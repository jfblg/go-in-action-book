package main

import (
	"encoding/json"
	"fmt"
)

type Name struct {
	First string `json:"firstName" xml:"FirstName"`
	Last  string `json:"lastName,omitempty"`
}

func main() {
	n := &Name{"fero", "janus"}
	fmt.Printf("Type: %T\n", n)
	jdata, _ := json.Marshal(n)
	fmt.Printf("%s\n", jdata)

}
