package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to files..")
	content := "This need to go to a file.."

	file, err := os.Create("./mytest.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	len, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("length of lines..:", len)
	readfile("./mytest.txt")
}

func readfile(filname string) {
	databytes, err := ioutil.ReadFile(filname)

	if err != nil {
		fmt.Println("failed", err)
	}
	fmt.Println("text data in the file ", string(databytes))
}
