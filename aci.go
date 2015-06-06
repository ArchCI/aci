package main

import (
	"os"

	"github.com/codegangsta/cli"
)

// Aci command-line client for ArchCI
func main() {
	app := cli.NewApp()
	app.Name = "aci"
	app.Version = "0.1"
	app.Usage = "command-line client for ArchCI"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "t, token",
			Value:  "",
			Usage:  "server auth token",
			EnvVar: "ARCHCI_TOKEN",
		},
		cli.StringFlag{
			Name:   "s, server",
			Value:  "",
			Usage:  "server address",
			EnvVar: "ARCHCI_SERVER",
		},
	}

	app.Commands = []cli.Command{
		NewCreateAccountCommand(),
	}

	app.Run(os.Args)
}
