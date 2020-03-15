package code

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

type Code int

var (
	messages = map[Code]string{}
	mu       = sync.RWMutex{}
)

func (c Code) Error() string {
	return string(c)
}

func Add(code int, msg string) (Code, error) {
	mu.Lock()
	defer mu.Unlock()
	c := Code(code)
	if _, ok := messages[c]; ok {
		return c, errors.New(fmt.Sprintf("code %v exists", code))
	}

	messages[c] = msg

	return c, nil
}

func MustAdd(code int, msg string) Code {
	c, err := Add(code, msg)
	if err != nil {
		log.Panic(err)
	}
	return c
}

func (c Code) Message(args ...interface{}) string {
	return fmt.Sprintf(messages[c], args...)
}
