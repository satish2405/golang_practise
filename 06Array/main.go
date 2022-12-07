package main

import "fmt"

func main() {

	fmt.Println("welcome to array in golang")
	var fruit = [4]string{"apple", "banana", "orange", "poppaya"}
	fmt.Println("friuts..: ", fruit)

	fruitlist := append(fruit[:2], "guava")
	fmt.Println("fruit list ..:", fruitlist)
	fmt.Printf("frust list..: %T\n", fruitlist)

	highscrore := make([]int, 2)

	highscrore[1] = 100
	fmt.Println("slice value..: ", highscrore)
}
