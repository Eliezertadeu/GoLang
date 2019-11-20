package prefixo

import "strconv"

//Capital representa o nº do prefixo do telefone estado/provincia
var Capital = 11

var teste = "teste"

//TesteComPrefixo teste + pref
var TesteComPrefixo = teste + " " + strconv.Itoa(Capital)
