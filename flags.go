package main

import (
	"github.com/urfave/cli/v2"
	"time"
)

const (
	hostFlag    = "host"
	portFlag    = "port"
	testnetFlag = "testnet"
	docsFlag    = "generate-docs"

	prvKeyFlag = "privateKey"

	operatorFlag    = "operator"
	allowedListFlag = "allowedList"
	deniedListFlag  = "deniedList"
	expiredFlag     = "expired"

	granterFlag = "granter"
	granteeFlag = "grantee"
)

var defaultFlags = map[string]cli.Flag{
	hostFlag: &cli.StringFlag{
		Name:        hostFlag,
		Aliases:     aliases[hostFlag],
		Usage:       "Tendermint RPC host",
		Value:       "",
		Destination: &host,
	},
	portFlag: &cli.StringFlag{
		Name:        portFlag,
		Aliases:     aliases[portFlag],
		Usage:       "Tendermint RPC port",
		Value:       "",
		Destination: &port,
	},
	testnetFlag: &cli.BoolFlag{
		Name:        testnetFlag,
		Aliases:     aliases[testnetFlag],
		Usage:       "Whether to use testnet",
		Value:       false,
		Destination: &testnet,
	},
	docsFlag: &cli.BoolFlag{
		Name:  docsFlag,
		Usage: "Whether to generate docs only",
		Value: false,
	},

	prvKeyFlag: &cli.StringFlag{
		Name:     prvKeyFlag,
		Aliases:  aliases[prvKeyFlag],
		Usage:    "The Astra private key",
		Required: true,
		EnvVars:  []string{"PRIVATE_KEY"},
	},

	operatorFlag: &cli.StringFlag{
		Name:     operatorFlag,
		Usage:    "The address of the operator",
		Required: false,
		Value:    compoundOperator,
	},
	allowedListFlag: &cli.StringSliceFlag{
		Name:     allowedListFlag,
		Aliases:  aliases[allowedListFlag],
		Usage:    "The list of allowed validator addresses (default: all `bonded` validators). Example: --allowed VALIDATOR_1 --allowed VALIDATOR_2",
		Required: false,
	},
	deniedListFlag: &cli.StringSliceFlag{
		Name:     deniedListFlag,
		Aliases:  aliases[deniedListFlag],
		Usage:    "The list of denied validator addresses (default: no denied validators). Example: --denied VALIDATOR_1 --denied VALIDATOR_2",
		Required: false,
	},
	expiredFlag: &cli.StringFlag{
		Name:     expiredFlag,
		Aliases:  aliases[expiredFlag],
		Usage:    "The expiration duration (e.g, 1000s)",
		Required: false,
		Value:    (10 * 365 * 24 * time.Hour).String(),
	},

	granterFlag: &cli.StringFlag{
		Name:     granterFlag,
		Usage:    "The address of the granter",
		Required: true,
		Value:    "",
	},
	granteeFlag: &cli.StringFlag{
		Name:     granteeFlag,
		Usage:    "The address of the grantee",
		Required: false,
		Value:    "",
	},
}

// aliases for defaultFlags
var aliases = map[string][]string{
	testnetFlag:     {"test", "t"},
	prvKeyFlag:      {"p", "prvKey"},
	allowedListFlag: {"allowed"},
	deniedListFlag:  {"denied"},
	expiredFlag:     {"exp"},
}
