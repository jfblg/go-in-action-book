package main

import (
	"bytes"
	"html/template"
	"testing"
)

// execute with:
// go test -bench .

func BenchmarkTemplates(b *testing.B) {
	b.Logf("b.N is %d\n", b.N)
	tpl := "Hello {{.Name}}"
	data := &map[string]string{
		"Name": "World",
	}

	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		t, _ := template.New("test").Parse(tpl)
		t.Execute(&buf, data)
		buf.Reset() // clears the buffer to avoid memory-allocation issues
	}
}

func BenchmarkCompiledTemplates(b *testing.B) {
	b.Logf("b.N is %d\n", b.N)
	tpl := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl) // compiled just once
	data := &map[string]string{
		"Name": "World",
	}

	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		t.Execute(&buf, data)
		buf.Reset() // clears the buffer to avoid memory-allocation issues
	}
}

// go test -bench . -cpu=1,2,4
func BenchmarkParallelTemplates(b *testing.B) {
	b.Logf("b.N is %d\n", b.N)
	tpl := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl)
	data := &map[string]string{
		"Name": "World",
	}

	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			t.Execute(&buf, data)
			buf.Reset()
		}
	})
}

// go test -bench Race -race -cpu=1,2,4
func BenchmarkParallelWithRaceConditon(b *testing.B) {
	// b.Logf("b.N is %d\n", b.N)
	tpl := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl)
	data := &map[string]string{
		"Name": "World",
	}

	var buf bytes.Buffer
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t.Execute(&buf, data)
			buf.Reset()
		}
	})
}
