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

	os.Setenv("SAMPLE_ENV", "sample env")

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "language for the greeting",
		},
		cli.StringFlag{
			Name:  "meridian, m",
			Value: "AM",
			Usage: "meridian for the greeting",
		},
		cli.StringFlag{
			Name:  "time, t",
			Value: "07:00",
			Usage: "'your time' for the greeting",
		},
		cli.StringFlag{
			Name:   "aaa, a",
			Value:  "sample",
			EnvVar: "SAMPLE_ENV",
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("-- Action ==")

		fmt.Printf("c.GlobalFlagNames()         :%+v\n", c.GlobalFlagNames())
		fmt.Printf("c.String(\"lang\")          :%+v\n", c.String("lang"))
		fmt.Printf("c.String(\"m\")             :%+v\n", c.String("m"))
		fmt.Printf("c.String(\"time\")          :%+v\n", c.String("time"))
		fmt.Printf("c.String(\"a\")             :%+v\n", c.String("a"))

		cli.ShowVersion(c)

		return nil
	}

	app.Run(os.Args)
}
