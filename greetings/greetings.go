package greetings 

import "fmt"
import "errors"

func Hey(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("嘿, %v. 欢迎!", name)
	return message, nil 
}

