package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/lovelacelee/clsgo/v1/log"
	"github.com/lovelacelee/clsgo/v1/utils"
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
	s, _ := io.ReadAll(resp.Body)
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

	result, _ := io.ReadAll(resp.Body)
	return string(result), nil
}

// descPrefix used for progressbar, format is like
// "[cyan][1/3][reset]", "Downloading..."
func Download(url string, filepath string, descPrefix ...string) (int64, error) {
	out, err := os.Create(filepath)
	if err != nil {
		return 0, err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0, utils.ErrNotFound
	}
	defer resp.Body.Close()
	bar := utils.ProgressBar(resp.ContentLength, 0, "█", descPrefix...)
	n, err := io.Copy(io.MultiWriter(out, bar), resp.Body)
	return n, err
}
