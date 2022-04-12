package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	parentDir := os.TempDir()
	logsDir, err := ioutil.TempDir(parentDir, "-logs")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(logsDir)

	globPattern := filepath.Join(parentDir, "*-logs")
	matches, err := filepath.Glob(globPattern)
	if err != nil {
		log.Fatalf("Failed to match %q: %v", globPattern, err)
	}

	for _, match := range matches {
		if err := os.RemoveAll(match); err != nil {
			log.Printf("Failed to remove %q: %v", match, err)
		}
	}
}
