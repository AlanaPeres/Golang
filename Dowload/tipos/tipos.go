package main

/*Um tipo determina um conjunto de valores juntamente com operações e métodos específicos a esses valores. O tipo de 
uma variável define as caracteristicas dos valores que ela pode assumir EX: seu tamanho(numero de bits), como elas são
representadas internamente, as operações intrínsecas que podem ser realizadas e os métodos associados a elas.
Na ciência da computação, os tipos primitivos de dados são um conjunto de tipos básicos de dados dos quais todos os outros
tipos de dados são construídos. Em geral, "tipos de dados primitivos" podem se referir aos tipos de dados padrão 
incorporados em uma linguagem de programação. Os tipos de dados que não são primitivos são referidos como 
derivados ou compostos.

Um tipo pode ser denotado por um nome de tipo, se tiver um, que deve ser seguido por argumentos de tipo se o tipo for
genérico Um tipo também pode ser especificado usando um tipo literal, que compõe um tipo de tipo a partir de tipos 
existentes.

A linguagem pré-declara certos nomes de tipos. Outros são introduzidos com declarações de tipo ou listas de parâmetros 
de tipo. 
*Tipos compostos* — matriz, estructs, ponteiro, função, interface, slice, maps e tipos de canais — podem ser 
construídos usando literais de tipo.

Tipos pré-declarados, tipos definidos e parâmetros de tipo são chamados tipos nomeados. Um pseudônimo denota um tipo 
nomeado se o tipo dado na declaração de alias for um tipo nomeado.

*/

import "fmt"

var y = 45
var x = "Rosa"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)//mostra qual o tipo implícito da variavel
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}

/*Golang por padrão inclui vários tipos primitivos pré-declarados, embutidos

○ booleano 
○ int e float
○ string

● Tipos pré-declarados são usados ​​para construir outros tipos compostos

○ matriz
○ estruct
○ ponteiros
○ slice
○ maps
○ canal*/