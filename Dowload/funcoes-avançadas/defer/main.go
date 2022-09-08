package main //DEFER = ADIAR esta clausula serve para adiar a execução de uma função.

import "fmt"

func func1(){
	fmt.Println("Executando a função 1")
}

func func2(){
	fmt.Println("Executando a função 2")
}

func main(){
	defer func1()//aqui eu estou pedindo para que o programa execute a func1 depois de todas as outras.
	func2()
}