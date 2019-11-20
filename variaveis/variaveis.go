package main

import "fmt"

var (
	//telefone é um valor importante e ter que ser publico - public
	telefone string

	//Endereco é um valor importante para nossa aula - private
	Endereco string

	qnt      int     // = 0
	comprou  bool    // = false
	valor    float64 // 0.00
	palavras rune
)

func main() {
	nome := "Eliézer"
	fmt.Println("Endereço: ", Endereco)
	fmt.Println("Telefone: ", telefone)
	fmt.Println("Quantidade: ", qnt)
	fmt.Println("Comprou: ", comprou)
	fmt.Println("Valor: ", valor)
	fmt.Println("Meu nome é", nome)
}
