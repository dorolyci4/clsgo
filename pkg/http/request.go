package http

import (
	"bytes"
	"encoding/json"
	"github.com/lovelacelee/clsgo/pkg/log"
	"io/ioutil"
	"net/http"
	"time"
)

func Get(url string, timeout ...int) (string, error) {
	var t time.Duration = 5
	if len(timeout) > 0 {
		t = time.Duration(timeout[0])
	}
	client := &http.Client{Timeout: t * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		log.Errori(err)
		return "{}", err
	}
	defer resp.Body.Close()
	s, _ := ioutil.ReadAll(resp.Body)
	return string(s), nil
}

func Post(url string, data interface{}, contentType string, timeout ...int) (string, error) {
	var t time.Duration = 5
	if len(timeout) > 0 {
		t = time.Duration(timeout[0])
	}
	client := &http.Client{Timeout: t * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Errori(err)
		return "{}", err
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result), nil
}
