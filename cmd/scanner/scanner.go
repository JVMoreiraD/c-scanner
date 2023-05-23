package scanner

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/JVMoreiraD/c-scanner/cmd/tokens"
)

func Scanner(filePath string) {
	readFile, err := os.Open(filePath)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	var word []string
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanWords)

	for fileScanner.Scan() {
		wr := fileScanner.Text()
		tokenRegex := regexp.MustCompile(`(\+\+|--)|!=|==|>=|<=|[(){;}\\=]+|\b\w+\b`)
		// fmt.Println(wr)
		tokensText := tokenRegex.FindAllString(wr, -1)
		fmt.Println(tokensText)

		// test := tokens.TokenMaker(wr)
		// fmt.Println(test)
		for _, tok := range tokensText {
			if tok != "" {
				if tokens.IsComparable(tok) || tokens.IsLogical(tok) || tokens.IsNoDerivable(tok) || tokens.IsOperator(tok) || tokens.IsReserved(tok) || tokens.IsIncrementOrDecrement(tok) {
					word = append(word, "<"+tok+">")
				} else if tokens.IsFloat(tok) {
					word = append(word, "<"+tok+", float>")
				} else if tokens.IsInteger(tok) {
					word = append(word, "<"+tok+", int>")
				} else {
					word = append(word, "<"+tok+", id>")
				}
			}
		}
	}
	word = append(word, "<eof>")

	fmt.Printf("%v\n", word)
}
