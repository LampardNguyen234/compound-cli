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
				Description: "This command helps register to the compound service. It basically creates a transaction composing of two following `sdk.Msg`s: MsgWithdrawDelegationReward, MsgDelegate. " +
					"This helps authorize the compound operator to perform compounding on behalf of the user.",
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
				Name:        "unregister",
				Aliases:     []string{"unreg"},
				Usage:       "Unregister from the compound service",
				Description: "This command helps un-register from the compound service. It basically revokes compound-related authorizations granted before.",
				Flags: []cli.Flag{
					defaultFlags[prvKeyFlag],
					defaultFlags[operatorFlag],
				},
				Action: unregisterCompound,
			},
		},
	},
}
