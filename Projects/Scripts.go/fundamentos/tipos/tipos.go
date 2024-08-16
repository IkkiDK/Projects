package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//Números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//Sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	//Com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do uint64 é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' //Representa mapeamento da tabela unicode(int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println("O valor do i2 é", i2)

	//Números reais (float 32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo do x é", reflect.TypeOf(x))
	fmt.Println("O tipo do x Literal é", reflect.TypeOf(49.99)) //Float64

	//Boolean
	bo := true
	fmt.Println("O tipo do bo é", reflect.TypeOf(bo))
	fmt.Println(!bo) //! antes inverte o valor do bo

	//String
	s1 := "Olá meu nome é Henrique"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	//String com múltiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Henrique`
	fmt.Println("O tamanho da string é", len(s2))

	//Char???
	char := 'a'
	fmt.Println("o tipo do char é", reflect.TypeOf(char))
	fmt.Println(char)
	//Char é um int32
}