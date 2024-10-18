package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco (R$)"`
	Tags  []string `json:"tags"`
}

func main() {
	//Struct para json
	p1 := produto{1, "Notebook", 1899.90, []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	//Json para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"Caneta Azul","preco (R$)":1.50,"tags":["Papelaria","Importando"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
}