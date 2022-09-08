package main /*O slice também se comporta como uma lista de valores, porém seu tamanho é dinâmico. 
O slice se parece com um Array, a verdade é que ele é um pedaço de um Array*/

import "fmt"

func main() {
	
	fmt.Println("Slice ")

	Array2 := [5]int {5, 4, 3, 2, 1}
	fmt.Println(Array2)

	slice1 := []int {3, 5}// essa é a lista mais utilizada em GO
	fmt.Println(slice1)

	//a função append serve para inserir mais um valor no slice
	slice1 = append(slice1, 7)
	fmt.Println(slice1)

	//O slice se parece com um Array, a verdade é que ele é um pedaço de um array
	Slice3 := Array2 [0:3] //aqui ele funciona como se fosse um ponteiro, que aponta para essas posições [0:3]
	fmt.Println(Slice3)

	//Quando o Array sofre alguma alteração o mesmo acontece com o slice:
	Array2[0] = 9 //aqui eu alterei a posição 0 do Array para o número 9
	fmt.Println(Slice3)

}