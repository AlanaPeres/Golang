package main /*Canais é a forma mais comum de organizar uma goroutine, o nome canais vem de comunicação, onde podemos enviar e receber dados
Essas duas operações de enviar e receber dados bloqueiam a execução do programa, então enquanto o programa não recebe o dado o programa fica parado*/

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)// é assim que chamamos um canal. chan é a palavra chave de canal e especificamos um tipo desse canal
	go escrever("Olá Mundo", canal)		//então ele só vai poder trafegar dados desse tipo, nesse caso aqui são dados de string
	
	fmt.Println("Depois da função escrever começa a ser executada")
	
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close (canal) //go tem essa proprimedade de verificar se o canal está aberto ou fechado, se estiver fechado não vai mais enviar nem receber dados

}