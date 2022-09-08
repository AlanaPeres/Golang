package main /**/
/*Você pode ter uma condição if sem um else.*/


import (
	"fmt"
)
func main() {

	fmt.Println("Estrutura de controle")

	numero := 10

	if numero > 15 {
	fmt.Println(numero, " É maior que 15. ")
	} else {
		fmt.Println(numero, "É menor que 15")
	}
	//ESSE É O IF INIT
	if outroNumero := numero; outroNumero > 0 {//aqui eu estou declarando uma variável, atribuindo valor a ela e já avaliando se esse valor atende uma condição. 
		fmt.Println("Número é maior que zero")
	} else {
		fmt.Println("Numero é menor que zero")

		fmt.Println(outroNumero)//Quando eu crio uma variável usando o if init eu estou limitando ela a esse escopo, se eu chamar ela fora daqui o programa acusa que a var não foi declarada.
	}
	 

}