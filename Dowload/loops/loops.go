package main

/*Alem do fluxo de controle sequencial, há duas declarações que podem afetar como o computador lê o código:
    - Uma delas é o fluxo de controle de repetição (loop). Nesse caso, o computador vai repetir a leitura de um mesmo código de uma maneira específica.
	 O fluxo de controle de repetição tambem é conhecido como fluxo de controle iterativo, é possível ter um loop dentro de outro loop e outro loop(loop aninhado)...*/

import (
	"fmt"
	"time"
)

func main() { //essa é a forma mais comum e simples de utilizar a condição For. Enquanto uma condição for verdadeira, o programa repete o loop.
	/*Em sua forma mais simples, uma instrução "for" especifica a execução repetida de um bloco, desde que uma condição booleana seja avaliada como verdadeira. 
	A condição é avaliada antes de cada iteração. Se a condição estiver ausente, é equivalente ao valor booleano true.*/
	i := 0
	for i < 10{
		i++
		fmt.Println("Incrementando I")
		time.Sleep(time.Second)//coloca o programa para dormir 1 segundo
	}

	/*Uma instrução "for" com uma ForClause também é controlada por sua condição, mas adicionalmente pode especificar uma instrução init e uma
	 instrução post, como uma atribuição, uma instrução de incremento ou decremento. A instrução init pode ser uma declaração de variável curta ,
	  mas a instrução post não deve. As variáveis ​​declaradas pela instrução init são reutilizadas em cada iteração.*/

	//enquanto j for menor que 10 o programa vai incrementando o valor de j.
	for j := 0; j < 10; j++{//[ Init ] ";" [ Condição ] ";" [ Post ] = aqui o J só está sendo declarado dentro do for, então saindo desse escopo a variável não existe.
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	//Agora vamos ver o For com a cláusula range
	/*Uma instrução "for" com uma cláusula "range" itera por todas as entradas de uma matriz, fatia, string ou mapa, ou valores recebidos em um canal. 
	Para cada entrada, ele atribui valores de iteração às variáveis ​​de iteração correspondentes, se presentes, e então executa o bloco.
	Iterar é a ação de repetir algo. Na programação, iteração significa a repetição de um conjunto de instruções por uma quantidade finita de vezes ou então, 
	enquanto uma condição seja aceita*/
	//Array
	nomes := [3]string{"Alana", "Pedro", "Paulo"}
	for indice, nome := range nomes{ //o indice mostra a posição do iteravel e o segundo é o valor, esse é sempre o padrão.
		fmt.Println(indice, nome) //também é possível usar o _ para não mostrar determinada info
	}

	//também é possível iterar uma STRING, fazer o programa passar por todas as letras da palavra
	for indice, letra := range "PALAVRA"{
		fmt.Println(indice, string(letra))
	}

	//Iteração em um maps- Estrutura de guardar dados:
	usuario := map[string]string{
		"nome": "Alana",
		"sobrenome": "Peres",
	}

	for chave, valor := range usuario{
		fmt.Println(chave, valor) //no maps o índice é a chave 
	}

	//o Range não funciona em struct!!
}