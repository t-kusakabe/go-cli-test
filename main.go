package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Test"
	app.Usage = "Sample is foo!"
	app.Version = "1.0"


	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Value: "tarou",
			Usage: "your name",
		},
    cli.BoolFlag {
      Name: "gay, g",
      Usage: "are you gay boy?",
    },
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("-- Action ==")

		fmt.Printf("c.GlobalFlagNames()       :%+v\n", c.GlobalFlagNames())
		fmt.Printf("c.String(\"name\")        :%+v\n", c.String("name"))
		fmt.Printf("c.Bool(\"g\")             :%+v\n", c.Bool("g"))

		cli.ShowVersion(c)

		return nil
	}

	app.Run(os.Args)
}
