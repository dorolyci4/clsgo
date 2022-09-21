package utils

import (
	"bytes"
	"io"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"

	"github.com/gogf/gf/v2/encoding/gcharset"
)

var (
	// Chinese
	GB2312  = simplifiedchinese.HZGB2312
	GB18030 = simplifiedchinese.GB18030
	GBK     = simplifiedchinese.GBK
	Big5    = traditionalchinese.Big5
	// Korean
	EUCKR = korean.EUCKR
	// Japanese
	EUCJP     = japanese.EUCJP
	ISO2022JP = japanese.ISO2022JP
	ShiftJIS  = japanese.ShiftJIS
	// Unicode
	UTF8             = unicode.UTF8
	UTF8BOM          = unicode.UTF16(unicode.BigEndian, unicode.UseBOM)
	UCS2BigEndian    = unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	UCS2LittleEndian = unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
)

// Wrapper for gcharset of goframe
func StringConvert(from string, to string, src string) (string, error) {
	return gcharset.Convert(to, from, src)
}

func BytesDecode(coding encoding.Encoding, s []byte) ([]byte, string, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, coding.NewDecoder())
	d, e := io.ReadAll(O)
	return d, string(d), e
}

func CharsetEncode(coding encoding.Encoding, src []byte) []byte {
	encoder := coding.NewDecoder()
	dst := make([]byte, len(src))
	ndst, _, err := encoder.Transform(dst, src, true)
	if err != nil {
		return dst[:ndst]
	} else {
		return dst
	}
}

func CharsetDecode(coding encoding.Encoding, src []byte) []byte {
	decoder := coding.NewDecoder()
	dst := make([]byte, len(src))
	ndst, _, err := decoder.Transform(dst, src, true)
	if err != nil {
		return dst[:ndst]
	} else {
		return dst
	}
}

// Probe the bytes encoding and try to convert to UTF-8
func DetermineDecode(s []byte, contentType ...string) ([]byte, string, error) {
	determineEncoding, _, _ := charset.DetermineEncoding(s, Param(contentType, ""))
	utf8Reader := transform.NewReader(bytes.NewReader(s), determineEncoding.NewDecoder())
	buf := make([]byte, len(s))
	_, err := utf8Reader.Read(buf)
	return buf, string(buf), err
}
