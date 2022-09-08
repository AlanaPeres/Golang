package app

import (
	"github.com/urfave/cli" // a gente sempre referencia o pacote pelo ultimo nome que vem após a /
	"log"
	"net"
	"fmt"
)
//gerar vai retornar a aplicação de Linha de comando pronta para ser executada.
func Gerar() *cli.App { // como esse não é o pacote main, eu não crio uma func main e o nome da minha função precisa ter a letra Maiuscula, para q o pck main importe essa função
	app := 	cli.NewApp()			//o App é um tipo que está dentro do pacote cli
	app.Name = "Aplicação de linha de comando"	//nome da aplicação
	app.Usage = "Busca Ips e Nomes de Servidores na Internet" // para que serve a aplicação

	flags := []cli.Flag{
		cli.StringFlag {
			Name: "host",
			Value: "amazon.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca IPS de endereço na Internet",
			Flags: flags, //flags é como se fosse o parâmetro que a gente passa para esse comando funcionar
			Action: buscarIps,	
		},
	
		{
			Name: "Servidores",
			Usage: "Busca o nome do Servidor na internet",
			Flags: flags,
			Action: buscarServidores,
		},
	}
	
	return app
}
	
	func buscarIps(c *cli.Context) {
		host := c.String("host")//

		ips, erro := net.LookupIP(host)
		if erro != nil {
			log.Fatal(erro)
		}

		for _, ip := range ips {
			fmt.Println(ip)
		}
	}

	func buscarServidores(c *cli.Context) {
		host := c.String("host")

		servidores, erro := net.LookupNS(host) // NAME SERVER

		if erro != nil{
			log.Fatal(erro)
		}

		for _, servidor := range servidores {
			fmt.Println(servidor.Host)
		}
	}
