package main

import "fmt"

type ShoppingList []string

func (list ShoppingList) Categorize() ([]string, []string, []string) {
	var vegetables, meats, others []string

	for _, e := range list {
		switch e {
		case "Alface", "Tomate":
			vegetables = append(vegetables, e)
		case "Peixe", "Frango":
			meats = append(meats, e)
		default:
			others = append(others, e)
		}
	}

	return vegetables, meats, others
}

func main() {
	list := make(ShoppingList, 6)
	list[0] = "Alface"
	list[1] = "Tomate"
	list[2] = "Azeite"
	list[3] = "Peixe"
	list[4] = "Frango"
	list[5] = "Chocolate"

	fmt.Println(list)
	fmt.Println(list.Categorize())
}
