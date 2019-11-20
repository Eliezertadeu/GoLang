package main

import (
	"fmt"

	"github.com/Eliezertadeu/funcoes_basicas/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Mult, 2, 2)
	fmt.Printf("Resultado: %d\r\n", resultado)
	resultado = matematica.Soma(4, 6)
	fmt.Printf("Resultado: %d\r\n", resultado)
}
