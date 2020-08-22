package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

var tpl = `package {{.Package}}

// this is a generated code

type {{.MyType}}Queue struct {
	q []{{.MyType}}
}

func New{{.MyType}}Queue() *{{.MyType}}Queue {
	return &{{.MyType}}Queue{
		q: []{{.MyType}}{},
	}
}

func (o *{{.MyType}}Queue) Insert(v {{.MyType}}) {
	o.q = append(o.q, v)
}

func (o *{{.MyType}}Queue) Remove() {{.MyType}} {
	if len(o.q) == 0 {
		panic("{{.MyType}}Queue is empty")
	}
	first := o.q[0]
	o.q = o.q[1:]
	return first
}
`

func main() {
	tt := template.Must(template.New("queue").Parse(tpl))
	for i := 1; i < len(os.Args); i++ {
		dest := strings.ToLower(os.Args[i]) + "_queue.go"
		file, err := os.Create(dest)
		if err != nil {
			fmt.Printf("Could not create file %q. Error: %s Skipping.\n", dest, err)
			continue
		}
		defer file.Close()

		vals := map[string]string{
			"MyType":  os.Args[i],
			"Package": os.Getenv("GOPACKAGE"), // Note: $GOPACKAGE is pupulated by go generate and
			// is set to a the name of package where the //go:genearate header was found
		}
		tt.Execute(file, vals)
	}
}
