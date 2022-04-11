package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type TemplateFile struct {
	name     string
	contents string
}

func CreateTestDir(files []TemplateFile) string {
	dir, err := os.MkdirTemp("", "template")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		f, err := os.Create(filepath.Join(dir, file.name))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err = io.WriteString(f, file.contents)
		if err != nil {
			log.Fatal(err)
		}
	}
	return dir
}

func main() {
	dir := CreateTestDir([]TemplateFile{
		{"T0.tmpl", `T0 invokes T1: ({{template "T1"}})`},
		{"T1.tmpl", `{{define "T1"}}T1 invokes T2: ({{template "T2"}}){{end}}`},
		{"T2.tmpl", `{{define "T2"}}This is T2{{end}}`},
	})
	defer os.RemoveAll(dir)

	pattern := filepath.Join(dir, "*.tmpl")

	tmpl := template.Must(template.ParseGlob(pattern))

	err := tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}
