package main /*FUNÇÕES RECURSIVAS SÃO FUNÇÕES QUE CHAMAM ELAS MESMAS direta ou indiretamente, para que sua execução aconteça ela precisa de uma outra execução dela mesma*/
	//Não é recomendada para funções muito grandes pois ele empilha muitas execuções.

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 { //aqui eu pedi para minha função que se o número da posição for menou ou igual a zero ele me retorna o próprio número da posição.
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
func main() {
	fmt.Println("Funções Recursivas")

	//1 1 2 3 5 8 13

	posicao := uint(15)//nesta função eu estou pedindo para q o programa me retorne o número que está na posição x
	fmt.Println(fibonacci(posicao))

	//eu posso fazer também toda a sequencia de fibonnaci e não apenas uma posição.
	for i := uint(0); i < posicao; i++{ //i recebe um número inteiro, enquanto i for menor que a posição, i é incrementado
		fmt.Println(fibonacci(i))
	}
}