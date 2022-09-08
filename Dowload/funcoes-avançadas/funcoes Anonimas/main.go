package main //FUNÇÕES ANONIMAS - Recursão é um processo no qual uma função chama a si mesma implícita ou explicitamente e a função correspondente é 
//chamada de função recursiva. É uma função que não contém nenhum nome.

import (
	"fmt"
)

func main() { // A função anonima chama a si mesma
	func() { 
		fmt.Println("Olá mundo")
	}()// para chamar uma func anonima basta colocar esses parênteses no final. 

	func (texto string) {//é possível passar parâmetros.
		fmt.Println(texto)
	}("passando parâmetro")//caso a função tenha algum parâmetro é só colocar aqui dentro. 
}