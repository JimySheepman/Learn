package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	dir1 := CreateTestDir([]TemplateFile{
		{"T1.tmpl", `T1 invokes T2: ({{template "T2"}})`},
	})

	dir2 := CreateTestDir([]TemplateFile{
		{"T2.tmpl", `{{define "T2"}}This is T2{{end}}`},
	})

	defer func(dirs ...string) {
		for _, dir := range dirs {
			os.RemoveAll(dir)
		}
	}(dir1, dir2)

	// Here starts the example proper.
	// Let's just parse only dir1/T0 and dir2/T2
	paths := []string{
		filepath.Join(dir1, "T1.tmpl"),
		filepath.Join(dir2, "T2.tmpl"),
	}
	tmpl := template.Must(template.ParseFiles(paths...))

	err := tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}
