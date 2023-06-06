There are two options for you to run the AtStella CLI by:
1. Downloading the pre-compiled executable binary file, you can find it in the [releases](https://github.com/LampardNguyen234/compound-cli/releases).
2. Compiling your own executable binary file from source as in the Installation instruction above.

Then execute the binary file with the following commands.

```shell
$ atstella-cli help
NAME:
   atstella-cli - A simple CLI application for doing weird stuff.

USAGE:
   compound-cli [global options] command [command options] [arguments...]

VERSION:
   v0.0.1

DESCRIPTION:
   A simple CLI for doing things that are beyond the capabilities of the regular SDK.

AUTHOR:
   AtStella Inc.

COMMANDS:
   help, h  Shows a list of commands or help for one command
   COMPOUND:
     compound, cmpd  Manage compound functionality

GLOBAL OPTIONS:
   --host value           Tendermint RPC host
   --port value           Tendermint RPC port
   --testnet, --test, -t  Whether to use testnet (default: false)
   --help, -h             show help (default: false)
   --version, -v          print the version (default: false)

COPYRIGHT:
   This tool is developed and maintained by the AtStella Devs Team. It is free for anyone. However, any commercial usages should be acknowledged by the AtStella Devs Team.
```
# Commands
<!-- commands -->
* [`COMPOUND`](#compound)
	* [`compound`](#compound)
		* [`compound register`](#compound_register)
		* [`compound unregister`](#compound_unregister)
## COMPOUND
### compound
This command helps perform compound-related actions
```shell
$ atstella-cli help compound
NAME:
   compound-cli compound - Manage compound functionality

USAGE:
   compound

CATEGORY:
   COMPOUND

DESCRIPTION:
   This command helps perform compound-related actions
```

#### compound_register
Register to the compound service
```shell
$ atstella-cli compound help register
NAME:
   atstella-cli compound register - Register to the compound service

USAGE:
   compound register --privateKey PRIVATE_KEY [--operator OPERATOR] [--allowedList ALLOWED_LIST] [--deniedList DENIED_LIST] [--expired EXPIRED]

   OPTIONAL flags are denoted by a [] bracket.

OPTIONS:
   --privateKey value, -p value, --prvKey value  The Astra private key [$PRIVATE_KEY]
   --operator value                              The address of the operator (default: "astra10y496hn28u6j8z80wuqvclta4jt9zcp7jfzqwa")
   --allowedList bonded, --allowed bonded        The list of allowed validator addresses (default: all bonded validators). Example: --allowed VALIDATOR_1 --allowed VALIDATOR_2
   --deniedList value, --denied value            The list of denied validator addresses (default: no denied validators). Example: --denied VALIDATOR_1 --denied VALIDATOR_2
   --expired value, --exp value                  The expiration duration (e.g, 1000s) (default: "87600h0m0s")
   
```

#### compound_unregister
Unregister from the compound service
```shell
$ atstella-cli compound help unregister
NAME:
   atstella-cli compound unregister - Unregister from the compound service

USAGE:
   compound unregister --privateKey PRIVATE_KEY [--operator OPERATOR]

   OPTIONAL flags are denoted by a [] bracket.

OPTIONS:
   --privateKey value, -p value, --prvKey value  The Astra private key [$PRIVATE_KEY]
   --operator value                              The address of the operator (default: "astra10y496hn28u6j8z80wuqvclta4jt9zcp7jfzqwa")
   
```

<!-- commandsstop -->
