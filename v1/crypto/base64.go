package crypto

import (
	"encoding/base64"
)

// Base64FromString encode string to base64 string in standard mode
func Base64FromString(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64 encode data to string in stanard mode
func Base64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64Decode decode cipher string to origin data in bytes
func Base64Decode(cipher string) []byte {
	assumelen := base64.StdEncoding.DecodedLen(len(cipher))
	decode := make([]byte, assumelen)
	n, _ := base64.StdEncoding.Decode(decode, []byte(cipher))
	return decode[0:n]
}

// Base64ToString decode cipher string to origin string
func Base64ToString(cipher string) string {
	return string(Base64Decode(cipher))
}

func Base64UrlSafe(in string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(in))
}

func Base64DecodeUrlSafe(cipher string) []byte {
	assumelen := base64.RawURLEncoding.DecodedLen(len(cipher))
	decode := make([]byte, assumelen)
	n, _ := base64.RawURLEncoding.Decode(decode, []byte(cipher))
	return decode[0:n]
}

func Base64DecodeStringUrlSafe(cipher string) string {
	return string(Base64DecodeUrlSafe(cipher))
}
