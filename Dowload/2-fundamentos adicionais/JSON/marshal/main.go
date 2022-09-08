package main
//como boa prática os campos do json possui sempre letras minusculas

import (
	"encoding/json" //pacote utilizado para fazer as conversoes de JSON
	"fmt"
	"log"
	"bytes"
)
//aqui eu estou dizendo que quando minha struct for convertida em json eu quero que ela fique nessas chaves
type cachorro struct {
	Nome string `json:"nome"` //como boa prática os campos do json possui sempre letras minusculas
	Raca string `json:"raca"` //o ultimos valor não precisa ser igual ao primeiro, mas faz mais sentido.
	Idade uint `json:"idade"`
}

func main() {
	//   json.Marshal()          //transforma uma struct em um JSON E Unmarshal() faz o processo reverso e transforma um json em struct ou em map
	c := cachorro{"Lucke", "York", 2} //passei os valores para minha struct

	cachorroEmJSON, erro := json.Marshal(c) // transformando minha struct em json e já trantando o erro.
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroEmJSON) //QUANDO EU CHAMO A STRUCT ASSIM O GO VAI ME RETORNAR UM SLICE DE ByTES
	fmt.Println(bytes.NewBuffer(cachorroEmJSON)) // utilizando o pacote para transformar a struct em json e importar as informações desse slice de bytes.

	//PARA TRANSFORMAR UM MAP EM JSON:
	c2 := map[string]string {
		"nome": "Lucke",
		"raca": "York",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorro2EmJSON)
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}