package main

import "math"

//Iniciando com letra maiúscula é PÚBLICO (visível dentro e fora do package)!
//Iniciando com letra minúscula é PRIVADO (visível apenas dentro do package)!

//Ponto - gerará algo público (maiúsculo)
//ponto ou _Ponto - gerará algo privado (minúsculo)

//Ponto representa uma cordenada no plano cartesiano

type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64){
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return 
}

func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}