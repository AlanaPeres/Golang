package main
//Também é uma estrutura de controle em Go, a instrução switch nos permite executar um bloco de código entre muitas alternativas, simplificando uma cadeia
//if-else. 
import (
	"fmt"
)
func DiadaSemana(numero int)string {
	switch numero {

	case 1:
		return "Domindo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default: // caso nenhuma das opções acima seja executada essa cláusula é chamada.
		return "Número Inválido"
	}
}
//outra forma de utilizar o swicth

func DiadaSemana2(numero int)string{
	var DiadaSemana string
	switch{
	case numero == 1:
		DiadaSemana = "Domingo"
	case numero == 2:
		DiadaSemana = "Segunda"
	case numero == 3:
		DiadaSemana = "Terça"
	case numero == 4:
		DiadaSemana = "Quarta"
	case numero == 5:
		DiadaSemana = "Quinta"
	case numero == 6:
		DiadaSemana = "Sexta"
	case numero == 7:
		DiadaSemana = "Sábado"
	default:
		DiadaSemana = "Número Inválido"
	}
	return DiadaSemana
}
func main() {
	fmt.Println("Swicth ")
	dia :=  DiadaSemana(5)
	fmt.Println(dia)

	fmt.Println("------------------")
	dia2 :=DiadaSemana2(7)
	fmt.Println(dia2)



}	