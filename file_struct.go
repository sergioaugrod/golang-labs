package main

import "fmt"

type File struct {
	name   string
	lenght float64
	chars  int
	words  int
	lines  int
}

func main() {
	file := File{"article.txt", 12.68, 12986, 1862, 220}
	fmt.Println(file)
}
