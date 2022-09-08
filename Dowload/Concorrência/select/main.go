package main // SELECT é utilizado apenas em concorrencias.
//quando eu executo essas 2 func anonimas sem usar o select eu acabo perdendo tempo de execução, pois uma fica pronta para executar a cada meio segundo
//e a outra fica pronta a cada 2 segundos. O select é utilizado para que cada função seja executada quando estiver pronta, independente das demais goroutines
import (
	"fmt"
	"time"
)
func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() { // essa é uma goroutine e func anonima
		for {
			time.Sleep(time.Millisecond * 500) // a cada meio segundo essa func deve ser executada
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for{
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()


	for { //aqui eu dei o comando para que quando a mensagem do canal tenha sido recebida, minha func execute o programa, independente do outro canal.
		select{//os canais ficam independentes e não esperam o outro para executar.
		case mensagem := <-canal1:
		fmt.Println(mensagem)
		
		case mensagem2 := <-canal2:
		fmt.Println(mensagem2)

		}
	}
}