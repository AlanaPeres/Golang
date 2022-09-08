//É UMA DAS FORMAS DE SINCRONIZAR GORROTINAS

package main

import (
	"time"
	 "fmt"
	 "sync"
)

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2) // Aqui eu estou passando para o meu programa q ele vai ter 2 goroutines q ele vai precisar esperar terminar para encerrar o programa
//não precisa referenciar, só passar a quantidade
	go func() {// crio as func anônimas 
		escrever("Olá mundo!")	
		waitGroup.Done()//esse método vai tirar um do contador, então eu passai que seriam 2 goroutines com esse done ele vai -1

	}()
		//usando o go na frete eu faço essas duas funções serem executadas simultaneamente
	go func() {
		escrever("Programando em Go!")	
		waitGroup.Done()
	}()
	waitGroup.Wait() // aqui eu passei a ordem ao programa para ele esperar meu contador chegar a zero, então ele vai esperar as duas funções terminarem para ele encerrar
}