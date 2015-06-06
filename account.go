package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
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

			urlStr := "http://127.0.0.1:8080/v1/account"
			data := url.Values{}
			data.Set("id", id)
			data.Add("email", email)
			data.Add("username", username)
			data.Add("password", password)

			client := &http.Client{}
			r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
			//r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
			//r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			resp, _ := client.Do(r)
			fmt.Println(resp.Status)

			os.Exit(0)
		},
	}
}
