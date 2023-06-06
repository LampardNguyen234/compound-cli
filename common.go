package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/camelcase"
	"github.com/urfave/cli/v2"
	"log"
	"strings"
)

func jsonPrintWithKey(key string, val interface{}) error {
	return jsonPrint(map[string]interface{}{key: val})
}

func jsonPrint(val interface{}) error {
	jsb, err := json.MarshalIndent(val, "", "\t")
	if err != nil {
		return err
	}
	log.Println(string(jsb))
	return nil
}

// flagToVariable gets the variable representation for a flag.
// The variable representation of a flag is a ALL_UPPER_CASE form of a flag.
//
// For example, the variable resp of the flag `privateKey` is `PRIVATE_KEY`.
func flagToVariable(f string) string {
	f = strings.Replace(f, "Flag", "", 1)

	words := camelcase.Split(f)
	res := ""
	for _, word := range words {
		if res == "" {
			res += strings.ToUpper(word)
		} else {
			res += "_" + strings.ToUpper(word)
		}
	}

	return res
}

func buildUsageTextFromCommand(command *cli.Command, parents ...string) {
	parent := ""
	if len(parents) > 0 {
		parent = parents[0]
	}
	res := command.Name
	hasOptionalFlags := false
	for _, f := range command.Flags {
		flagString := fmt.Sprintf(" --%v %v", f.Names()[0], flagToVariable(f.Names()[0]))
		if requiredFlag, ok := f.(cli.RequiredFlag); ok {
			if !requiredFlag.IsRequired() {
				// optional flag is put inside a [] symbol.
				flagString = fmt.Sprintf(" [--%v %v]", f.Names()[0], flagToVariable(f.Names()[0]))
				hasOptionalFlags = true
			}
		}
		res += flagString
	}
	if parent != "" {
		res = fmt.Sprintf("%v %v", parent, res)
	}

	command.UsageText = res
	if hasOptionalFlags {
		command.UsageText += "\n\n\t OPTIONAL flags are denoted by a [] bracket."
	}
}
