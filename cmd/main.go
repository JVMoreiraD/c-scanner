package main

import (
	"os"

	"github.com/JVMoreiraD/c-scanner/cmd/scanner"
)

func main() {
	filePath := os.Args[1]

	scanner.Scanner(filePath)
}
