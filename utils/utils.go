package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var client http.Client

type Response struct {
	*http.Response
}

func Request(url string) (*Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("HTTP error code %d", resp.StatusCode))
	}

	return &Response{resp}, nil
}

func (resp *Response) Json(t interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.Unmarshal(body, &t)
	if err != nil {
		return err
	}
	return nil
}

func SecondToTime(tt float64) string {
	secs, msec := int64(tt), int64(tt*1000)%1000
	mins, secs := secs/60, secs%60
	hrs, mins := mins/60, mins%60
	return fmt.Sprintf("%02d:%02d:%02d,%03d", hrs, mins, secs, msec)
}

func CleanText(t string) string {
	toBeReplaces := []string{"\"", "?", "/", ":", "\\", "*", "<", ">", "|"}
	for _, elem := range toBeReplaces {
		t = strings.ReplaceAll(t, elem, "_")
	}
	t = strings.ReplaceAll(t, "\n", " ")

	return strings.TrimSpace(strings.Trim(t, "."))
}

func WriteFile(filename string, content []byte) error {
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		return err
	}

	_, err = f.Write(content)
	if err != nil {
		return err
	}
	return nil
}
