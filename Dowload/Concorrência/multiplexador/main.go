package main // o padrão multiplexador é pegar dois ou mais canais e juntar em um só.


import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	canal := multiplexar(escrever("olá mundo!"), escrever ("Programando em GO."))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canalDeEntrada1, canalDeEntrada2 <- chan string) <- chan string { // nessa func estamos recebendo 2 canais, mas poderiam ser n..
	canalDeSaída := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalDeEntrada1:
				canalDeSaída <- mensagem
			case mensagem := <- canalDeEntrada2:
				canalDeSaída <- mensagem
			}
		}
	}()
	return canalDeSaída 
}
 func escrever(texto string) <- chan string { //a func escrever vai retornar pra gente um canal de string
	canal := make(chan string)

	go func() {
		for{
		canal <- fmt.Sprintf("Valor recebido %s", texto)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}