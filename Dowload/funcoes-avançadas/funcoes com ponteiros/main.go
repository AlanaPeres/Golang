package main
/*Ponteiro é o endereço de uma varíavel, é o local em que o valor é armazenado. Nem todo o valor tem um endereço, mas toda variável tam. Com o ponteiro podemos
podemos ler ou utilizar o valor de uma variável indiretamente, sem usar ou sequer saber o nome da variável, se é que ela tem um nome. */

import (
	"fmt"
)

func inverterSinal(numero int) int{ //AQUI EU ESTOU PASSANDO UM PARÂMETRO POR CÓPIA, então quando durante o trabalho eu quiser utilizar uma variável sem alterar seu valor 
	return numero * -1              //em todo o programa eu posso utilizar esse modelo de cópia.
}

func inverterComPonteiro(numero *int) { //AQUI EU ESTOU ALTERANDO O VALOR DE UMA VARIÁVEL UTILIZANDO O PONTEIRO.
	*numero = *numero * -1
}

func main() {//aqui nessa primeira função o programa não altera o valor dentro da variável, ele apenas cria uma cópia dessa variável e faz a alteração.
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	//utilizando ponteiro para inverter a info dentro da variável.
	NovoNumero := 50
	fmt.Println(NovoNumero)
	inverterComPonteiro(&NovoNumero)// aqui eu passei o endereço de memória onde esta variável e fiz a alteração dentro deste endereço.
	fmt.Println(NovoNumero)
}