package matematica

// Calculo de numeros
func Calculo(funcao func(int, int) int, x int, y int) int {
	return funcao(x, y)
}

//Mult dois numeros
func Mult(x int, y int) int {
	return x * y
}
