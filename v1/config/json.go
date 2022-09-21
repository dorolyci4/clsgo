package config

import (
	"bytes"
	"encoding/json"
	"io"
	"os"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/lovelacelee/clsgo/v1/utils"
)

type Json = gjson.Json
type Options = gjson.Options

// A JSON contains the JSON standard byte stream, and the file path to which you want to store it
type JSON struct {
	path    string
	content []byte
}

func NewJsonWith(file string, in []byte) *JSON {
	return &JSON{
		content: in,
		path:    file,
	}
}

// Create an empty JSON object
func NewJson() *JSON {
	return NewJsonWith("", []byte{})
}

func (obj *JSON) FromString(s string) {
	obj.content = []byte(s)
}

func (obj *JSON) OutputTo(file string) {
	obj.path = file
}

// Open and read json file
func (obj *JSON) FromFile(path string) error {
	obj.path = path
	f, err := os.Open(obj.path)
	if err != nil {
		return err
	}
	defer f.Close()
	content, _ := io.ReadAll(f)
	obj.content = bytes.TrimPrefix(content, []byte("\xef\xbb\xbf"))
	return err
}

// Json string
func (obj *JSON) String() string {
	return string(obj.content)
}

// Parse byte strem to struct
func (obj *JSON) Unmarshal(model interface{}) error {
	// interface{} is the type of anything
	return JsonUnmarshal([]byte(string(obj.content)), model)
}

// Encode golang object to json stream
func (obj *JSON) Marshal(model interface{}) error {
	b, err := JsonMarshal(model)
	obj.content = b
	return err
}

// Save byte stream to json file
func (obj *JSON) Save(filepath ...string) error {
	file := utils.Param(filepath, "output.json")
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	var buffer bytes.Buffer
	err = json.Indent(&buffer, obj.content, "", "\t")
	_, werr := buffer.WriteTo(f)

	return utils.NewError(err).Join(werr)
}

// Functions implement by gjson

func JsonDecode(data interface{}, options ...Options) (interface{}, error) {
	return gjson.Decode(data, options...)
}
func JsonDecodeTo(data interface{}, v interface{}, options ...Options) (err error) {
	return gjson.DecodeTo(data, v, options...)
}
func JsonEncode(value interface{}) ([]byte, error) {
	return gjson.Encode(value)
}
func JsonEncodeString(value interface{}) (string, error) {
	return gjson.EncodeString(value)
}
func JsonIsValidDataType(dataType string) bool {
	return gjson.IsValidDataType(dataType)
}
func JsonMarshal(v interface{}) (marshaledBytes []byte, err error) {
	return gjson.Marshal(v)
}
func JsonMarshalIndent(v interface{}, prefix, indent string) (marshaledBytes []byte, err error) {
	return gjson.MarshalIndent(v, prefix, indent)
}
func JsonMustEncode(value interface{}) []byte {
	return gjson.MustEncode(value)
}
func JsonMustEncodeString(value interface{}) string {
	return gjson.MustEncodeString(value)
}
func JsonUnmarshal(data []byte, v interface{}) (err error) {
	return gjson.Unmarshal(data, v)
}
func JsonValid(data interface{}) bool {
	return gjson.Valid(data)
}
