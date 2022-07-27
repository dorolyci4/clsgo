package http

import (
	"io/ioutil"
	"net/http"

	"github.com/lovelacelee/clsgo/pkg/log"
)

func GetAsJson(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "{}", err
	}
	s, _ := ioutil.ReadAll(resp.Body)
	log.Info(s)
	return string(s), nil
}
