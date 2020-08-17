package main

import (
	"fmt"
	"io"
	"reflect"
)

// to reflect on an interface type at runtime
// use reflect.Type

type Name struct {
	First, Last string
}

func (n *Name) String() string {
	return n.First + " " + n.Last
}

func main() {

	n := &Name{First: "Fero", Last: "Janus"}

	stringer := (*fmt.Stringer)(nil) // creates a nil pointer of type fmt.Stringer
	implements(n, stringer)

	writer := (*io.Writer)(nil)
	implements(n, writer)
}

func implements(concrete, target interface{}) bool {
	iface := reflect.TypeOf(target).Elem()

	v := reflect.ValueOf(concrete)
	t := v.Type()

	if t.Implements(iface) {
		fmt.Printf("%T is a %s\n", concrete, iface.Name())
		return true
	}
	fmt.Printf("%T is not a %s\n", concrete, iface.Name())
	return false
}
