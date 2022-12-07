package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com/"

func main() {
	fmt.Println("Web requests basics..")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Println("data in the response code..:", response.StatusCode)
	//fmt.Printf("type of response..: %T", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("full data.:", string(databytes))

}
