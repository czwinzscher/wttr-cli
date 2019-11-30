package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "wttr-cli"
	app.Description = "A cli application for getting infos from wttr.in"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "celcius,c",
			Usage: "show temp in Celcius",
		},
		cli.BoolFlag{
			Name:  "fahrenheit,f",
			Usage: "show temp in Fahrenheit",
		},
	}

	app.Action = func(c *cli.Context) error {
		location := c.Args().Get(0)

		wttr, err := GetWeather(location)
		if err != nil {
			return err
		}

		curr := wttr.CurrentCondition[0]

		celcius := c.Bool("celcius")
		fahrenheit := c.Bool("fahrenheit")

		if !celcius && !fahrenheit {
			fmt.Println(curr.TempC + "°C")
		} else {
			if celcius {
				fmt.Println(curr.TempC + "°C")
			}

			if fahrenheit {
				fmt.Println(curr.TempF + "°F")
			}
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
