package main
//generator é um padrão utilizado para encapsular uma ou mais go routine

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá Mundo!") //esse aqui é um canal que só recebe

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <- chan string { //a func escrever vai retornar pra gente um canal de string
	canal := make(chan string)

	go func() {
		canal <- fmt.Sprintf("Valor recebido %s", texto)
		time.Sleep(time.Millisecond * 500)
	}()
//aqui eu estou encapsulando a clausula go e quando eu for chamar ela na func main eu posso fazer normalmente, como uma função. 
//em uma go routine grande, ou uma série de go routine, esse encapsulamento faz todo o sentido, pois eu abstraio toda essa parte.
	return canal
}