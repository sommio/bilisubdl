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

func GetJson(t interface{}, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("HTTP error code %d", resp.StatusCode))
	}

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

	return strings.TrimSpace(strings.Trim(t, "."))
}

func CreateSubFile(filename, content string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
