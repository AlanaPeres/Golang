package main // Array é uma lista de valores, que possui o tamanho e tipo fixos. 

import "fmt"

func main() {
	fmt.Println("Array e Slice")

	var Array1 [5]int //Array só armazena ítens do mesmo tipo e é obrigatório especificar o seu tamanho
	Array1[0] = 5 //para popular o array podemos fazer assim
	fmt.Println(Array1)

	//OUTRA FORMA DE POPULAR O ARRAY
	Array2 := [5]int {5, 4, 3, 2, 1}
	fmt.Println(Array2)

	fruta := [3]string {"melão", "abacate", "uva"}
	fmt.Println(fruta)

	//uma forma de tentar deixar o tamanho do array mais maleável é fazendo da seguinte maneira:

	Array3 := [...]int {1, 2, 3} //enquanto eu for aumentando os ítens dentro das {} o array vai aumentando seu tamanho
	fmt.Println(Array3)
}