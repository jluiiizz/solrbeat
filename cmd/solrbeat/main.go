package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func read_files(folder string) []string {
	entries, err := os.ReadDir(folder)

	if err != nil {
		log.Fatal(err)
	}

	files := []string{}

	for _, entry := range entries {
		name := entry.Name()
		if strings.HasSuffix(name, ".log") {
			files = append(files, name)
		}
	}

	return files
}

func main() {
	fmt.Println("Starting solrbeat ...")
}
