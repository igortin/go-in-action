package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Count up or down."
	app.Commands = []cli.Command{
		{
			// Define cli.Command
			Name:      "up",
			ShortName: "u",
			Usage:     "Count Up",
			// Define flag for cli.Command
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "value, v",
					Usage: "./urfave-cli up --value 5",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				// Get value
				v := c.Int("value")

				// Check parameter restrictions
				if v <= 0 {
					fmt.Println("value cannot be negative.")
				}

				// Bussiness logic
				for i := 1; i <= v; i++ {
					fmt.Println(i)
				}

				return nil
			},
		},
		{
			// Define cli.Command
			Name:      "down",
			ShortName: "d",
			Usage:     "Count Down",
			// Define flag for cli.Command
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "value, v",
					Usage: "./urfave-cli down --value 5",
					Value: 5,
				},
			},
			Action: func(c *cli.Context) error {
				// Get value
				v := c.Int("value")

				// Check parameter restrictions
				if v <= 0 {
					fmt.Println("value cannot be negative.")
				}

				// Bussiness logic
				for i := v; i > 0; i-- {
					fmt.Println(i)
				}

				return nil
			},
		},
	}

	app.Run(os.Args)
}
