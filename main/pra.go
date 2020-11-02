package main

import (
	"fmt"

	"github.com/youk-h/practice/a"
)

func main() {
	fmt.Println(Pang(a.NewA(5)))
}

func Pang(a a.A) int {
	return a.Hoge(5)
}
