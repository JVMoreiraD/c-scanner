package tokens

import (
	"strconv"
	"strings"
)

func isValid(str string) bool {

	for _, char := range str {
		if strings.ContainsAny(string(char), "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			return true
		}
	}
	return false
}

func isOperator(str string) bool {

	for _, char := range str {
		if char != '+' && char != '-' && char != '*' && char != '/' {
			return false
		}

	}
	return len(str) > 0
}

func isWhiteSpace(str string) bool {
	for _, char := range str {
		if char != ' ' && char != '\t' && char != '\n' && char != '\r' {
			return false
		}
	}
	return len(str) > 0
}
func isNoDerivable(str string) bool {
	for _, char := range str {
		if char != ')' && char != '(' && char != '}' && char != '{' && char != ';' {
			return false
		}
	}
	return len(str) > 0
}

func isLogical(str string) bool {
	for _, char := range str {
		if char != '&' && char != '|' {
			return false
		}
	}
	return len(str) > 0
}

func isComparable(str string) bool {

	operators := []string{"<=", ">=", "==", "!=", "=", ">", "<"}

	for _, op := range operators {
		if str == op {
			return true
		}
	}

	return false
}

func isInteger(str string) bool {

	_, err := strconv.Atoi(str)
	return err == nil
}

func isFloat(str string) bool {

	sides := strings.Split(str, ".")
	if len(sides) != 2 {
		return false
	}
	rightSide := sides[0]
	leftSide := sides[1]
	_, err := strconv.Atoi(rightSide)
	if err != nil {
		return false
	}
	_, err = strconv.Atoi(leftSide)
	if err != nil {
		return false
	}
	return true

}

func isReserved(str string) bool {
	operators := []string{"for", "while", "do", "return", "int", "float", "if", "else"}

	for _, op := range operators {
		if str == op {
			return true
		}
	}

	return false
}

func TokenMaker(str string) []string {
	var tokens []string
	var stack string
	for _, char := range str {
		switch {
		case isLogical(string(char)):
			if stack != "" {
				if !isReserved(stack) {
					tokens = append(tokens, tokenParser(stack))
					stack = string(char)
				} else {
					stack += string(char)
				}
			} else {
				stack = string(char)
			}
		case isWhiteSpace(string(char)):
			if stack != "" {
				tokens = append(tokens, tokenParser(stack))
				stack = ""
			}
		case string(char) == ".":
			if stack != "" {
				if isInteger(stack) {
					stack += string(char)
				} else {
					tokens = append(tokens, tokenParser(stack))
					stack = ""
				}
			} else {
				stack += string(char)
			}
		case isInteger(string(char)):
			if stack != "" {
				if isFloat(stack + string(char)) {
					stack += string(char)
				} else if !isInteger(stack) {
					tokens = append(tokens, tokenParser(stack))
					stack = string(char)
				} else {
					stack += string(char)
				}
			} else {
				stack = string(char)
			}

		case isNoDerivable(string(char)):
			if stack != "" {
				tokens = append(tokens, tokenParser(stack))
				stack = ""
			}
			tokens = append(tokens, tokenParser(string(char)))

		case isOperator(string(char)):
			if stack == string(char) {
				tokens = append(tokens, tokenParser(stack+string(char)))
				stack = ""
			} else {
				if stack != "" {
					tokens = append(tokens, tokenParser(stack))
				}
				stack = string(char)
			}
		case isComparable(string(char)):
			if stack != "" {
				if !isComparable(stack) {
					tokens = append(tokens, tokenParser(stack))
					stack = string(char)
				} else {
					stack += string(char)
				}
			} else {
				stack = string(char)
			}

		case isValid(string(char)):
			if stack != "" {
				if !isValid(stack) {
					tokens = append(tokens, tokenParser(stack))
					stack = string(char)
				} else {
					stack += string(char)
				}
			} else {
				stack = string(char)
			}
		default:
			stack += string(char)
		}

	}
	if stack != "" {
		tokens = append(tokens, tokenParser(stack))

	}
	return tokens
}

func tokenParser(str string) string {
	switch {
	case isValid(str) && !isReserved(str):
		str = "<" + str + ", id>"
	case isFloat(str):
		str = "<" + str + ", float>"
	case isInteger(str):
		str = "<" + str + ", int>"
	default:
		str = "<" + str + ">"

	}

	return str
}
