package main

import (
	"fmt"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "Test App Name"
	app.Usage = "Sample Command Line Arguments"
	app.Flags = []cli.Flag{
		// This defines the arguments --name, -n in one go
		// This flag takes in a string value
		cli.StringFlag{
			Name:  "name, n",
			Value: "Default Name Value", // This is the default value in case none is provided
			Usage: "Contains the name text for the binary usage",
		},
		// This defines the arguments --write_mode, -wm in one go
		// This is a boolean flag, returning a true if present
		// returning a false if not present
		cli.BoolFlag{
			Name:  "write_mode, wm",
			Usage: "When defined, write mode is enabled for the task",
			// The default value is `false`
		},
	}
	app.Action = func(c *cli.Context) error {
		// Get the value of the `name` flag
		name := c.GlobalString("name")
		fmt.Printf("Hello %s!\n", name)

		// Get the value of the `write_mode` flag
		mode := c.GlobalBool("write_mode")
		if mode == true {
			fmt.Println("Write mode is on!")
		} else {
			fmt.Println("Write mode is off!")
		}

		// Return a `nil`
		return nil
	}

	app.Run(os.Args)
}
