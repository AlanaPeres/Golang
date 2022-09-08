package main 

/*Herança é um mecanismo que permite que características comuns a diversas classes sejam fatoradas em uma classe base, ou superclasse.
 A partir de uma classe base, outras classes podem ser especificadas*/

import (
	"fmt"
)
type pessoa struct {
	nome	string
	sobrenome	string
	idade	uint8
	altura	uint8
}

type estudante struct {
	pessoa //aqui eu busquei minha struct anterior e não especifiquei o tipo dela. Faz c/ q o programa não fique redundante e tenha q perguntar novamente as mesmas informações.
	curso	string
	faculdade string
}
func main() {
	fmt.Println("HERANÇA")

	p1 := pessoa{"João", "Pedro", 20, 178}
	fmt.Println(p1)

	e1 := estudante {p1, "Engenharia", "USP"}
	fmt.Println(e1) //se eu quiser pegar uma info isolada é só (e1.altura)
}
