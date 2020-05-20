package utils

import "strings"

// JoinWihtCommas join the phrase 
func JoinWihtCommas(phrases []string) string{
	result:= strings.Join(phrases[:len(phrases)-1], ", ")
	result += " and "
	result += phrases[len(phrases)-1]
	return result;
}