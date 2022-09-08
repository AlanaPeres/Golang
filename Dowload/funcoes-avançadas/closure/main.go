package main 
/*Em Golang, um closure é uma função que referencia variáveis ​​fora de seu escopo(corpo). Um encerramento pode sobreviver ao escopo no qual foi criado.
 Assim, ele pode acessar variáveis ​​dentro desse escopo, mesmo depois que o escopo for destruído. */
import "fmt"

func closure() func() {//estou criando um func que se chama closure e ela vai me retornar outra função
	texto := "dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}
 /*Em Go, as funções podem retornar outras funções. Isso pode ser feito, por exemplo, retornando uma função anônima.

 Um closure é um valor de função que faz referência a variáveis ​​fora de seu próprio corpo de função. Uma closure é capaz de sobreviver ao escopo
 da função na qual ela é definida. Isso significa que ele pode acessar as variáveis ​​fora do escopo.*/

func main() {
	texto := "Dentro da func main" 
	fmt.Println(texto) 

	funcNova := closure()//quando eu chamo o closure o programa entende que estou chamando o texto de dentro da funcão q foi criada anteriormente.
	funcNova()
}