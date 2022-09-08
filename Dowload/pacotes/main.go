package main /*Pacote em Go é simplesmente um diretório (pasta) com um ou mais arquivos .go dentro dele. Pacotes Go
 fornecem isolamento e organização de código semelhante a diretórios e organização de arquivos em um computador. Todo
  o código Go vive em um pacote e um pacote é o ponto de entrada para acessar o código Go*/

import (
	"fmt"
	"modulo/externo" //para eu conseguir exportar um pkg q eu criei eu vou no import e chamo o nome do múdulo/pacote
)

func main() {
	fmt.Println("do pacote main")
	externo.Ola("Alana")
}
/*Ao criar um pacote em Go , normalmente, o objetivo final é tornar o pacote acessível para uso por outros 
desenvolvedores, em pacotes de ordem mais elevada ou em programas inteiros. Ao importar o pacote , seu código poderá 
servir como bloco de construção para outras ferramentas ou blocos mais complexos. No entanto, apenas pacotes estão 
disponíveis para distribuição*/