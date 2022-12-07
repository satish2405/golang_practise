package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome for rand")
	//rand.Seed(time.Now().UnixNano()) //entire algo is driven by seed
	//fmt.Print(rand.Intn(5) + 1)
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)
}
