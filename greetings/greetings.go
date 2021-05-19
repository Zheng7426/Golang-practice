package greetings 

import "fmt"

func Hey(name string) string {
	message := fmt.Sprintf("嘿, %v. 欢迎!", name)
	return message 
}


