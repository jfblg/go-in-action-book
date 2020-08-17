package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type MyFloat64 float64

func main() {
	var a uint8 = 10
	var b int = 5
	var c float64 = 2.5
	var d string = "2.5"
	var e MyFloat64 = 10
	res := sumType(a, b, c, d, e)
	fmt.Printf("Result %f\n", res)
	res = sumKind(a, b, c, d, e)
	fmt.Printf("Result %f\n", res)
}

func sumType(v ...interface{}) float64 {
	var res float64 = 0
	for _, val := range v {
		switch val.(type) {
		case int:
			res += float64(val.(int))
		case int64:
			res += float64(val.(int64))
		case float64:
			res += val.(float64)
		case uint8:
			res += float64(val.(uint8))
		case string:
			a, err := strconv.ParseFloat(val.(string), 64)
			if err != nil {
				panic(err)
			}
			res += a
		default:
			fmt.Printf("Unsupported type %T. Skipping.\n", val)
		}
	}
	return res
}

// using a Kind switch to detect a kind of a value
// advantage: also
func sumKind(v ...interface{}) float64 {
	var res float64 = 0
	for _, val := range v {
		ref := reflect.ValueOf(val)
		switch ref.Kind() {
		case reflect.Int, reflect.Int64:
			res += float64(ref.Int())
		case reflect.Uint8:
			res += float64(ref.Uint())
		case reflect.Float64:
			res += ref.Float()
		case reflect.String:
			a, err := strconv.ParseFloat(ref.String(), 64)
			if err != nil {
				panic(err)
			}
			res += a
		default:
			fmt.Printf("Unsupported type %T. Skipping.\n", val)

		}
	}

	return res
}
