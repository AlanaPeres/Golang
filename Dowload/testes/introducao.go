package main

import (
	"introducao-testes/endereco" //modulo/pacote
	"fmt"
)

func main() {
	tiposEndereco := endereco.TipoDeEndereco("Avenida da Saudade.")
	fmt.Println(tiposEndereco)
}