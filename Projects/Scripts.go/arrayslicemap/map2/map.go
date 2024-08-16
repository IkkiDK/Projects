package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Henrique": 1000.,
		"Leonardo": 1200.,
		"Mário":    1500.,
	}
	funcsESalarios["Josnei"] = 69.0
	delete(funcsESalarios, "Arnaldo")
	for nome, salario := range funcsESalarios {
		fmt.Printf("%s -> salário (%.2f$)\n", nome, salario)
	}
}
