package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("labeling data: Maps");
	
	var mymap map[string]int
	fmt.Printf("mymap: %#v, type: %s \n", mymap, reflect.TypeOf(mymap))

	mymap = make(map[string]int) // this actuall creates the map

	ranks := make(map[string]int) // short hand create a map and declare a varable to hold it
	ranks["gold"]=1
	ranks["silver"]=2
	ranks["bronze"]=3
	fmt.Println(ranks["gold"])
	fmt.Println(ranks)

	//map literals
	elements:= map[string]string{"Li": "Lithium", "H": "Hydrogen"}
	fmt.Println(elements);

	// As with arrays and slices, if you access a map key that hasn’t been assigned to, you’ll get a zero value back
	// The zero value for a map variable is nil

	// How to tell zero values apart from assigned values

	goldValue, ok := ranks["gold"];
	fmt.Println(goldValue, ok)
	diamondValue, ok := ranks["diamond"] // diamond is not a key so how to decide it there or not as it will return default
	if ok{
		fmt.Println(diamondValue, ok)
	}

	someSlice:= []string{"a", "e", "c", "e", "a"}
	countMap := make(map[string]int)
	for _, item:= range someSlice{
		countMap[item]++;
	}
	someSlice2:= []string{"a", "b", "c", "d", "e"}
	for _,item:= range someSlice2{
		value, ok:= countMap[item]
		if !ok{
			fmt.Printf("%s not found \n", item)
		}
		if ok{
			fmt.Printf("%s found with count: %d \n", item, value)
		}
	}

	grades:= map[string]int{"Nitin": 61, "Sachin": 84, "Vipin": 76}

	for key, value:= range grades{
		fmt.Printf("%s: %d \n", key, value)
	}
}