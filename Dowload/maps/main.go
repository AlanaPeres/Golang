package main /*MAPS TAMBÉM É UMA ESTRUTURA PARA GUARDAR DADOS, é uma estrutura de chave e valor mais rígida,
as chaves precisam ser do mesmo tipo e os valores precisam ser do mesmo tipo, mas as chaves e os valores não precisam ser 
do mesmo tipo.
*/

import (
	"fmt"
)

func main() {
	fmt.Println("MAPS")

	usuario := map[string]string{ //essa é a sintaxe do map, estou definindo o tipo das chaves e o valor como string. 

		"nome": "Alana",
		"sobrenome": "Peres",
	}
	fmt.Println(usuario) 
	fmt.Println(usuario["nome"]) // Os elementos são acessados por meio da notação usual de indexação.

	idade:= map[string]int {
		"alice" : 31,
		"Julia" : 15,
	}
		fmt.Println(idade)

	//PODEM SER REMOVIDAS COM A FUNÇÃO EMBUTIDA DELETE:
	delete(idade,"alice")
	fmt.Println(idade)

	//As formas abreviadas de atribuição (x += y e x++)também funcionam em maps
	idade["Julia"]++
	fmt.Println(idade, "Parabéns")
}
