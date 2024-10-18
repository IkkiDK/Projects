package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //Array
	s1 := []int{1, 2, 3}  //Slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))
	//Slice não é um array! Slice define um pedaço do array

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[:3]  //Começa no 0 e não considera o índice 3
	s3 := a2[3:5] //Um novo slice apontando para o mesmo array e não considera o índice 5
	fmt.Println(s2, s3)

	//Considerações: slice tem tamanho e um ponteiro para um elemento do array 
	s4 := s2[:2] //Slice de slice
	fmt.Println(s4)
}
