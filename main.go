package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("./tests/mocks/mock1.json")
	if err != nil {
		log.Fatal(err)
	}

	result := Markdown(string(data))
	fmt.Println(result)
}
