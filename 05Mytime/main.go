package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now().Add(10)
	fmt.Println(currentTime.Format("01-02-2006 15:06:05 Monday"))
}
