/*A janela, que geralmente é chamada de linha de comando ou interface de linha de comando, é uma aplicação de texto para ver e manipular arquivos em seu
computador. É como o Windows Explorer ou o Finder no Mac, mas sem a interface gráfica. Outros nomes para a linha de comando são: cmd, CLI, prompt, console
 ou terminal.
 Como nessa aplicação vamos utilizar mais de um pacote foi necessário criar um módulo, depois disso dei um go get no pacote github.com/urfave/cli, que é
 o responsável por fazer a aplicação funcionar.*/
package main

import (
	"log"
	"linha-de-comando/app"//linha de comando é o nome do módulo e app é o nome do pacote que foi criado.
	"os"
)

 func main() {
	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {//serve para q o comando do meu S.O reconheça 
		log.Fatal(erro)
    }
 
 }


