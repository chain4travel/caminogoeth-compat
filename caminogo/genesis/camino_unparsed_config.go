// Copyright (C) 2022-2024, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"errors"
)

var errCannotParseInitialAdmin = errors.New("cannot parse initialAdmin from genesis")

type UnparsedCamino struct {
	VerifyNodeSignature      bool                       `json:"verifyNodeSignature"`
	LockModeBondDeposit      bool                       `json:"lockModeBondDeposit"`
	InitialAdmin             string                     `json:"initialAdmin"`
	DepositOffers            []UnparsedDepositOffer     `json:"depositOffers"`
	Allocations              []UnparsedCaminoAllocation `json:"allocations"`
	InitialMultisigAddresses []UnparsedMultisigAlias    `json:"initialMultisigAddresses"`
}

type UnparsedCaminoAllocation struct {
	ETHAddr             string                       `json:"ethAddr"`
	AVAXAddr            string                       `json:"avaxAddr"`
	XAmount             uint64                       `json:"xAmount"`
	AddressStates       AddressStates                `json:"addressStates"`
	PlatformAllocations []UnparsedPlatformAllocation `json:"platformAllocations"`
}
type UnparsedPlatformAllocation struct {
	Amount            uint64 `json:"amount"`
	NodeID            string `json:"nodeID,omitempty"`
	ValidatorDuration uint64 `json:"validatorDuration,omitempty"`
	DepositDuration   uint64 `json:"depositDuration,omitempty"`
	TimestampOffset   uint64 `json:"timestampOffset,omitempty"`
	DepositOfferMemo  string `json:"depositOfferMemo,omitempty"`
	Memo              string `json:"memo,omitempty"`
}

// UnparsedMultisigAlias defines a multisignature alias address.
// [Alias] is the alias of the multisignature address. It's encoded to string
// the same way as ShortID String() method does.
// [Addresses] are the addresses that are allowed to sign transactions from the multisignature address.
// All addresses are encoded to string the same way as ShortID String() method does.
// [Threshold] is the number of signatures required to sign transactions from the multisignature address.
type UnparsedMultisigAlias struct {
	Alias     string   `json:"alias"`
	Addresses []string `json:"addresses"`
	Threshold uint32   `json:"threshold"`
	Memo      string   `json:"memo,omitempty"`
}

type UnparsedDepositOffer struct {
	InterestRateNominator   uint64                    `json:"interestRateNominator"`
	StartOffset             uint64                    `json:"startOffset"`
	EndOffset               uint64                    `json:"endOffset"`
	MinAmount               uint64                    `json:"minAmount"`
	MinDuration             uint32                    `json:"minDuration"`
	MaxDuration             uint32                    `json:"maxDuration"`
	UnlockPeriodDuration    uint32                    `json:"unlockPeriodDuration"`
	NoRewardsPeriodDuration uint32                    `json:"noRewardsPeriodDuration"`
	Memo                    string                    `json:"memo"`
	Flags                   UnparsedDepositOfferFlags `json:"flags"`
}

type UnparsedDepositOfferFlags struct {
	Locked bool `json:"locked"`
}
