// Copyright (C) 2022-2025, Chain4Travel AG. All rights reserved.
//
// This file is a derived work, based on ava-labs code whose
// original notices appear below.
//
// It is distributed under the same license conditions as the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********************************************************
// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"errors"

	"github.com/chain4travel/caminogoeth-compat/caminogo/ids"
)

var errInvalidETHAddress = errors.New("invalid eth address")

type LockedAmount struct {
	Amount   uint64 `json:"amount"`
	Locktime uint64 `json:"locktime"`
}

type UnparsedAllocation struct {
	ETHAddr        string         `json:"ethAddr"`
	AVAXAddr       string         `json:"avaxAddr"`
	InitialAmount  uint64         `json:"initialAmount"`
	UnlockSchedule []LockedAmount `json:"unlockSchedule"`
}

type UnparsedStaker struct {
	NodeID        ids.NodeID `json:"nodeID"`
	RewardAddress string     `json:"rewardAddress"`
	DelegationFee uint32     `json:"delegationFee"`
}

// UnparsedConfig contains the genesis addresses used to construct a genesis
type UnparsedConfig struct {
	NetworkID uint32 `json:"networkID"`

	Allocations []UnparsedAllocation `json:"allocations"`

	StartTime                  uint64           `json:"startTime"`
	InitialStakeDuration       uint64           `json:"initialStakeDuration"`
	InitialStakeDurationOffset uint64           `json:"initialStakeDurationOffset"`
	InitialStakedFunds         []string         `json:"initialStakedFunds"`
	InitialStakers             []UnparsedStaker `json:"initialStakers"`
	Camino                     *UnparsedCamino  `json:"camino"`

	CChainGenesis string `json:"cChainGenesis"`

	Message string `json:"message"`
}
