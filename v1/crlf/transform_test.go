package crlf_test

import (
	"io"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/v1/crlf"
	"github.com/lovelacelee/clsgo/v1/log"
	"golang.org/x/text/transform"
)

func TestNormalize(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{"hello, world\r\n", "hello, world\n"},
		{"hello, world\r", "hello, world\n"},
		{"hello, world\n", "hello, world\n"},
		{"", ""},
		{"\r\n", "\n"},
		{"hello,\r\nworld", "hello,\nworld"},
		{"hello,\rworld", "hello,\nworld"},
		{"hello,\nworld", "hello,\nworld"},
		{"hello,\n\rworld", "hello,\n\nworld"},
		{"hello,\r\n\r\nworld", "hello,\n\nworld"},
	}

	n := new(crlf.Normalize)

	for _, c := range testCases {
		got, _, err := transform.String(n, c.in)
		if err != nil {
			t.Errorf("error transforming %q: %v", c.in, err)
			continue
		}
		if got != c.want {
			t.Errorf("transforming %q: got %q, want %q", c.in, got, c.want)
		}
	}
}

func TestToCRLF(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{"hello, world\n", "hello, world\r\n"},
		{"", ""},
		{"\n", "\r\n"},
		{"hello,\nworld", "hello,\r\nworld"},
		{"hello,\n\nworld", "hello,\r\n\r\nworld"},
	}

	for _, c := range testCases {
		got, _, err := transform.String(crlf.ToCRLF{}, c.in)
		if err != nil {
			t.Errorf("error transforming %q: %v", c.in, err)
			continue
		}
		if got != c.want {
			t.Errorf("transforming %q: got %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCRLF(t *testing.T) {
	testCases := []struct {
		in   string
		want string
		conv func(string) string
	}{
		{"hello, world\n", "hello, world\r\n", crlf.CRLFString}, //CR->CRLF
		{"", "", crlf.CRLFString},                               //NIL
		{"\r", "\r", crlf.CRLFString},                           //LF->CRLF
		{"\r", "\n", crlf.CRString},                             //LF->CR
		{"hello,\r\nworld", "hello,\nworld", crlf.CRString},     //CRLF->CR
		{"hello,\r\n\nworld", "hello,\n\nworld", crlf.CRString}, //CRLF->CR
	}
	gtest.C(t, func(t *gtest.T) {
		t.Run("", func(_ *testing.T) {
			for _, tc := range testCases {
				t.Assert(tc.conv(tc.in), tc.want)
			}
		})
	})
	gtest.C(t, func(t *gtest.T) {
		t.Run("", func(_ *testing.T) {
			var dst [5]byte
			var src = "\r\n\r\r\r\n\r\n\r\r"
			var dstX [20]byte
			n := new(crlf.Normalize)
			n.Transform(dst[0:5], []byte(src), false)
			x := new(crlf.ToCRLF)
			x.Transform(dst[0:5], []byte(src), false)
			x.Transform(dstX[0:10], []byte(src), false)
		})
	})
	gtest.C(t, func(t *gtest.T) {
		t.Run("", func(_ *testing.T) {
			r := crlf.NewReader(io.MultiReader())
			r.Read([]byte{})
			w := crlf.NewWriter(io.MultiWriter())
			w.Write([]byte{})
		})
	})
}

func TestMain(m *testing.M) {
	log.Green("")
	m.Run()
}
