/*
 * Test result match: https://www.mklab.cn/utils/aes
 */

package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"encoding/base64"
	"encoding/hex"
	"errors"
	"github.com/lovelacelee/clsgo/pkg/log"
	"io"
)

const (
	ECB = 1
	CBC = 2
	CFB = 3
)

const (
	PKCS5 = iota
	PKCS7
	NoPadding
)

type AES struct {
	// Private
	key []byte
	// Public
	Encrypted []byte
	// Public
	Origin []byte
	// PKCS5/PKCS7/NoPadding
	Padding int
	// ECB/CBC/CFB
	Mod int
	// IV length must be 16 bytes long, IvGen could generate one
	Iv []byte
}

// Fix keystr to 16 bytes key[]byte
// Must called before encrypt or decrypt method
// keylen only accept [16/24/32]
func (aes *AES) Key(keystr string, keylen int) {
	if keylen != 16 && keylen != 24 && keylen != 32 {
		log.Error("AES key only accept [16/24/32] len, use default 16.")
		keylen = 16
	}
	key := []byte(keystr)
	aes.key = make([]byte, keylen)
	copy(aes.key, key)
	for i := keylen; i < len(key); {
		for j := 0; j < keylen && i < len(key); j, i = j+1, i+1 {
			aes.key[j] ^= key[i]
		}
	}
}

func (aes *AES) IvGen() []byte {
	iv := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil
	}
	aes.Iv = iv
	return iv
}

func (aes *AES) EncryptedBase64() string {
	return base64.StdEncoding.EncodeToString(aes.Encrypted)
}

func (aes *AES) EncryptedHex() string {
	return hex.EncodeToString(aes.Encrypted)
}

func (aes *AES) DecryptedString() string {
	return string(aes.Origin)
}

func (aes *AES) Encrypt(data []byte) ([]byte, error) {
	var err error
	switch aes.Mod {
	case ECB:
		aes.Encrypted = aesEncryptECB(data, aes.key)
		return aes.Encrypted, nil
	case CBC:
		if aes.Padding == NoPadding {
			return nil, errors.New("aes CBC mode only support PKCS #5 or PKCS #7")
		}
		aes.Encrypted, err = aesEncryptCBC(data, aes.Padding, aes.key, aes.Iv)
		return aes.Encrypted, err
	case CFB:
		aes.Encrypted, err = aesEncryptCFB(data, aes.Padding, aes.key, aes.Iv)
		return aes.Encrypted, err
	}

	return nil, errors.New("AES mod is not supported")
}

func (aes *AES) Decrypt(cipher []byte) ([]byte, error) {
	var err error
	switch aes.Mod {
	case ECB:
		aes.Origin = aesDecryptECB(cipher, aes.key)
		return aes.Origin, nil
	case CBC:
		aes.Origin, err = aesDecryptCBC(cipher, aes.Padding, aes.key, aes.Iv)
		return aes.Origin, err
	case CFB:
		aes.Origin, err = aesDecryptCFB(cipher, aes.Padding, aes.key, aes.Iv)
		return aes.Origin, err
	}

	return nil, errors.New("AES mod is not supported")
}

func aesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(key)
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// encrypt by group
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}
func aesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(key)
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

// The length of iv must be the same as the Block's block size.
func aesEncryptCBC(origData []byte, padding int, key []byte, iv []byte) (encrypted []byte, err error) {
	// len(key) must be 16/24/32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()

	switch padding {
	case PKCS5:
		origData = PadPKCS5(origData, blockSize)
		err = nil
	case PKCS7:
		origData, err = PadPKCS7(origData, aes.BlockSize)
	case NoPadding:
		err = nil
	}
	blockMode := cipher.NewCBCEncrypter(block, iv)
	encrypted = make([]byte, len(origData))
	blockMode.CryptBlocks(encrypted, origData)
	return encrypted, err
}

// The length of iv must be the same as the
// Block's block size and must match the iv used to encrypt the data.
func aesDecryptCBC(encrypted []byte, padding int, key []byte, iv []byte) (decrypted []byte, err error) {
	// len(key) must be 16/24/32
	block, _ := aes.NewCipher(key)
	blockMode := cipher.NewCBCDecrypter(block, iv)
	decrypted = make([]byte, len(encrypted))
	blockMode.CryptBlocks(decrypted, encrypted)

	switch padding {
	case PKCS5:
		decrypted = UnpadPKCS5(decrypted)
		err = nil
	case PKCS7:
		decrypted, err = UnpadPKCS7(decrypted, aes.BlockSize)
	case NoPadding:
		err = nil
	}
	return decrypted, err
}

func aesEncryptCFB(origData []byte, padding int, key []byte, iv []byte) (encrypted []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	switch padding {
	case PKCS5:
		origData = PadPKCS5(origData, block.BlockSize())
		err = nil
	case PKCS7:
		origData, err = PadPKCS7(origData, aes.BlockSize)
	case NoPadding:
		err = nil
	}
	encrypted = make([]byte, len(origData))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted, origData)
	return encrypted, err
}

func aesDecryptCFB(encrypted []byte, padding int, key []byte, iv []byte) (decrypted []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	if len(encrypted) < aes.BlockSize {
		return nil, errors.New("cipher content(encrypted) too short")
	}
	decrypted = make([]byte, len(encrypted))
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decrypted, encrypted)
	switch padding {
	case PKCS5:
		decrypted = UnpadPKCS5(decrypted)
		err = nil
	case PKCS7:
		decrypted, err = UnpadPKCS7(decrypted, aes.BlockSize)
	case NoPadding:
		err = nil
	}
	return decrypted, err
}
