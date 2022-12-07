package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"` // if value is nil dont show it at all
}

func main() {
	fmt.Println("welcome to json understanding")
	fullcoursesinfo := []course{
		{"Java", 100, "windows", "abc123", []string{"corejava", "advaced java"}},
		{"c", 200, "windows", "abc123", []string{"clang", "cpluscplus"}},
		{"python", 500, "windows", "abc123", nil},
	}
	finalJson, err := json.MarshalIndent(fullcoursesinfo, "", "\t")
	if err != nil {
		fmt.Println("Json error..:", err)
	}
	fmt.Printf("final json..: \n %s", finalJson)

}
