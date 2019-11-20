package main

import (
	"fmt"

	"github.com/Eliezertadeu/pacotes/operadora"
	"github.com/Eliezertadeu/pacotes/prefixo"
)

//NomeDoUsuario é o nome do usuário do sistema
var NomeDoUsuario = "Eliézer"

func main() {
	fmt.Printf("Nome do usuário: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Valor do teste com pref: %s\r\n", prefixo.TesteComPrefixo)
}
