/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-07-14 17:50:17
 * @LastEditTime    : 2022-07-15 10:00:07
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /pkg/crypto/pkcs5.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
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
