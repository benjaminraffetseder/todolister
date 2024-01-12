package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a folder to scan.")
		return
	}

	folderToScan := os.Args[1]

	// Create "stats" folder
	statsFolder := filepath.Join(folderToScan, "stats")
	os.MkdirAll(statsFolder, os.ModePerm)

	// Create CSV file in "stats" folder
	file, err := os.Create(filepath.Join(statsFolder, "todos.csv"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"File", "TODO", "Notes"}) // Added "Notes" column

	fmt.Println("Starting to scan files...")

	todoCount := 0

	err = filepath.Walk(folderToScan, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if info.Name() == "node_modules" || info.Name() == "stats" {
				return filepath.SkipDir
			}
		} else {
			ext := filepath.Ext(path)
			if ext == ".exe" || ext == ".bin" {
				return nil
			}

			fmt.Printf("Scanning file: %s\n", path)

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			reader := bufio.NewReader(file)
			lineNumber := 1
			for {
				line, err := reader.ReadString('\n')
				if err != nil {
					if err == io.EOF {
						break
					}
					return err
				}
				if strings.Contains(line, "TODO: ") {
					relativePath, _ := filepath.Rel(folderToScan, path)
					writer.Write([]string{relativePath, fmt.Sprintf("Line %d: %s", lineNumber, strings.TrimSpace(line)), ""}) // Added empty "Notes" cell
					todoCount++
				}
				lineNumber++
			}

			fmt.Printf("Finished scanning file: %s\n", path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Write total count
	writer.Write([]string{"Total", fmt.Sprintf("%d TODOs found", todoCount), ""}) // Added empty "Notes" cell

	fmt.Println("Finished scanning all files.")
}
