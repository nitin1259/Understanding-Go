package utils

import "strings"

// JoinWihtCommas join the phrase 
func JoinWihtCommas(phrases []string) string{
	if len(phrases) ==1{
		return phrases[0]
	}
	if len(phrases) ==2{
		return phrases[0]+ " and " + phrases[1]
	}
	
	result:= strings.Join(phrases[:len(phrases)-1], ", ")
	result += ", and "
	result += phrases[len(phrases)-1]
	return result;

	

}