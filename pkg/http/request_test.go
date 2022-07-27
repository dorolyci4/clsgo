package http_test

import (
	"github.com/lovelacelee/clsgo/pkg/http"
	"github.com/lovelacelee/clsgo/pkg/log"
	"testing"
)

func ExampleGetAsJson() {
	s, err := http.GetAsJson("https://restapi.amap.com/v3/assistant/coordinate/convert")
	log.Info(s)
	log.Error(err)
}

func TestGetAsJson(t *testing.T) {
	ExampleGetAsJson()
}
