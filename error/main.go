package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
)

func main() {
	i, err := anyFunc(10)
	fmt.Println(i, err)

	i, err = anyFunc(100)
	if err != nil {
		log.Printf("err isn't nil\n")
	}
	fmt.Println(i, err)
}

func anyFunc(i int) (int, error) {
	if i > 10 {
		return i, errors.New("i should not be bigger than 10")
	}

	return i, nil
}
