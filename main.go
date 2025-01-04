package main

import (
	"log"
	"os"

	"github.com/Ajstraight619/gogit/cmd"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:     "gogit",
		Usage:    "A worse version of git for version control",
		Commands: cmd.GetCommands(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
