package main

import (
	"fmt"
)

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) { //esse é o retorno nomeado. 
	soma = n1 + n2
	subtracao = n1 - n2
	return //aqui eu não preciso colocar return soma, subtração, pois seria como jogar fora  o retorno nomeado. Essa é uma forma de declarar uma função de forma mais bonita.
}

func main() {
	fmt.Println("Funções com retornos nomeados")

	soma, subrtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subrtracao)

}