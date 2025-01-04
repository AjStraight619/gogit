package cmd

import (
	"fmt"
	"os"

	"github.com/Ajstraight619/gogit/internal/utils"
	"github.com/urfave/cli/v2"
)

var addCommand = &cli.Command{
	Name:      "add",
	Usage:     "Add a new item",
	ArgsUsage: "[files...]",

	Action: func(c *cli.Context) error {
		args := c.Args().Slice()
		if len(args) == 0 {
			return fmt.Errorf("no files specified to add")
		}

		resolvedPaths := make([]string, 0, len(args))

		for _, arg := range args {
			// fmt.Printf("%s ", arg)
			resolvedPath, err := utils.ResolvePath(arg)
			if err != nil {
				return fmt.Errorf("error resolving path '%s': %w", arg, err)
			}
			resolvedPaths = append(resolvedPaths, resolvedPath)

		}

		filesInfo := []os.FileInfo{}

		for _, path := range resolvedPaths {

			err := utils.ProcessPaths(path, &filesInfo)

			if err != nil {
				return fmt.Errorf("error stating file '%s': %w", path, err)
			}

		}

		for _, file := range filesInfo {

			fmt.Println("Current file: ")
			fmt.Printf("File name: %s\n", file.Name())
			fmt.Printf("File size: %d\n", file.Size())
			fmt.Printf("File is dir?: %t\n", file.IsDir())
			fmt.Printf("File mod time: %s\n", file.ModTime())
			fmt.Println("---------------------")
		}
		return nil
	},
}
