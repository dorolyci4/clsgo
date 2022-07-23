package config

import (
	"bytes"
	"encoding/json"
	"github.com/gogf/gf/v2/encoding/gjson"
	"io/ioutil"
	"log"
	"os"
)

type Json = gjson.Json
type Options = gjson.Options

type JsonFile struct {
	path    string
	content []byte
}

// Open and read json file
func (obj *JsonFile) Load(path string) error {
	obj.path = path
	f, err := os.Open(obj.path)
	if err != nil {
		log.Fatalf("Open %s error %v\n", obj.path, err)
	}
	defer f.Close()
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("Read %s error %v\n", obj.path, err)
	}
	obj.content = bytes.TrimPrefix(content, []byte("\xef\xbb\xbf"))
	return err
}

// Json string
func (obj *JsonFile) String() string {
	return string(obj.content)
}

// Json to struct
func (obj *JsonFile) Parse(model interface{}) error {
	// interface{} is the type of anything
	return json.Unmarshal([]byte(string(obj.content)), model)
}

// Struct
func (obj *JsonFile) Save(model interface{}, filepath string) error {
	b, err := json.Marshal(model)
	if err != nil {
		return err
	}
	obj.content = b
	f, err := os.Create(filepath)
	if err != nil {
		log.Printf("%s create failed\n", filepath)
		return err
	}
	defer f.Close()

	var buffer bytes.Buffer
	err = json.Indent(&buffer, obj.content, "", "\t")
	if err != nil {
		log.Printf("%s format failed\n", filepath)
	}

	_, err = buffer.WriteTo(f)
	return err
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
