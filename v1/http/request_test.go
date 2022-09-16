package http_test

import (
	"fmt"
	"github.com/lovelacelee/clsgo/v1/http"
	"testing"
)

func ExampleGet() {
	s, _ := http.Get("https://restapi.amap.com/v3/assistant/coordinate/convert")
	fmt.Println(s)
	// Output:
	// {"status":"0","info":"INVALID_USER_KEY","infocode":"10001"}
}

func TestGet(t *testing.T) {
	// ExampleGet()
}
