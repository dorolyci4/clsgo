package clsgo_test

import (
	"reflect"
	"testing"

	"github.com/lovelacelee/clsgo/pkg/config"
)

type JsonModel struct {
	Version string `json:"version"`
	Author  string `json:"author"`
	Github  string `json:"github"`
}

func TestJson(t *testing.T) {
	var json config.JsonFile
	//Load
	json.Load("test.json")

	var m JsonModel
	err := json.Parse(&m)
	l.Infof("%v %v %v %v", err, m.Author, m.Version, m.Github)
	//Read check
	if !reflect.DeepEqual(m.Author, "Lovelace") {
		t.Errorf("Json value not match: %s != %s\n", m.Author, "Lovelace")
		t.Logf(json.String())
	}
	//Modification
	m.Author = "Lovelace"
	json.Save(m, "test.json")
	l.Info(json.String())
}

func BenchmarkJson(b *testing.B) {
	var json config.JsonFile
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Load("test.json")
		var m JsonModel
		json.Parse(&m)
		json.Save(m, "test.json")
	}
}
