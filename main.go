package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/atotto/clipboard"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cpn <file>")
		os.Exit(1)
	}

	// Lấy basename (tên file, bỏ đường dẫn)
	filename := filepath.Base(os.Args[1])

	// Copy vào clipboard
	err := clipboard.WriteAll(filename)
	if err != nil {
		fmt.Println("Error copying to clipboard:", err)
		os.Exit(1)
	}

	fmt.Println("Copied:", filename)
}
