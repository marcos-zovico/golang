package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

// Gerar vai retornara a aplicação cli para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca Ips e nomes de servidores na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca endereços de IPs na internet",
			Flags:  flags,
			Action: buscarIp,
		},
		{
			Name: "server",
			Usage: "Busca oos nomes dos servidores",
			Flags: flags,
			Action: buscarServidores,
		},
	}
	return app
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")
	servers, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

func buscarIp(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
