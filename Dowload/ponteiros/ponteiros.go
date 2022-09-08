package main
/* Uma variável é um item de armazenagem que contém um valor.
Um ponteiro é o endereço de uma variável. é o local em que o valor é armazenado. Nem todo valor tem um endereço, mas toda variavel 
tem. Como um ponteiro, podemos ler ou utilizar o valor de uma variável indiretamente, sem usar ou sequer saber o nome da 
variável, se é que ela tem nome.*/

import "fmt"

func main() {

	fmt.Println("Ponteiro é uma referencia de memória")

	var v1 int = 10 
	var v2 int = v1 // estou atribuindo a v2 o mesmo valor da v1
	fmt.Println(v1, v2)

	var v3 int = 100 //essa variável guarda um inteiro 
	var ponteiro = &v3 // essa variável guarda o endereço de memória onde a v3 está armazenada.
	fmt.Println(v3, *ponteiro, ponteiro)//quando eu coloco o * eu estou pedindo o valor que está armazenado neste endereço

}