package main

import (
	"encoding/json"
	"fmt"
)

// if you don't know the structure, parse into an interface{} type

var ks = []byte(`{
    "firstName": "Jean",
    "lastName": "Bartik",
    "age": 86,
    "education": [
        {
            "institution": "Northwest Missouri State Teachers College",
            "degree": "Bachelor of Science in Mathematics"
        },
        {
            "institution": "University of Pennsylvania",
            "degree": "Masters in English"
        }
    ],
    "spouse": "William Bartik",
    "children": [
        "Timothy John Bartik",
        "Jane Helen Bartik",
        "Mary Ruth Bartik"
    ]
}`)

func main() {
	var f interface{}
	err := json.Unmarshal(ks, &f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)

	m := f.(map[string]interface{})
	fmt.Println(m["firstName"])
	printJSON(f)
}

func printJSON(v interface{}) {
	switch vv := v.(type) {
	case string:
		fmt.Println("is string: ", v)
	case float64:
		fmt.Println("is float64: ", v)
	case []interface{}:
		fmt.Println("is an array")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	case map[string]interface{}:
		fmt.Println("is an object:")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	default:
		fmt.Println("unknown type")
	}
}
