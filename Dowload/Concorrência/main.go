/*CONCORRÊNCIA É O PRINCIPAL FORTE DA LINGUAGEM GO,
A concorrência é a composição de processos de execução independente, enquanto o paralelismo é o 
execução simultânea de cálculos (possivelmente relacionados). A concorrência é sobre lidar com muitos 
coisas de uma só vez. Paralelismo é sobre fazer muitas coisas ao mesmo tempo."*/
package main

import (
	"time"
	"fmt"
)

func escrever(texto string) {
	for{//esse é um laço infinito, se eu não interromper o programa ele vai ser executado para sempre.
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {
	go escrever("Olá mundo!") // quando eu coloco a palavra GO na frente da func eu estou dando o comando para que meu programa não espere terminar uma linha para executar outra. Para que ele faça isso de forma simultânea. 
	escrever("Programando em Go!")
}