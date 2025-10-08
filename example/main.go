package main

import (
	"fmt"

	"github.com/nxtgo/scrapper"
)

func main() {
	elements, err := scrapper.ScrapByURL("https://github.com/nxtgo", "h1[dir=\"auto\"]")
	if err != nil {
		panic("do something")
	}

	for _, element := range elements {
		fmt.Println(element.Value)
	}
}
