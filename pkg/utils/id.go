package utils

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

func SessionId() string {
	buf := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(buf)
}
