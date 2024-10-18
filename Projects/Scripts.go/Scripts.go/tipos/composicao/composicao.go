package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	//pode adicionar mais métodos
}

type bmw struct{}

func (b bmw) ligarTurbo() {
	fmt.Println("Turbo...")
}
func (b bmw) fazerBaliza() {
	fmt.Println("Baliza...")
}

func main()  {
	var b esportivoLuxuoso = bmw{}
	b.fazerBaliza()
	b.ligarTurbo()
}