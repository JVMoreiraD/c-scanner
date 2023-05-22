package scanner

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/JVMoreiraD/c-scanner/cmd/tokens"
)

func Scanner() {
	fmt.Println("Insira as instruções (Digite 'fim' para encerrar):")

	reader := bufio.NewReader(os.Stdin)
	var word []string

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "fim" {
			break
		}
		// Regex pattern to match operators and separators
		tokenRegex := regexp.MustCompile(`!=|==|>=|<=|[(){;}]|[^\s(){};]+`)

		// Find all tokens in the input
		tokensText := tokenRegex.FindAllString(input, -1)
		for _, tok := range tokensText {
			if tok != "" {
				if tokens.IsComparable(tok) || tokens.IsLogical(tok) || tokens.IsNoDerivable(tok) || tokens.IsOperator(tok) || tokens.IsReserved(tok) {
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
