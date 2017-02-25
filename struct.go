package main

import "fmt"

type State struct {
	name       string
	population int
	capital    string
}

func main() {
	states := make(map[string]State, 6)

	states["GO"] = State{"Goias", 6434052, "Goiania"}
	states["MG"] = State{"Minas Gerais", 5513131, "Belo Horizonte"}

	fmt.Println(states)

	_, found := states["SP"]

	if !found {
		fmt.Println("Doesnt exists this key")
	}

	for initials, state := range states {
		fmt.Printf("\nState(%s): %s Capital: %s Population: %d", initials, state.name, state.capital, state.population)
	}
}
