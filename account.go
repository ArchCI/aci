package main

import (
	"github.com/codegangsta/cli"
)

const EXIT_STATUS = 1

// NewAccountCommand refers to account
func NewAccountCommand() cli.Command {
	return cli.Command{
		Name:  "account",
		Usage: "create or get account",

		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "id",
				Value: "",
				Usage: "user id",
			},
			cli.StringFlag{
				Name:  "email",
				Value: "",
				Usage: "user email",
			},
		},

		Action: func(c *cli.Context) {
			//buildCommandFunc(c)
			//fmt.Println("call action fuc")
		},
	}
}
