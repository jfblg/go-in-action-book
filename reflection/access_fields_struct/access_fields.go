package main

import (
	"fmt"
	"reflect"
	"strings"
)

type MyInt int

type Person struct {
	Name    *Name
	Address *Address
}

type Name struct {
	Title, First, Last string
}

type Address struct {
	Street, Region string
}

func main() {
	fmt.Println("Walking a simple integer")
	var ten MyInt = 10
	walk(ten, 0)

	fmt.Println("Walking a simple struct")
	s := struct{ Name string }{"foo"}
	walk(s, 0)

	fmt.Println("Walking a struct with struct fields")
	p := &Person{
		Name:    &Name{"Prof.", "Oliver", "Twist"},
		Address: &Address{"London", "UK"},
	}
	walk(p, 0)

}

func walk(u interface{}, depth int) {
	val := reflect.Indirect(reflect.ValueOf(u))
	t := val.Type()
	tabs := strings.Repeat("\t", depth+1)

	fmt.Printf("%sValue type %q (%s)\n", tabs, t, val.Kind())
	if val.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldVal := reflect.Indirect(val.Field(i))

			tabs := strings.Repeat("\t", depth+2)
			fmt.Printf("%sFiled %q is type %q (%s)\n", tabs, field.Name, field.Type, fieldVal.Kind())

			if fieldVal.Kind() == reflect.Struct {
				walk(fieldVal.Interface, depth+1)
			}
		}
	}

}
