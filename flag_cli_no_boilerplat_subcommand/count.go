package main

import (
	"fmt"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Count up or down"
	app.Commands = []cli.Command{
		{
			Name:      "up",
			ShortName: "u",
			Usage:     "Count numbers from 0",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "stop, s",
					Usage: "Value to count up to",
					Value: 12,
				},
			},
			Action: func(c *cli.Context) error {
				stop := c.Int("s")
				if stop <= 0 {
					fmt.Println("Stop can't be lower than or equal to 0")
				}
				for i := 0; i <= stop; i++ {
					fmt.Println(i)
				}
				return nil
			},
		},
		{
			Name:      "down",
			ShortName: "d",
			Usage:     "Count numbers to 0",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "start, s",
					Usage: "Value to count up to",
					Value: 12,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("s")
				if start <= 0 {
					fmt.Println("Stop can't be lower than or equal to 0")
				}
				for i := start; i >= 0; i-- {
					fmt.Println(i)
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}
