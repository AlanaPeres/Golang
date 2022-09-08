package main //o nome do arquivo não importa quando vamos executar e sim o nome do pacote

import "fmt"

var y = 45 //var é uma palavra chave
//todas as variáveis possuem um tipo (int, string, float, bool) estando esse tipo explicito ou não 

func main() {
	// com Declarações curtas de variáveis ( := )é possível atribuir/declarar ao mesmo tempo
	x := 10 //estamos declarando x e lhe atribuindo o valor 10
	fmt.Println(x)
	fmt.Println(y) //se vc precisa declarar uma variável fora do corpo de uma função, use VAR
	
	var ( // dessa maneira posso declarar muitas variáveis ao mesmo tempo.
		exemplo1 int8 = 65 // = é apenas uma atribuição.
		exemplo2 int8 = 69
	)
	fmt.Println(exemplo1, exemplo2)

    variavel4, variavel5 := 25, 45
	fmt.Println(variavel4, variavel5)
	
	/*limite o escopo de suas variáveis e tanto quanto possível, tente utilizar :=.
	escopo de uma declaração é a parte do código-fonte em que um uso do nome declarado refere-se a essa declaração. É uma região 
	do texto do programa. Caso você mantenha todas as suas funções com Escopo global, há grandes chancer para erro. Elas devem estar 
	diponíveis apenas onde são necessárias
	existem várias formas de declarar uma variável em GO e o mesmo se aplica a constantes. VAR E CONST se diferenciam apenas pela const 
	ter seu valor imutável.*/
}