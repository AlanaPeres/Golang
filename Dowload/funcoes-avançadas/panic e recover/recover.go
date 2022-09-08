/*a Clausula RECOVER serve para recuperar um programa que está em panic. Desistir geralmente é a resposta correta para um panic, mas nem sempre. 
Talvez seja possível recuperar-se de algum modo ou, pelo menos arrumar a bagunça antes de encerrar. 
Se a função RECOVER for chamada em um função adiada 
e a função que contém a instrução DEFER gerar panic, recover finaliza o estado atual de panic e devolve seu valor. A função que gerou panic não continua
do ponto em que parou, mas retorna normalmente. Se recover for chamado em qualquer outro momento, ele não terá nenhum efeito e devolverá nil.*/

package main

import "fmt"

//como chamamos o recover:
func recuperar() {
	if r := recover(); r != nil{ //aqui o recover vai armazenar o seu valor e caso esse valor seja diferente de nil a função vai retomar a execução.
		fmt.Println("Recuperado.")
	}
}//Oq ocorre na execução é que quando o panic foi gerado o programa chamou a função DEFER, que está apontando para a função recover, que conseguiu cumprir 
//com a função de não deixar o programa morrer.

func alunoAprovado(n1, n2 float64) bool{//nesse caso da recuperação o programa vai retornar o valor zero do bool que é false
	media := (n1 + n2) / 2
	defer recuperar()
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