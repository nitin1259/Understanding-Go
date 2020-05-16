package main

import (
	"fmt"
	"log"
)

func main() {

	// format output wiht Printf and Sprintf
	// fmt.Printf("one third equal to %0.2f\n", 1.0/3.0)

	//errors
	/* err := errors.New("height can't be negative")
	fmt.Println(err.Error())
	// or
	// err := fmt.Errorf("height can't be negative %0.2f ", -23.0)
	fmt.Println(err)
	log.Fatal(err)
	*/

	paintReq, err := paintNeededForWall(4.0, -3.2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Paint requried: ", paintReq)

}

func paintNeededForWall(width float64, height float64) (float64, error) {
	if width < 0 || height < 0 {
		return 0, fmt.Errorf("Width: %0.2f or height: %0.2f parameter cannot be negative", width, height)
	}
	area := width * height
	return area / 10.0, nil
}
