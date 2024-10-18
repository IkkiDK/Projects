package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{"SF90", false, 0}
	carro1.ligarTurbo()

	//Outro modo de criar
	var carro2 esportivo = &ferrari{"SF90", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
