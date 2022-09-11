package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

// ord use char as parameter, rerturn ascii number
var Ord = func(char byte) int {
	return int(char)
}

// Covert bytes to hex bytes, aim to make []byte{0x5B, 0x47, 0x4F, 0x5D}) ==> []byte("5b474f5d")
func BytesToHex(src []byte) []byte {
	maxEnLen := hex.EncodedLen(len(src))
	dst := make([]byte, maxEnLen)
	hex.Encode(dst, src)
	return dst
}

// BytesToHex BytesToHex, aim to make []byte("5b474f5d")) ==> []byte{0x5B, 0x47, 0x4F, 0x5D}
func HexToBytes(src []byte) []byte {
	maxEnLen := hex.DecodedLen(len(src))
	dst := make([]byte, maxEnLen)
	hex.Decode(dst, src)
	return dst
}

func StringToHex(s string) []byte {
	return BytesToHex([]byte(s))
}

func HexToString(src []byte) string {
	return string(HexToBytes(src))
}

func HexString(src []byte) string {
	return hex.EncodeToString(src)
}

func HexDump(src []byte, show ...func(fmt string, args ...any)) {
	callback := Param(show, InfoWithoutHeader)
	str := hex.Dump(src)
	callback(str)
}

// T can be bool, int8/16/32/64, uint8/16/32/64, float32/64
// T also could be slice of type bool/int8~64, uint8~64, float32~64
func NumberToBytes[T any](n T, bigEndian ...bool) []byte {
	bytebuf := bytes.NewBuffer([]byte{})
	if Param(bigEndian, false) {
		binary.Write(bytebuf, binary.BigEndian, n)
	} else {
		binary.Write(bytebuf, binary.LittleEndian, n)
	}
	return bytebuf.Bytes()
}

// T can be bool, int8/16/32/64, uint8/16/32/64, float32/64
// T also could be slice of type bool/int8~64, uint8~64, float32~64
func BytesToNumber[T any](nbytes []byte, bigEndian ...bool) (data T) {
	bytebuff := bytes.NewBuffer(nbytes)
	if Param(bigEndian, false) {
		binary.Read(bytebuff, binary.BigEndian, &data)
	} else {
		binary.Read(bytebuff, binary.LittleEndian, &data)
	}
	return data
}
