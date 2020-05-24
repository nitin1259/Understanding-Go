package main

import (
	"fmt"

	"github.com/nitin1259/Understanding-Go/utils"
)

func main() {
	fmt.Println("code quality assurance: Automated Testing")

	res:= utils.JoinWihtCommas([]string{"apple", "orange", "pears", "banana"})
	fmt.Println(res)

}