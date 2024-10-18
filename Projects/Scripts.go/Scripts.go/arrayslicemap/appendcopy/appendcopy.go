package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	//array1 = append(array1, 4, 5, 6)
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2, 10)
	copy(slice2, slice1) //Copy não pode passar array, somente slice, e não aumenta o tamanho caso não caiba
	fmt.Println(slice2)
}
