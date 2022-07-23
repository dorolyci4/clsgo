/*
 * PKCS #5 only padding 8bytes(BlockSize=8), padding data will be 0x01-0x08
 */

package crypto

import (
	"bytes"
)

func PadPKCS5(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func UnpadPKCS5(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
