package main //UNMARSHAL SERVE PARA CONVERTER UM JSON EM STRUCT OU MAP


import (
	"encoding/json" //pacote utilizado para fazer as conversoes de JSON
	"fmt"
	"log"
	"bytes"
)
//aqui eu estou dizendo que quando minha struct for convertida em json eu quero que ela fique nessas chaves
type cachorro struct {
	Nome string `json:"nome"` 
	Raca string `json:"raca"` 
	Idade uint `json:"idade"`
}

func main() {

	cachorroEmJSON := `{"nome":"Lucke", "Raca":"York", "Idade":2}` //aqui eu já tenho um JSON e vou transforma-lo em struct
	var c cachorro //criei uma variável do tipo cachorro
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil{ //json só envia e recebe dados em SLICE DE bytes []bytes, precisamos sempre converter.
		log.Fatal(erro)
	}

	fmt.Println(c)

	cachorro2EmJSON := `{"nome":"Lucke", "Raca":"York"}`
	c2 := make[map[string]string]
	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil{
		log.Fatal(erro)
	}

	fmt.Println(c2)

}