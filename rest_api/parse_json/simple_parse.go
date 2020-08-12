package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
}

var JSON = `{
	"name": "Fero"
}
`

func main() {
	var p Person
	err := json.Unmarshal([]byte(JSON), &p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Value: %q - Type: %T\n", p, p)

}
