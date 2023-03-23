package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	// create new app
	app := cli.NewApp()

	// add flags to the app
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "name", Value: "stranger", Usage: "your wonderful name"},
		cli.IntFlag{Name: "time", Value: 0, Usage: "time practicing go"},
		cli.StringFlag{Name: "future", Value: "exciting", Usage: "your prediction of ur future"},
	}

	// this function parses and brings data in cli.Context (cmd) struct
	app.Action = action
	// pass os.Args to cli app to parse content
	app.Run(os.Args)
}

func action(c *cli.Context) error {
	// c.String, c.INt looks for value of given flag
	fmt.Printf("Hello sir %s who has wondered this lands for %d weeks, and certainly more to come, Welcome "+
		"to the command line world\nThis is the forecast of your future with Go %s\n",
		c.String("name"), c.Int("time"), c.String("future"))
	return nil
}
