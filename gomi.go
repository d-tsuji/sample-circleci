package sample_circleci

import (
	"errors"
	"fmt"
)

func Addd(a, b int) int {
	{
		a := a
		fmt.Println(a)
	}
	return a + b
}

type Hoge struct{}

var X = 3

func unused() {
	if errors.New("abc") == errors.New("abc") {
		// test SA4000
	}
}
