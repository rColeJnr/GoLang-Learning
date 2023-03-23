package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	// define flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "save", Value: "no", Usage: "Should save to database (y/n)",
		},
	}

	app.Version = "1.0"

	// define action
	app.Action = action

	app.Run(os.Args)
}

func action(c *cli.Context) error {
	var args []string
	if c.NArg() > 0 {
		// Fetch arguments in an array
		args = c.Args()
		personName := args[0]
		marks := args[1:len(args)]
		fmt.Println("Person: ", personName)
		fmt.Println("marks: ", marks)
	}

	if c.String("save") == "no" {
		fmt.Println("Skipping saving to the db")
	} else {
		// add db logic here, but not for real
		fmt.Println("Saving to the db", args)
	}
	return nil
}
