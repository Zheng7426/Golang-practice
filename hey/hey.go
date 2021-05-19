package main

import (
	"fmt"
	"github.com/Zheng7426/Golang-practice/greetings"
)

func main() {
	message := greetings.Hey("荷马")
	fmt.Println(message)
}
