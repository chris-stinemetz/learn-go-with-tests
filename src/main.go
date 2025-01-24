package main

import (
	"fmt"

	"github.com/chris-stinemetz/learn-go-with-tests/src/pkg/integers"
)

func main() {
	result := integers.Add(2, 4)
	fmt.Println(result)
}
