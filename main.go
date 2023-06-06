package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
)

func main() {
	app := &cli.App{
		Name:        "atstella-cli",
		Usage:       "A simple CLI application for doing weird stuff.",
		Version:     "v0.0.1",
		Description: "A simple CLI for doing things that are beyond the capabilities of the regular SDK.",
		Authors: []*cli.Author{
			{
				Name: "AtStella Inc.",
			},
		},
		Copyright: "This tool is developed and maintained by the AtStella Devs Team. It is free for anyone. However, any " +
			"commercial usages should be acknowledged by the AtStella Devs Team.",
	}
	app.EnableBashCompletion = true
	app.Before = defaultBeforeHook

	// set app defaultFlags
	app.Flags = []cli.Flag{
		defaultFlags[testnetFlag],
		defaultFlags[hostFlag],
		defaultFlags[portFlag],
	}

	app.Commands = make([]*cli.Command, 0)
	app.Commands = append(app.Commands, compoundCommands...)

	for _, command := range app.Commands {
		if len(command.Subcommands) > 0 {
			sort.Sort(cli.CommandsByName(command.Subcommands))
			for _, subCommand := range command.Subcommands {
				buildUsageTextFromCommand(subCommand, command.Name)
			}
		}
		buildUsageTextFromCommand(command)
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	var err error
	//_ = generateDocsToFile(app, "commands.md") // un-comment this line to generate docs for the app's commands.
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
