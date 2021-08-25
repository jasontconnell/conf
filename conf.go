package conf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Conf struct{}

func LoadConfig(configFile string, conf interface{}) error {
	file, err := os.Open(configFile)

	if err != nil {
		return fmt.Errorf("opening config file %s. %w", configFile, err)
	}
	defer file.Close()

	return fromReader(file, conf)
}

func DecodeRawMessage(raw json.RawMessage) (interface{}, error) {
	b := bytes.NewBuffer(raw)
	var decoded interface{}
	err := fromReader(b, &decoded)
	return decoded, err
}

func FromUrl(configUrl string, conf interface{}) error {
	return FromUrlBasicAuth(configUrl, "", "", conf)
}

func FromUrlBasicAuth(configUrl, username, password string, conf interface{}) error {
	req, err := http.NewRequest("GET", configUrl, nil)
	if err != nil {
		return err
	}

	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("couldn't get url: %s", configUrl)
	}

	return fromReader(resp.Body, conf)
}

func fromReader(r io.Reader, conf interface{}) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(conf)
}
