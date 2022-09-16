package crypto_test

import (
	"testing"

	"github.com/lovelacelee/clsgo/v1/crypto"
	"github.com/lovelacelee/clsgo/v1/log"
)

func TestAes(t *testing.T) {
	key := "lovelacelee"
	want := "clsgo is a framework of common project work."
	// ECB test
	aesECB := crypto.AES{Mod: crypto.ECB, Padding: crypto.PKCS7}
	aesECB.Key(key, 16)
	enc, _ := aesECB.Encrypt([]byte(want))
	aesECB.Decrypt(enc)
	if aesECB.DecryptedString() != want {
		t.Errorf("Not passed\n")
	} else {
		log.Info(aesECB.EncryptedBase64())
	}
	// CBC test
	aesCBC := crypto.AES{Mod: crypto.CBC, Padding: crypto.PKCS5}
	aesCBC.IvGen()
	aesCBC.Key(key, 16)
	encCBC, _ := aesCBC.Encrypt([]byte(want))
	aesCBC.Decrypt(encCBC)
	if aesCBC.DecryptedString() != want {
		t.Errorf("Not passed\n")
	} else {
		log.Info(aesCBC.EncryptedBase64())
	}
	// CFB test
	aesCFB := crypto.AES{Mod: crypto.CFB, Padding: crypto.NoPadding, Iv: []byte("1111111111111111")}
	aesCFB.Key(key, 16)
	encCFB, _ := aesCFB.Encrypt([]byte(want))
	aesCFB.Decrypt(encCFB)
	if aesCFB.DecryptedString() != want {
		t.Errorf("Not passed\n")
	} else {
		log.Info(aesCFB.EncryptedBase64())
	}
}
