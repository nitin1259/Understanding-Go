package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("code quality assurance: Automated Testing")

	res:= joinWihtCommas([]string{"apple", "orange", "pears", "banana"})
	fmt.Println(res)

}


func joinWihtCommas(phrases []string) string{
	result:= strings.Join(phrases[:len(phrases)-1], ", ")
	result += " and "
	result += phrases[len(phrases)-1]
	return result;
}