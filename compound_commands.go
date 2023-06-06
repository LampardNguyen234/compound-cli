package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var compoundCommands = []*cli.Command{
	{
		Name:        "compound",
		Aliases:     []string{"cmpd"},
		Usage:       "Manage compound functionality",
		Description: fmt.Sprintf("This command helps perform compound-related actions"),
		Category:    compoundCategory,
		Subcommands: []*cli.Command{
			{
				Name:    "register",
				Aliases: []string{"reg"},
				Usage:   "Register to the compound service",
				Flags: []cli.Flag{
					defaultFlags[prvKeyFlag],
					defaultFlags[operatorFlag],
					defaultFlags[allowedListFlag],
					defaultFlags[deniedListFlag],
					defaultFlags[expiredFlag],
				},
				Action: registerCompound,
			},
			{
				Name:    "unregister",
				Aliases: []string{"unreg"},
				Usage:   "Unregister from the compound service",
				Flags: []cli.Flag{
					defaultFlags[prvKeyFlag],
					defaultFlags[operatorFlag],
				},
				Action: unregisterCompound,
			},
		},
	},
}
