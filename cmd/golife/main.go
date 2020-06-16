package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/ayoisaiah/golife"
	"github.com/urfave/cli/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const version = "v0.1.0"

func main() {
	app := &cli.App{
		Name: "golife",
		Authors: []*cli.Author{
			{
				Name:  "Ayooluwa Isaiah",
				Email: "ayo@freshman.tech",
			},
		},
		Usage:     "An implementation of John Conway's Game of Life simulation, written in Go!",
		UsageText: "[options]",
		Version:   version,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "refresh-rate",
				Aliases: []string{"fps"},
				Usage:   "Set the speed of simulation",
				Value:   15,
			},
			&cli.StringFlag{
				Name:    "theme",
				Aliases: []string{"t"},
				Usage:   "Select a theme. Use 'golife themes' to see available themes",
				Value:   "WhiteOnBlack",
			},
			&cli.StringFlag{
				Name:    "rule",
				Aliases: []string{"r"},
				Usage:   "Select a rule. Use 'golife rules' to see available rules",
				Value:   "Default",
			},
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "Load a preset pattern from a file",
			},
			&cli.StringFlag{
				Name:    "url",
				Aliases: []string{"u"},
				Usage:   "Load a preset pattern from a url",
			},
			&cli.StringFlag{
				Name:    "input-format",
				Aliases: []string{"i"},
				Usage:   "Specify format when loading preset pattern from url. Valid values are rle, life106, and plaintext",
				Value:   "rle",
			},
			&cli.BoolFlag{
				Name:    "wrap",
				Aliases: []string{"w"},
				Usage:   "Specify if the cells should wrap around the edges of the board",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "themes",
				Usage: "List available themes",
				Action: func(c *cli.Context) error {
					golife.ListThemes()
					return nil
				},
			},
			{
				Name:  "rules",
				Usage: "List available rules",
				Action: func(c *cli.Context) error {
					golife.ListRules()
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			g, err := golife.NewGame(c)
			if err != nil {
				return err
			}

			return g.Start()
		},
	}

	// Override the default help template
	cli.AppHelpTemplate = `DESCRIPTION:
	{{.Usage}}

USAGE:
   {{.HelpName}} {{if .UsageText}}{{ .UsageText }}{{end}}
{{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}{{end}}
{{if .Version}}
VERSION:
	 {{.Version}}{{end}}
{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{ if .VisibleFlags }}
FLAGS:{{range .VisibleFlags}}
	 {{.}}
	 {{end}}{{end}}
WEBSITE:
	https://github.com/ayoisaiah/golife
`

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
