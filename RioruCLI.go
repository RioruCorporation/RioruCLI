package main

import (
	"RioruCli/src"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	commands := []*cli.Command{
		{
			Name: "download",
			Aliases: []string { "d" },
			Description: "Download all Repositories of Rioru",
			Action: func(c *cli.Context) error {
				src.GetRepositories()
				return nil
			},
		},
	}

	app := &cli.App{
		Name: "Rioru CLI",
		Usage: "A Simple Command Line Interface to manage Rioru Projects",
		UsageText: "./RioruCli command",
		CommandNotFound: func(c *cli.Context, command string) {
			fmt.Println("Not have the command \"" + command + "\" in the Rioru CLI")
		},
		Commands: commands,
		HideHelpCommand: true,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
