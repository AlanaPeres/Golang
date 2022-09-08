package main
/*A programação orientada a Objetos tem sido o paradigma da programação dominante no mercado e na educação e todas as linguagens amplamente usadas 
dessenvolvidas desde então incluiram suporte a ela. Em go Um objeto é simplesmente um valor ou uma variável que tem métodos, e um método é uma função
associada a um tipo particular.  
Um programa orientado a objetos usa métodos para expressar as propriedades e operações de cada estruturas de dados, de modo que os clientes não precisem
acessar a representação do objeto diretamente.
Um método é declarado como uma variante da declaração normal de função, em que um parâmetro extra aparece antes do nome da função. O parâmetro associa a função
ao tipo de parâmetro.
Os métodos podem ser declarados em qualquer tipo nomeado definido no mesmo pacote, desde que seu tipo subjacente não seja um ponteiro nem uma interface.
Também é possível passar parâmetros e pedir retornos assim como nas funções.*/

import "fmt"
//a diferença entre um método e uma função é que o método sempre precisa estar associada a um tipo particular.
type usuario struct{
	nome string
	idade uint
}
//um método é uma função que é declarada com um receptor.
/*O PARÂMETRO EXTRA u é chamado de receptor do método- um legado das primeiras linguagens orientadas a objetos que descreviam uma chamada de método como 
o envio de uma mensagem ao objeto. Como o nome do receptor será usado com frequencia, sugere-se colocar um nome curto e consistente entre os métodos. Uma
opção é sempre colocar a primeira letra do nome do tipo.*/

func (u usuario) salvar() { //salvar é o nome do método, usuário é o nome da strutura a qual o método está ligado, u é o nome do receptor(paramentro extra) Em uma chamada de 
	fmt.Println("Dentro do método Salvar") //método o argumento para o receptor aparece antes do nome do método. 
}

func main() {
	fmt.Println("Métodos")
	usuario1 := usuario{"Usuário 1", 20}
	fmt.Println(usuario1)
}