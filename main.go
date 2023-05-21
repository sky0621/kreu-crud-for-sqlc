package main

import (
	"errors"
	"log"
	"os"

	"github.com/sky0621/kreu-crud-for-sqlc/internal"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "kreu-crud-for-sqlc",
		Usage: "create a CRUD diagram based on the file for sqlc.",
		Action: func(c *cli.Context) error {
			args := c.Args().Slice()
			if len(args) == 0 {
				return errors.New("need 'root path' arg")
			}
			return execMain(args[0], args[1:])
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func execMain(rootPath string, targetSQLName []string) error {
	result, err := internal.CollectSQLParseResult(rootPath, targetSQLName)
	if err != nil {
		return err
	}

	if err := internal.Output(result); err != nil {
		return err
	}
	return nil
}
