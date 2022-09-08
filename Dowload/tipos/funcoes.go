package main /*funções em go também são consideradas um tipo, onde podemos passar parâmetros, esperar retornos e aninhar 
uma na outra*/	

/*as funções são como receitas de bolo, elas contém o passo a passo do que o seu programa pode fazer*/
import (
	"fmt"
)
func somar(n1, n2 int8) int8{ //aqui estou definindo q minha func possui 2 numeros e retorna a soma entre eles.
	return n1 + n2
}
	//em GO as funções podem ter mais de um retorno, para isso basta utilizar os ()
func Calculos(n1, n2 int8) (int8, int8) { //estou dizendo q esta minha função recebe 2 números como parâmetros e pode me retornar 2 números.
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 25) // essa é uma forma simples de executar uma função e passar seus parâmetros. 
	fmt.Println(soma)

 	/*também é possível que o parâmetro de uma funcão seja outra função ou o seu retorno tbm seja outra função.
 	var f = func() //a função é um tipo, eu passei ele para a variável "f" 
		fmt.Println("função f")
	 f()*/

	 resultadoSoma, resultadoSubtracao := Calculos(25, 30)
	 fmt.Println(resultadoSoma, resultadoSubtracao)

}