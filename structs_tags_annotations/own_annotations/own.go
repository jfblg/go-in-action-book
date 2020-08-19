package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Processes ...
type Processes struct {
	Total     int     `init:"total"`
	Running   int     `init:"running"`
	Sleepting int     `init:"sleeping"`
	Threads   int     `init:"threads"`
	Load      float32 `init:"load"`
}

func main() {
	fmt.Println("Write a struct to output")
	p := &Processes{
		Total:     33,
		Running:   10,
		Sleepting: 1,
		Threads:   100,
		Load:      10.4,
	}
	data, err := Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	fmt.Println("Reading data back to a struct")
	proc2 := &Processes{}
	if err := Unmarshal(data, proc2); err != nil {
		panic(err)
	}
	fmt.Printf("Struct: %#v", proc2)

}

func fieldName(field reflect.StructField) string {
	if t := field.Tag.Get("ini"); t != "" {
		return t
	}
	return field.Name
}

// Marshal for my custom INI format
func Marshal(v interface{}) ([]byte, error) {
	var b bytes.Buffer
	val := reflect.Indirect(reflect.ValueOf(v))
	if val.Kind() != reflect.Struct {
		return []byte{}, errors.New("Marshall can only take structs")
	}

	t := val.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name := fieldName(f)
		raw := val.Field(i).Interface()
		fmt.Fprintf(&b, "%s=%v\n", name, raw)
	}
	return b.Bytes(), nil
}

// Unmarshal for my custom INI format
func Unmarshal(data []byte, v interface{}) error {
	val := reflect.Indirect(reflect.ValueOf(v))
	t := val.Type()

	b := bytes.NewBuffer(data)
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.Split(line, "=")
		if len(pair) < 2 {
			// skip
			continue
		}
		setField(pair[0], pair[1], t, val)
	}
	return nil
}

// TODO
func setField(name, value string, t reflect.Type, v reflect.Value) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if name == fieldName(field) {
			var dest reflect.Value
			switch field.Type.Kind() {
			default:
				fmt.Printf("Kind %s is not supported\n", field.Type.Kind())
				continue
			case reflect.Int:
				ival, err := strconv.Atoi(value)
				if err != nil {
					fmt.Printf("Could not convert %q to int: %s\n", value, err)
					continue
				}
				dest = reflect.ValueOf(ival)
			case reflect.Float64:
				fval, err := strconv.ParseFloat(value, 64)
				if err != nil {
					fmt.Printf("Could not convert %q to float64: %s\n", value, err)
					continue
				}
				dest = reflect.ValueOf(fval)
			case reflect.String:
				dest = reflect.ValueOf(value)
			case reflect.Bool:
				bval, err := strconv.ParseBool(value)
				if err != nil {
					fmt.Printf("Could not convert %q to bool: %s\n", value, err)
					continue
				}
				dest = reflect.ValueOf(bval)
			}
			v.Field(i).Set(dest)
		}
	}
}
