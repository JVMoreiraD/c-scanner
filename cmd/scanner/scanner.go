package scanner

import (
	"bufio"
	"fmt"
	"os"

	"github.com/JVMoreiraD/c-scanner/cmd/tokens"
)

func Scanner(filePath string) {
	readFile, err := os.Open(filePath)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	var result []string
	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		words := tokens.TokenMaker(line)
		// fmt.Println(words)
		result = append(result, tokens.TokenFormatter(words)...)

	}
	result = append(result, "<eof>")
	fmt.Printf("%v\n", result)
}
