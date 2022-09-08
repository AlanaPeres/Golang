/*O Sistema de tipos em GO captura muitos erros em tempo de compilação, mas outros, como um acesso fora dos limites de uma array ou desreferenciar um
ponteiro nil, exigem verificações em tempos de execução. Quando o runtime de GO identifica esses erros, um PANIC é gerado. Durante um pânico típico 
a execução normal é interrompida, todas as chamadas de funções adiadas nessa goroutin são executadas e o programa falha com uma mensagem de log. Essa
mensagem de log inclui o valor de panic  que geralmente é uma espécie de mensagem de erro e, para cada goroutine, uma stack trace (estado da pilha de execução)
mostrando a pilha das chamadas de funções que estavam ativas no momento do panic. Em geral, a mensagem de log tem informações  suficientes para diagnosticar
a causa-raiz do problema sem executar denovo o programa, portanto ela sempre deve ser incluída em um relatório de bug sobre um programa que gera pânico.
O PANIC MATA A EXECUÇÃO DO PROGRAMA.*/
package main

import "fmt"

func alunoAprovado(n1, n2 float64) bool{
	media := (n1 + n2) / 2

	if media > 6{
		return true
	} else if media < 6 {
		return false
	}
	panic("MÉDIA IGUAL A 6")//para o panic funcionar é necessário ter uma mensagem.
}

func main(){
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("Pós Execução") // APOS O PANIC NADA QUE ERA PARA SER EXECUTADO DEPOIS PODERÁ FUNCIONAR, POIS O PANIC MATA A EXECUÇÃO.
}