package root

import (
	"fmt"
	"log"
	"os"

	"github.com/canonical/promql-transform/pkg/transform"
	cli "github.com/urfave/cli/v2"
)

var app = &cli.App{
	Name:            "promql-transform",
	Usage:           "Transforms PromQL Expressions on the fly",
	HideHelpCommand: true,
	HideHelp:        true,
	ArgsUsage:       "expression",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:     "label-matcher",
			Required: true,
			Usage:    "Label matcher to inject into all vector selectors",
		},
	},
	Action: func(c *cli.Context) error {
		args := c.Args()

		if args.Len() != 1 {
			log.Fatal("Expected exactly one argument: the expression.")
		}

		inj, err := transform.GetLabelMatchers(c.StringSlice("label-matcher"))
		if err != nil {
			log.Fatal(err)
		}

		output, err := transform.Transform(args.First(), &inj)
		if err != nil {
			return err

		}

		fmt.Print(output)
		return nil
	},
}

func Execute() error {
	return app.Run(os.Args)
}
