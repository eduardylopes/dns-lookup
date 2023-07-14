package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate return the command line application ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()

	app.Name = "Command Line Application"
	app.Usage = "Search for public IP addresses based on the dns"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search for public IP in the internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "github.com",
				},
			},
			Action: searchIps,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)

	if error != nil {
		log.Fatal(error)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
