package main

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
	"github.com/urfave/cli/v2"
)

func main() {
	logger := slog.New(tint.NewHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cliApp := cli.NewApp()
	cliApp.Name = "photos"
	cliApp.Description = "an app to index & view photos & videos from your disks"

	cliApp.Commands = []*cli.Command{
		{
			Name:        "index",
			Description: "indexes photos & videos from given paths",
			Usage:       "migrate [command]",
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					Name:    "paths",
					Aliases: []string{"p"},
					Usage:   "paths to index",
				},
			},
			Action: index,
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}
