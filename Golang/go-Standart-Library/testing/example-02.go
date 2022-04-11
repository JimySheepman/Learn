package main

import (
	"bytes"
	"testing"
	"text/template"
)

func main() {
	testing.Benchmark(func(b *testing.B) {
		templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
		b.RunParallel(func(pb *testing.PB) {
			var buf bytes.Buffer
			for pb.Next() {
				buf.Reset()
				templ.Execute(&buf, "World")
			}
		})
	})
}
