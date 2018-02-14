package conf

import (
	"os"
	"fmt"
	"encoding/json"
	"net/http"
	"io"
	"errors"
)

type Conf struct{}

func LoadConfig(configFile string, conf interface{}) error {
	file, err := os.Open(configFile)
 
	if err != nil {
		fmt.Println("Error opening json file ", err, configFile)
	}
 
 	return fromReader(file, conf)
}

func FromUrl(configUrl string, conf interface{}) error {
	return FromUrlBasicAuth(configUrl, "", "", conf)
}

func FromUrlBasicAuth(configUrl, username, password string, conf interface{}) error {
	req,err := http.NewRequest("GET", configUrl, nil)
	if err != nil {
		return err
	}

	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	client := &http.Client{}

	resp,err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("couldn't get url: " + configUrl)
	}

	return fromReader(resp.Body, conf)
}

func fromReader(r io.Reader, conf interface{}) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(conf)
}