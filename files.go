package main

import (
	"fmt"
	"os"
)

func CreateFiles(dirBase string, files ...string) {
	for _, name := range files {
		fileLocation := fmt.Sprintf("%s%s.%s", dirBase, name, "txt")
		file, err := os.Create(fileLocation)

		defer file.Close()

		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", name, err)
			os.Exit(1)
		}

		fmt.Printf("File %s has been created.\n", file.Name())
	}
}

func main() {
	tmp := os.TempDir()

	CreateFiles(tmp)
	CreateFiles(tmp, "file1")
	CreateFiles(tmp, "file2", "file3", "file4")
}
