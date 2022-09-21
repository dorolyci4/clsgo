package http_test

import (
	// "fmt"
	"testing"

	// "github.com/lovelacelee/clsgo/v1/http"
	"github.com/lovelacelee/clsgo/v1/utils"
)

func clean() {
	utils.DeletePath("logs")
	utils.DeleteFiles(utils.Cwd(), "/*.yaml$")
	utils.DeleteFiles(utils.Cwd(), "/*.xml$")
}

// func ExampleGet() {
// 	// s, _ := http.Get("https://restapi.amap.com/v3/assistant/coordinate/convert")
// 	// fmt.Println(s)
// 	// Output:
// 	// {"status":"0","info":"INVALID_USER_KEY","infocode":"10001"}

// }

func TestGet(t *testing.T) {
	// ExampleGet()

	clean()
}
