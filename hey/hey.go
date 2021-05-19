package main

import (
	"fmt"
	"log"

	"github.com/Zheng7426/Golang-practice/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//message, err := greetings.Hey("荷马")
	message, err := greetings.Hey("")
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
