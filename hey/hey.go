package main

import (
	"fmt"
	"log"

	"github.com/Zheng7426/Golang-practice/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{
		"荷马",
		"曾子",
		"庄生",
		"苏格拉底",
	}

	message, err := greetings.Heys(names)
	//message, err := greetings.Hey("")
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
