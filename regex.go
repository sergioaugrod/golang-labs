package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	regex := regexp.MustCompile("\\b\\w")

	transform := func(s string) string {
		return strings.ToUpper(s)
	}

	text := "golang clojure"

	fmt.Println(transform(text))
	fmt.Println(regex.ReplaceAllStringFunc(text, transform))
}
