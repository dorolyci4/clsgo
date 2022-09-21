package crypto_test

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/v1/crypto"
	"github.com/lovelacelee/clsgo/v1/log"
)

func TestAes(t *testing.T) {
	key := "lovelacelee"
	want := "clsgo is a framework of common project work."

	gtest.C(t, func(t *gtest.T) {
		t.Run("ECB", func(_ *testing.T) {
			// ECB test
			aesECB := crypto.AES{Mod: crypto.ECB, Padding: crypto.PKCS7}
			aesECB.Key(key, 16)
			enc, _ := aesECB.Encrypt([]byte(want))
			aesECB.Decrypt(enc)
			t.Assert(aesECB.DecryptedString(), want)
			t.Assert(aesECB.EncryptedBase64(), "L/DafEZBAnVIBK+nP+jN+1TzljDyF1ttT6YxkXOO8QiLc+b0TI2qV3IVe4aVTr+Q")
		})
		t.Run("CBC", func(_ *testing.T) {
			// CBC test
			aesCBC := crypto.AES{Mod: crypto.CBC, Padding: crypto.PKCS5}
			aesCBC.IvGen("test iv")
			aesCBC.Key(key, 16)
			encCBC, _ := aesCBC.Encrypt([]byte(want))
			aesCBC.Decrypt(encCBC)
			t.Assert(aesCBC.DecryptedString(), want)
			t.Assert(aesCBC.EncryptedBase64(), "oIo68+csok8WOyCnH23lEZyvsTESZ8OaoKJH6x2NmRmVY+d8zpJdwAsM3vD8eOwE")

			aesCBC.IvGen("test iv too long: expected trimed result")
			aesCBC.EncryptedHex()

			aesCBC = crypto.AES{Mod: crypto.CBC, Padding: crypto.PKCS7}
			aesCBC.IvGen("test iv")
			aesCBC.Key(key, 32)
			aesCBC.EncryptedHex()
			encCBC, _ = aesCBC.Encrypt([]byte(want))
			aesCBC.Decrypt(encCBC)

			aesCBC = crypto.AES{Mod: crypto.CBC, Padding: crypto.NoPadding}
			aesCBC.IvGen("test iv")
			aesCBC.Key("key: this test key is to long, longer than 16 bytes", 169) //16 will be used
			aesCBC.Encrypt([]byte(want))
			encCBC, _ = aesCBC.Encrypt([]byte(want))
			aesCBC.Decrypt(encCBC)
		})

		t.Run("CBC-RandomIV", func(_ *testing.T) {
			// CBC test
			aesCBC := crypto.AES{Mod: crypto.CBC, Padding: crypto.PKCS5}
			aesCBC.IvGen()
			aesCBC.Key(key, 16)
			encCBC, _ := aesCBC.Encrypt([]byte(want))
			aesCBC.Decrypt(encCBC)
			t.Assert(aesCBC.DecryptedString(), want)
			// Generate different cipher everytime because of IvGen()
			t.AssertNE(aesCBC.EncryptedBase64(), "oIo68+csok8WOyCnH23lEZyvsTESZ8OaoKJH6x2NmRmVY+d8zpJdwAsM3vD8eOwE")
		})
		t.Run("CFB", func(_ *testing.T) {
			// CFB test
			aesCFB := crypto.AES{Mod: crypto.CFB, Padding: crypto.NoPadding, Iv: []byte("1111111111111111")}
			aesCFB.Key(key, 16)
			encCFB, _ := aesCFB.Encrypt([]byte(want))
			aesCFB.Decrypt(encCFB)
			t.Assert(aesCFB.DecryptedString(), want)
			t.Assert(aesCFB.EncryptedBase64(), "4QKYnxmTnsbOVjeNqRDHpAfndf5XxqHWBqWgUYCND9p+DZOCQcZoluA9NQM=")

			aesCFB = crypto.AES{Mod: crypto.CFB, Padding: crypto.PKCS5, Iv: []byte("1111111111111111")}
			aesCFB.Key(key, 16)
			encCFB, _ = aesCFB.Encrypt([]byte(want))
			aesCFB.Decrypt(encCFB)

			aesCFB = crypto.AES{Mod: crypto.CFB, Padding: crypto.PKCS7, Iv: []byte("1111111111111111")}
			aesCFB.Key(key, 16)
			encCFB, _ = aesCFB.Encrypt([]byte(want))
			aesCFB.Decrypt(encCFB)
			aesCFB.Key("4QKYnxmTnsbOVjeNqRDHpAfndf5XxqHWBqWgUYCND9p", 16)
			aesCFB.Decrypt([]byte("test")) //error
		})
		t.Run("INVALID", func(_ *testing.T) {
			aesCBC := crypto.AES{Mod: 255, Padding: crypto.PKCS5}
			aesCBC.Encrypt([]byte("test"))
			aesCBC.Decrypt([]byte("xxxx"))
		})
	})

}

func TestMain(m *testing.M) {
	log.Green("")
	m.Run()
}
