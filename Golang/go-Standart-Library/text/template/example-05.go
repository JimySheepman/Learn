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
		{"T1.tmpl", `{{define "T1"}}T1 invokes T2: ({{template "T2"}}){{end}}`},
		{"T2.tmpl", `{{define "T2"}}This is T2{{end}}`},
	})
	defer os.RemoveAll(dir)

	pattern := filepath.Join(dir, "*.tmpl")
	templates := template.Must(template.ParseGlob(pattern))
	_, err := templates.Parse("{{define `driver1`}}Driver 1 calls T1: ({{template `T1`}})\n{{end}}")
	if err != nil {
		log.Fatal("parsing driver1: ", err)
	}

	_, err = templates.Parse("{{define `driver2`}}Driver 2 calls T2: ({{template `T2`}})\n{{end}}")
	if err != nil {
		log.Fatal("parsing driver2: ", err)
	}

	err = templates.ExecuteTemplate(os.Stdout, "driver1", nil)
	if err != nil {
		log.Fatalf("driver1 execution: %s", err)
	}
	err = templates.ExecuteTemplate(os.Stdout, "driver2", nil)
	if err != nil {
		log.Fatalf("driver2 execution: %s", err)
	}
}
