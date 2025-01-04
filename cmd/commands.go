package cmd

import "github.com/urfave/cli/v2"

func GetCommands() cli.Commands {
	return cli.Commands{
		initCommand,
		addCommand,
		commitCommand,
	}
}
