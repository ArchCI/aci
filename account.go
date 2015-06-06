package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

const EXIT_STATUS = 1

// NewAccountCommand refers to access account data
func NewCreateAccountCommand() cli.Command {
	return cli.Command{
		Name:  "create_account",
		Usage: "create the new account",

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
			var id = c.String("id")
			var email = c.String("email")

			if len(c.Args()) <= 1 {
				fmt.Println("No eough parameters, need username and password")
				os.Exit(0)
			}

			var username = c.Args()[0]
			var password = c.Args()[1]

			// TODO: send post request to create account
			fmt.Println(id)
			fmt.Println(email)
			fmt.Println(username)
			fmt.Println(password)

			os.Exit(0)
		},
	}
}
