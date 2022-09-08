package main // é utilizado quando existe uma grande filas de tarefa para ser executada, o objetivo é deixar essas tarefas mais rápidas na medida do possível
//apesar dessa ferramenta melhorar muito na execução, ainda sim ficamos limitado pela potência da máquina, então cada uma executa de maneira mais rápida,
//de acordo com sua potência.
import "fmt"

func fibonacci(posicao int) int {// essa é uma função recursiva, então ela vai fazendo um montante de execuções e pode causar o estouro de pilha
	if posicao <= 1 { 
		return posicao
	}
	return fibonacci(posicao -2) + fibonacci(posicao -1)
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas{
		resultados <- fibonacci(numero)
	}
}

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados) //quanto mais eu chamar o go worker mais rápido o programa é executado.
	go worker(tarefas, resultados)//aqui eu estou usando a concorrência e tem 4 worker puxando a execução.
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++{
		resultado := <-resultados
		fmt.Println(resultado)
	}

}