package main

import (
	"fmt"
	"net/url"
)

const aurl string = "https://satish:3000"

func main() {
	fmt.Println("Learning URL's handling")
	fmt.Println("my url..:", aurl)

	//parsing the url
	result, err := url.Parse(aurl)
	if err != nil {
		fmt.Println("error parsing the url..:", err)
	}
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.User)
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Port())
	quearyparams := result.Query()
	fmt.Println(quearyparams["coursename"])
	for _, v := range quearyparams {
		fmt.Println("params are ..:", v)
	}

	// Reverse construct url using this chunks
	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "learngo.dev",
		Path:    "/here",
		RawPath: "user=SatishKumar",
	}
	anotherUrl := partsofUrl.String()
	fmt.Println("another url..:", anotherUrl)
}
