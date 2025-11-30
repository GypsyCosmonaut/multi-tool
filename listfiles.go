package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dirPath := "./"

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, entry := range entries {
		fullPath := filepath.Join(dirPath, entry.Name())
		fmt.Println(fullPath)
	}
}

