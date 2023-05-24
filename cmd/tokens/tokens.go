package tokens

import (
	"regexp"
)

func isId(str string) bool {
	match, _ := regexp.MatchString(`a-zA-Z`, str)

	return match
}
func isOperator(str string) bool {
	match, _ := regexp.MatchString(`[+\-*/]`, str)
	return match
}
func isNoDerivable(str string) bool {
	match, _ := regexp.MatchString(`[\(\)\;\}\{]`, str)
	return match
}
func isComparable(str string) bool {
	match, _ := regexp.MatchString(`<=|>=|==|!=|=`, str)
	return match
}
func isLogical(str string) bool {
	match, _ := regexp.MatchString(`&|\|`, str)
	return match
}

func isInteger(str string) bool {
	match, _ := regexp.MatchString(`\b\d+\b`, str)
	return match
}
func isFloat(str string) bool {
	match, _ := regexp.MatchString(`\b\d+\.\d+\b`, str)
	return match
}
func isReserved(str string) bool {
	match, _ := regexp.MatchString(`\b(for|while|do|return|int|float|if|else)\b`, str)
	return match
}

func isIncrementOrDecrement(str string) bool {
	match, _ := regexp.MatchString(`--|\+\+`, str)
	return match
}
func isWhiteSpace(str string) bool {
	for _, char := range str {
		if char != ' ' && char != '\t' && char != '\n' && char != '\r' {
			return false
		}
	}
	return true
}

func TokenMaker(str string) []string {
	var tokens []string
	var stack string
	for _, char := range str {
		switch {
		case isWhiteSpace(string(char)):
			if stack != "" {
				tokens = append(tokens, stack)
				stack = ""
			}

		case isNoDerivable(string(char)):
			if stack != "" {
				tokens = append(tokens, stack)
				stack = ""
			}
			tokens = append(tokens, string(char))
		case isOperator(string(char)):
			if stack == string(char) {
				tokens = append(tokens, stack+string(char))
				stack = ""
			} else {
				if stack != "" {
					tokens = append(tokens, stack)
				}
				stack = string(char)
			}

		default:
			stack += string(char)
		}

	}
	return tokens
}

func TokenFormatter(tokens []string) []string {
	var word []string
	for _, tok := range tokens {
		if tok != "" {
			if isComparable(tok) || isLogical(tok) || isNoDerivable(tok) || isOperator(tok) || isReserved(tok) || isIncrementOrDecrement(tok) {
				word = append(word, "<"+tok+">")
			} else if isFloat(tok) {
				word = append(word, "<"+tok+", float>")
			} else if isInteger(tok) {
				word = append(word, "<"+tok+", int>")
			} else {
				word = append(word, "<"+tok+", id>")
			}
		}
	}

	return word
}

// improvement
