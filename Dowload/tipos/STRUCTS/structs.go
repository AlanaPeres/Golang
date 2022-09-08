package main //STRUCT é basicamente uma coleção de campos e cada campo possui um nome e um tipo.

import (
	"fmt"
)

type usuario struct { // estamos criando um tipo chamado usuário que vai armazenar uma strutura de dados
	nome string 
	idade uint8
	endereco endereco //esse é o tipo endereço que criei abaixo.
}

type endereco struct { //tbm é possível aninhar structs
	Logradouro string
	numero uint8
}

func main() { // a primeira forma de declarar uma struct é assim:
	var u usuario //para popular a struct fazemos:
	u.nome = "Alana"
	u.idade = 27
	fmt.Println(u)

	endereco1 := endereco{"Rua dos bobos", 15}

	//a segunda maneira de declarar a struct é utilizando a :=
	usuario2 := usuario {"Alana", 27, endereco1} 
	fmt.Println(usuario2)
}

