package formas

import "math"

type Forma interface{ //o que o método faz dentro das chaves pode variar, mas a assinatura para ser considerada uma forma, precisa ser exatamente igual(area float64)
	Area() float64 //então o nome do método precisa ser o mesmo que area()float64
}

type Retangulo struct {  // SE EU FOSSE UTILIZAR A STRUCT PARA CALCULAR A ÁREA EU TERIA PROBLEMAS PARA CALCULAR A ÁREA DE TODAS AS FORMAS, TENDO EM VISTA QUE CADA FORMA POSSUI
	Altura float64 	// UM MÉTODO DIFERENTE DE CALCULAR SUA ÁREA, COMO NESSE EXEMPLO DO CIRCULO E DO RETÂNGULO. COM A INTERFACE EU POSSO IMPLEMENTAR DIVERSAS FORMAS DE FAZER 
	Largura float64 // O CALCULO DA ÁREA.
}

type Circulo struct {
	Raio float64
}
//PARA SATISFAZER A NECESSIDADE DA INTERFACE DE TER UM MÉTODO IGUAL EU VOU IMPLEMENTAR MINHAS FORMAS

func (r Retangulo) Area() float64 { // agora o meu retângulo possui o requisito de método igual.
	return r.Altura * r.Largura
}

func (c Circulo) Area() float64 { // a fórmula para calcular a área de um círculo é pi* o raio²(pi *(raio*raio))
	return math.Pi * (c.Raio * c.Raio)
}

