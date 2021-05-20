package greetings 

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

func Hey(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	//message := fmt.Sprintf("嘿, %v. 欢迎!", name)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil 
}

func Heys(names []string) (map[string]string, error) {
	// 将名字与信息联系起来的映射
	messages := make(map[string]string)

	// 循环名字的 slice，对每一个名字调用 Hey 函数
	for _, name := range names {
		message, err := Hey(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"嘿, %v. 欢迎!",
		"见到你真棒, %v!",
		"哟, %v! 很稳!",
	}

	return formats[rand.Intn(len(formats))]
}
