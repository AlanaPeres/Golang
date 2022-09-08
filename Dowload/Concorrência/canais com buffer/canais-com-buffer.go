package main

import "fmt"

func main() {
	canal := make(chan string, 2)//aqui foi criando um canal que recebe duas mensagens, oque faz com q o meu programa não dê dedlock ao receber a primeira mensagem
	canal <- "Olá Mundo" //Então quando um canal possui um número na frente significa que a capacidade do canal está sendo definida com buffer
	canal <- "Programando em GO"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}