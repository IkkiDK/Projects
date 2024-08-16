package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"H": {
			"Henrique": 1000.,
		},
		"J": {
			"Josnei": 69.,
		},
		"M": {
			"Mário": 500.,
		},
	}
	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}