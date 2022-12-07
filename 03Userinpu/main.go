package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "welecome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating of pizza")

	//comma ok || comma err syntax

	input, _ := reader.ReadString('\n')

	fmt.Println("thanks your rating is ..: ", input)
	fmt.Printf("Ratings type is ..: %T\n", input)

	numRating, err := strconv.ParseInt(strings.TrimSpace(input), 0, 64)

	if err != nil {
		fmt.Println("some error occured while conversion", err)
	} else {
		//numRating = numRating + 1
		fmt.Println("Rating added by 1 : ", numRating+1)
	}
}
