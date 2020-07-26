package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Note: fields must be exported so that they are populated with values by Decoder
type configuration struct {
	Enabled bool
	Name    string
	Path    string
}

func main() {
	file, _ := os.Open("config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("name: %s, enabled: %v, path: %s\n", conf.Name, conf.Enabled, conf.Path)

}
