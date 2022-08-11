package utils

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/google/uuid"
	"io"
)

// length must > 0
func SessionId(length ...int) string {
	buf := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return ""
	}
	base64 := base64.URLEncoding.EncodeToString(buf)

	if !IsEmpty(length) && length[0] > 0 && length[0] <= len(base64) {
		return base64[:length[0]]
	}
	return base64
}

func UUID() string {
	return uuid.New().String()
}

// Time based and error ignored, not safe way
func UUIDV1() string {
	uuid, _ := uuid.NewUUID()
	return uuid.String()
}
