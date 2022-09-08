package main
/*INTERFACE: Uma interface permite que uma coisa funcione com outra. OS tipos interface expressam generalizações ou abstrações sobre os comportamentos 
de outros tipos. Ao generalizar, as interfaces permitem escrever funções mais flexíveis e adaptáveis pq não estão amarradas aos detalhes de uma 
implementação em particular. Interfaces são coleções nomeadas de assinaturas de métodos. 
Quando você vê que um func tem um parâmetro desse tipo "...interface{}", isso é chamado de "parâmetro variádico" e significa que o func pode receber
 quantos valores desse tipo você quiser passar.*/

import (
	"fmt"
	"math"
)
//forma é uma interface que tem um método chamado area / Esse método não recebe parametros e retorna um float64
type forma interface{ //o que o método faz dentro das chaves pode variar, mas a assinatura para ser considerada uma forma, precisa ser exatamente igual(area float64)
	area() float64 
}

type retangulo struct {  // SE EU FOSSE UTILIZAR A STRUCT PARA CALCULAR A ÁREA EU TERIA PROBLEMAS PARA CALCULAR A ÁREA DE TODAS AS FORMAS, TENDO EM VISTA QUE CADA FORMA POSSUI
	altura float64 	// UM MÉTODO DIFERENTE DE CALCULAR SUA ÁREA, COMO NESSE EXEMPLO DO CIRCULO E DO RETÂNGULO. COM A INTERFACE EU POSSO IMPLEMENTAR DIVERSAS FORMAS DE FAZER 
	largura float64 // O CALCULO DA ÁREA.
}

type circulo struct {
	raio float64
}
//PARA SATISFAZER A NECESSIDADE DA INTERFACE DE TER UM MÉTODO IGUAL EU VOU IMPLEMENTAR MINHAS FORMAS
//QUALQUER ESTRUTURA QUE TENHA UM MÉTODO AREA() COM A MSM ASSINATURA, AUTOMATICAMENTE JÁ A IMPLEMENTA.
func (r retangulo) area() float64 { // agora o meu retângulo possui o requisito de método igual.
	return r.altura * r.largura
}

func (c circulo) area() float64 { // a fórmula para calcular a área de um círculo é pi* o raio²(pi *(raio*raio))
	return math.Pi * (c.raio * c.raio)
}

func CalculoArea(f forma) { //Aqui eu tenho uma função que ao invés de receber uma estrutura concreta como é o caso do Struct, ela recebe uma estrutura abstrata que é o caso da interface.
	fmt.Printf ("A área da forma é %0.2f\n", f.area())
}

func main() {
	//para passar os valores da área é assim:
	r := retangulo {10, 15}
	CalculoArea(r)// ai eu chamo a função calculo de área passando o retangulo como parâmetro. Para que o retângulo possa utilizar essa interface ele precisa obrigatoriamente ter um método chamado área que não recebe parâmetros e que retorne um float64.
	
	c := circulo {15}
	CalculoArea(c)
}
/*dizer que eu preciso ir daqui para LA Vou precisar de um veículo para chegar lá. carro, caminhão, moto, avião todos esses diferentes modos de transporte implementam a "interface do veículo"
todos eles satisfazem os critérios para o que significa ser um veículo. Um policial pode ter uma regra, "um veículo não pode passar por um sinal vermelho"
um policial poderia então pará-lo se você passasse por um sinal vermelho em uma moto, moto, carro ... todos esses diferentes modos de transporte implementam a "interface do veículo"
todos eles satisfazem os critérios para o que significa ser um veículo.

Você poderia dizer que as Interfaces nos permitem agrupar as coisas por funcionalidade, o que elas fazem e seus métodos.  um tipo que define um conjunto de métodos
Uma interface é como ter mais de um tipo: eu sou um avião e sou um veículo, eu sou um carro e sou um veículo, eu sou um barco e sou um veículo.
Sou um arquivo e sou uma interface de leitura. Sou uma string e sou uma interface. 

Em sua forma mais básica, uma interface especifica uma lista (possivelmente vazia) de métodos. O conjunto de tipos definido por tal interface é o 
conjunto de tipos que implementam todos esses métodos, e o conjunto de métodos correspondentes consiste exatamente dos métodos especificados pela 
interface. Interfaces cujos conjuntos de tipos podem ser definidos inteiramente por uma lista de métodos são chamadas interfaces básicas.*/