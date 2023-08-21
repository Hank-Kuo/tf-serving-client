package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	HOST = "localhost:8501/v1/models/bert:predict"
)

const TIMEOUT = 1000 * time.Millisecond

func QueryRestful(data map[string][][]string, url string, target map[string][][]float32) error {
	buffer, err := json.Marshal(data)
	if err != nil {
		return err
	}

	client := http.Client{Timeout: TIMEOUT}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(buffer))

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	respBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(string(respBuf))
	}

	err = json.Unmarshal(respBuf, &target)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	query := make(map[string][][]string)
	query["inputs"] = [][]string{{"Helle world"}}
	response := make(map[string][][]float32)

	if err := tf_serving.QueryRestful(query, URL, response); err != nil {
		panic(err)
	}
}
