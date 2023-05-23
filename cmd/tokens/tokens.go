package tokens

import (
	"regexp"
)

func IsId(str string) bool {
	match, _ := regexp.MatchString(`a-zA-Z`, str)

	return match
}
func IsOperator(str string) bool {
	match, _ := regexp.MatchString(`[+\-*/]`, str)
	return match
}
func IsNoDerivable(str string) bool {
	match, _ := regexp.MatchString(`[\(\)\{\};]`, str)
	return match
}
func IsComparable(str string) bool {
	match, _ := regexp.MatchString(`<=|>=|==|!=|=`, str)
	return match
}
func IsLogical(str string) bool {
	match, _ := regexp.MatchString(`&|\|`, str)
	return match
}

func IsInteger(str string) bool {
	match, _ := regexp.MatchString(`\b\d+\b`, str)
	return match
}
func IsFloat(str string) bool {
	match, _ := regexp.MatchString(`\b\d+\.\d+\b`, str)
	return match
}
func IsReserved(str string) bool {
	match, _ := regexp.MatchString(`\b(for|while|do|return|int|float|if|else)\b`, str)
	return match
}

func IsIncrementOrDecrement(str string) bool {
	match, _ := regexp.MatchString(`--|\+\+`, str)
	return match
}
func IsWhiteSpace(str string) bool {
	match, _ := regexp.MatchString(`\s`, str)
	return match
}

func TokenMaker(str string) []string {
	var tokens []string
	currentToken := ""
	if IsIncrementOrDecrement(str) {
		tokens = append(tokens, str)
	}
	for i := 0; i < len(str); i++ {
		if IsNoDerivable(string(str[i])) {
			if currentToken != "" {
				tokens = append(tokens, currentToken)
				currentToken = ""
			}
			tokens = append(tokens, string(str[i]))
		} else if IsOperator(string(str[i])) {
			if currentToken != "" {
				tokens = append(tokens, currentToken)
				currentToken = ""
			}
			tokens = append(tokens, string(str[i]))
		} else if IsWhiteSpace(string(str[i])) {
			if currentToken != "" {
				tokens = append(tokens, currentToken)
				currentToken = ""
			}
		} else {
			currentToken += string(string(str[i]))
		}
		if currentToken != "" {
			tokens = append(tokens, currentToken)
		}
	}
	return tokens
}

// improvement
