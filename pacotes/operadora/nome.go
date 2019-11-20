package operadora

import (
	"strconv"

	"github.com/Eliezertadeu/pacotes/prefixo"
)

//NomeOperadora representa o nome da operadora
var NomeOperadora = "XPTO Telecom"

//PrefixoCapital prefixo mais nome da operadora
var PrefixoCapital = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora
