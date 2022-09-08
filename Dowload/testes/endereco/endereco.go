/*Os testes automatizados em GO é uma função que vai testar outra funções e ver se o resultado dessa função é oq vc espera.
Serve para garantir q a função vai retornar oq é esperado. Da mais segurança para fazer alterações em programas que está em andamento.*/
package endereco

import (
 	"strings"
)


//TiposDeEndereco verifica se um endereço tem um tipo válido e o retorna.
func TipoDeEndereco(endereco string) string { //slice de string
	tiposValidos := []string {"rua", "avenida", "estrada", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)//estou deixando todas as letras minúsculas.
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0] // Split é uma função do pacote strings que divide a string em slice de acordo com o separador.Ex: RUA DOS BOBOS = "RUA" "DOS" "BOBOS". 
	//É RECORTADA ENTRE OS ESPAÇOS. e nesse caso estou chamando a posição [0]
	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"

}

