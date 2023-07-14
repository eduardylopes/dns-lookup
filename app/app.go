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

	app.Name = "DNS Lookup"
	app.Usage = "Search for public IP addresses based on the host domain name"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Retrieve the IP addresses from the host",
			Flags:  flags,
			Action: getHostIp,
		},
		{
			Name:   "ns",
			Usage:  "Retrieve the DNS from the host",
			Flags:  flags,
			Action: getHostDNS,
		},
	}

	return app
}

func getHostIp(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func getHostDNS(c *cli.Context) {
	host := c.String("host")

	servers, error := net.LookupNS(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
