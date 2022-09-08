package formas

import (
	"testing"
	"math"
)

func TestArea(t *testing.T) {
	t.Run("	Retangulo", func(t *testing.T){ // t.Run é utilizado quando queremos separar os testes
		ret := Retangulo{ 10, 12 }
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida{
			t.Fatalf("A área %f é diferente da área %f.", areaEsperada, areaRecebida)
		}
	})

	t.Run("Circulo", func(t *testing.T){
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida{
			t.Fatalf("A área %f é diferente da área %f.", areaEsperada, areaRecebida)
		}
	})
}