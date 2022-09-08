package main //FUNÇÕES VARIÁTICAS = é uma função que pode receber N parâmetros, não precisa especificar a quantidade de parâmetros que ela vai receber

import (
	"fmt"
)

func soma(numeros...int){ //aqui estou passando que minha função vai receber de 0 a N numeros inteiros.
	fmt.Println(numeros)
} 

/* Usado as funções variáticas eu posso iterar um função, então eu pego o valor que é passado entre (), somo e peço o total dessa soma*/

func calculos(num...int)int{ // a func variática está
	total := 0
	for _, num := range num{
		total += num
	}
	return total
}

func mista(texto string, n...int) { // também é possível criar funções variáticas parâmetros misturados, string e inteiros, nesse caso o n é um variático.
	for _, n := range n { //a limitação é que não podemos ter mais de um parâmetro variático por função e a func variática tem que ser sempre colocada por último.
		fmt.Println(texto, n)
	}
}

func main() {
	fmt.Println("Funções variáticas.")
	soma (1, 2, 3, 4, 5, 6)
	fmt.Println(soma)

	totalDaSoma := calculos(10, 15, 20) // criei a func totalDaSoma que recebe os parâmetros da func Calculos onde passei os valores 10, 15, 20
	fmt.Println(totalDaSoma)

	mista("Olá mundo", 13, 14, 15)
	fmt.Println(mista)
}
/* Usado as funções variáticas eu posso iterar um função*/