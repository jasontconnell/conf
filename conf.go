package conf

import (
	"os"
	"fmt"
	"encoding/json"
)

type Conf struct{}

func LoadConfig(configFile string, conf interface{}) {
	file, err := os.Open(configFile)
 
	if err != nil {
		fmt.Println("Error opening json file ", err, configFile)
	}
 
	decoder := json.NewDecoder(file)
	err2 := decoder.Decode(conf)
	if err2 != nil {
		fmt.Println("error decoding json file", err2, configFile);
	}
}