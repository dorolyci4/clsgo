package crypto_test

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/v1/crypto"
	"github.com/lovelacelee/clsgo/v1/log"
)

const want = "clsgo is a framework of common project work."

func TestAes(t *testing.T) {
	key := "lovelacelee"

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

func TestPKCS7(t *testing.T) {
	src := want
	gtest.C(t, func(t *gtest.T) {
		t.Run("PAD-UNPAD", func(_ *testing.T) {
			var err error
			crypto.PadPKCS7([]byte("test"), 256)
			crypto.PadPKCS7([]byte("test"), 0)
			crypto.UnpadPKCS7([]byte{}, 0)
			_, err = crypto.UnpadPKCS7([]byte{}, 32)
			t.Assert(err, crypto.ErrEmpty)
			_, err = crypto.UnpadPKCS7([]byte("test"), 32)
			t.Assert(err, crypto.ErrNotFullBlocks)
			_, err = crypto.UnpadPKCS7([]byte("test"), 2)
			t.Assert(err, crypto.ErrBadPadding)

			pad, err := crypto.PadPKCS7([]byte(src), 160)
			pad[len(pad)-1] = 0x9E //158 < blocksize(160), 158 > len(pad):116
			unpad, err := crypto.UnpadPKCS7(pad, 160)
			t.Assert(err, crypto.ErrBadPadding)
			t.Assert(unpad, nil)
			crypto.UnpadPKCS7([]byte{}, 256)
		})
	})
}

func TestMD5(t *testing.T) {
	// Test case match http://md5.cn/
	src := want
	gtest.C(t, func(t *gtest.T) {
		t.Run("Sum", func(_ *testing.T) {
			t.Assert(crypto.MD5Sum(src), "848129e9a404ac48092a3920911c660f")
			md5 := crypto.NewMD5(src)
			t.Assert(md5.Sum(), "848129e9a404ac48092a3920911c660f")
			t.Assert(md5.SumUpper(), "848129E9A404AC48092A3920911C660F")
			t.Assert(crypto.Md5Any("clsgo is a f", "ramework of com", "mon project work."), "848129e9a404ac48092a3920911c660f")
			t.Assert(md5.SumUpper(crypto.MD5_16), "A404AC48092A3920")
			t.Assert(crypto.MD5Sum(crypto.MD5Sum("admin")), "c3284d0f94606de1fd2af172aba15bf3")
		})
	})
}

func TestMain(m *testing.M) {
	log.Green("")
	m.Run()
}
