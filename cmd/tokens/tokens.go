package tokens

import (
	"regexp"
)

var isId = regexp.MustCompile(`a-zA-Z`)

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
