package main

import "fmt"

func main() {
	fmt.Println("...Variables..")
	var username string = "satish kumar"
	fmt.Println(username)
	fmt.Printf("variable type : %T\n", username)

	var small byte = 255
	fmt.Printf("small type : %T and value : %d \n", small, small)

	// implicit type
	var web = "https://hello"
	fmt.Println(web)

	//no var style
	x := 100 //volus operator
	fmt.Println(x)
}
