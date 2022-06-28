package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ClsJson struct {
	path    string
	content []byte
}

// Open and read json file
func (obj *ClsJson) Load(path string) error {
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
func (obj *ClsJson) String() string {
	return string(obj.content)
}

// Json to struct
func (obj *ClsJson) Parse(model interface{}) error {
	// interface{} is the type of anything
	return json.Unmarshal([]byte(string(obj.content)), model)
}

// Struct
func (obj *ClsJson) Save(model interface{}, filepath string) error {
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
