package main

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/account"
	"github.com/LampardNguyen234/astra-go-sdk/client"
	"github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"time"
)

var (
	cosmosClient *client.CosmosClient
	testnet      bool
	host         string
	port         string
	chainID      = "astra_11110-1"
)

func defaultBeforeHook(_ *cli.Context) error {
	return initNetwork()
}

func initTxParams(c *cli.Context) (msg_params.TxParams, error) {
	prvKey := c.String(prvKeyFlag)
	if err := validatePrivateKey(prvKey); err != nil {
		return msg_params.DefaultTxParams(), err
	}
	return msg_params.TxParams{
		PrivateKey:    prvKey,
		GasLimit:      msg_params.DefaultTxParams().GasLimit,
		GasAdjustment: 1,
		GasPrice:      msg_params.DefaultTxParams().GasPrice,
	}, nil
}

func initNetwork() error {
	var err error

	cfg := client.DefaultMainnetConfig()
	if host+port != "" {
		cfg = client.CosmosClientConfig{
			Endpoint:       host,
			TendermintPort: port,
			ChainID:        chainID,
		}
	} else if testnet {
		cfg = client.DefaultTestnetConfig()
	}

	cosmosClient, err = client.NewCosmosClient(cfg)
	if err != nil {
		return fmt.Errorf("failed to initialize network: %v", err)
	}

	chainID = cfg.ChainID

	return nil
}

func validatePrivateKey(k string) error {
	if _, err := account.NewPrivateKeyFromString(k); err != nil {
		return errors.Wrapf(newAppError(errInvalidPrivateKey, err), k)
	}

	return nil
}

func validateAddress(addr string) error {
	if _, err := account.ParseCosmosAddress(addr); err != nil {
		return errors.Wrapf(newAppError(errInvalidAddress, err), addr)
	}

	return nil
}

func validateValidatorAddress(addr string) error {
	if _, err := account.ParseCosmosValidatorAddress(addr); err != nil {
		return errors.Wrapf(newAppError(errInvalidValidatorAddress, err), addr)
	}

	return nil
}

func parseDuration(d string) (time.Duration, error) {
	if d == "" {
		return 10 * 365 * 24 * time.Hour, nil
	}
	return time.ParseDuration(d)
}
