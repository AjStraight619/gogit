package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var initCommand = &cli.Command{
	Name:      "init",
	Usage:     "Initialize a new gogit repository",
	ArgsUsage: "[Specify a directory to initialize gogit in]",
	Action: func(c *cli.Context) error {
		args := c.Args().Slice()
		if len(args) == 0 {
			return fmt.Errorf("no directory specified to initialize gogit")
		}

		if len(args) > 1 {
			return fmt.Errorf("specify only one directory to initialize gogit (e.g., `.` or `path/to/proj`)")
		}

		// Define subdirectories
		subdirs := []string{
			".gogit/objects",
			".gogit/refs/heads",
			".gogit/refs/tags",
			".gogit/hooks",
			".gogit/info",
			".gogit/logs/refs",
			".gogit/branches",
		}

		// Create directories
		for _, dir := range subdirs {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", dir, err)
			}
			fmt.Printf("Created directory: %s\n", dir)
		}

		// Create files
		files := map[string]string{
			".gogit/HEAD":   "ref: refs/heads/master\n",
			".gogit/config": "",
			".gogit/index":  "",
		}

		for path, content := range files {
			if err := createFile(path, content); err != nil {
				return fmt.Errorf("failed to create file %s: %w", path, err)
			}
		}

		dir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get working directory: %w", err)
		}

		fmt.Printf("Successfully initialized repository in: %s\n", dir)
		return nil
	},
}

func createFile(path, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if content != "" {
		_, err = file.WriteString(content)
		if err != nil {
			return err
		}
	}

	return nil
}
