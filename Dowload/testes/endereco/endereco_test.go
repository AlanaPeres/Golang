                   //teste de unidade, é a função que vai testar a menor unidade do meu código, nesse caso é uma função. 
package endereco   //os testes são os únicos que podem conter mais de um pacote,
                   //o teste de uma func nunca fica no mesmo arquivo que ela, para isso precisamos abrir O arquivo precisa _text.go para que o go reconheça que é um arquivo de test.

import "testing"   //esse pacote é o padrão do go para testes

type cenarioDeTeste struct{ 
	enderecoInserido string            // Nesse struct o 1º parâmetro é o valor que estamos passando para a função TiposDeEndereco
	retornoEsperado string            // e o 2º campo é o valor que eu espero receber de volta dessa
}

func TestTipoDeEndereço(t *testing.T) { //essa sintaxe é que o go vai saber que essa função é para ser testada. Sintaxe obrigatória.
	cenariosDeTeste := []cenarioDeTeste {
		{"rua Abc", "Rua"},            //aqui podemos testar infinitamente, quanto mais melhor, mas focar apenas nos principais.
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das flores", "Tipo Inválido"},
		{"Estrada Geral", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
	}// go test ./... esse comando no terminal fala pro go entrar em todos os pacotes desse projeto e executar os testes de todo mundo.

	for _, cenario := range cenariosDeTeste {
		tiposDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tiposDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
			tiposDeEnderecoRecebido,
			cenario.retornoEsperado,
			)
		}
	}	
}

