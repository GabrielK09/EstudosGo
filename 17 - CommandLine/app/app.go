package app

import (
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplication in CLI"
	app.Usage = "Search names and IPs in netwoork"

	//Propriedade do struct do cli.NewApp() e tipo esperado
	app.Commands = []cli.Command{
		// Name = nome do comando no go run main.go ip
		// Usage = Uso
		// Flags = propriedades
		{
			Name:  "ip",
			Usage: "search IPs",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: searchIPs,
		},
		{
			Name:  "servers",
			Usage: "search names",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: searchNames,
		},
	}

	return app
}

func searchNames(c *cli.Context) {
	host := c.String("host")

	names, err := net.LookupNS(host)

	if err != nil {
		log.Fatal("erro for search Names:", err)
	}

	for _, name := range names {
		log.Println("Name:", name.Host)
	}
}

func searchIPs(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal("erro for search IPs:", err)
	}

	for _, ip := range ips {
		log.Println("IP:", ip)
	}
}
