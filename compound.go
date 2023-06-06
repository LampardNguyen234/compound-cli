package main

import (
	"github.com/LampardNguyen234/astra-go-sdk/account"
	"github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	distrTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"time"
)

// registerCompound performs compound registration.
func registerCompound(c *cli.Context) error {
	txParams, err := initTxParams(c)
	if err != nil {
		return err
	}
	operator := c.String(operatorFlag)
	if err = validateAddress(operator); err != nil {
		return err
	}
	allowed := c.StringSlice(allowedListFlag)
	if len(allowed) > 0 {
		for _, addr := range allowed {
			if err = validateValidatorAddress(addr); err != nil {
				return err
			}
		}
	}
	denied := c.StringSlice(deniedListFlag)
	if len(denied) > 0 {
		for _, addr := range denied {
			if err = validateValidatorAddress(addr); err != nil {
				return err
			}
		}
	}
	if len(allowed) > 0 && len(denied) > 0 {
		return errors.Wrapf(newAppError(errUnexpected), "expect either `allowed` or `denied` to be empty")
	}
	expired, err := parseDuration(c.String(expiredFlag))
	if err != nil {
		return errors.Wrapf(newAppError(errInvalidDuration), err.Error())
	}

	if len(allowed) == 0 && len(denied) == 0 {
		allValidators, err := cosmosClient.AllValidators(stakingTypes.Bonded)
		if err != nil {
			return errors.Wrapf(newAppError(errClient), err.Error())
		}

		for _, val := range allValidators {
			allowed = append(allowed, val.OperatorAddress)
		}
	}

	allowedVals := make([]sdk.ValAddress, 0)
	for _, addr := range allowed {
		allowedVals = append(allowedVals, account.MustParseCosmosValidatorAddress(addr))
	}
	deniedVals := make([]sdk.ValAddress, 0)
	for _, addr := range denied {
		deniedVals = append(deniedVals, account.MustParseCosmosValidatorAddress(addr))
	}

	stakeGrant, err := stakingTypes.NewStakeAuthorization(
		allowedVals, deniedVals,
		stakingTypes.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE, nil,
	)
	if err != nil {
		return errors.Wrapf(newAppError(errUnexpected), err.Error())
	}
	txResp, err := cosmosClient.TxGrantAuthorization(
		msg_params.TxGrantParams{
			TxParams:    txParams,
			Grantee:     operator,
			ExpiredTime: time.Now().Add(expired),
		},
		authz.NewGenericAuthorization(sdk.MsgTypeURL(&distrTypes.MsgWithdrawDelegatorReward{})),
		stakeGrant,
	)
	if err != nil {
		return errors.Wrapf(newAppError(errCreateTx), err.Error())
	}

	return jsonPrintWithKey("TxHash", txResp.TxHash)
}

// unregisterCompound performs compound un-registration.
func unregisterCompound(c *cli.Context) error {
	txParams, err := initTxParams(c)
	if err != nil {
		return err
	}
	operator := c.String(operatorFlag)
	if err = validateAddress(operator); err != nil {
		return err
	}
	operatorAcc := account.MustParseCosmosAddress(operator)

	p, _ := account.NewPrivateKeyFromString(txParams.PrivateKey)
	msgs := make([]sdk.Msg, 0)
	withdrawMsg := authz.NewMsgRevoke(p.AccAddress(), operatorAcc, sdk.MsgTypeURL(&distrTypes.MsgWithdrawDelegatorReward{}))
	msgs = append(msgs, &withdrawMsg)
	delegateMsg := authz.NewMsgRevoke(p.AccAddress(), operatorAcc, sdk.MsgTypeURL(&stakingTypes.MsgDelegate{}))
	msgs = append(msgs, &delegateMsg)

	txResp, err := cosmosClient.BuildAndSendTx(
		txParams, msgs...,
	)
	if err != nil {
		return errors.Wrapf(newAppError(errCreateTx), err.Error())
	}

	return jsonPrintWithKey("TxHash", txResp.TxHash)
}
