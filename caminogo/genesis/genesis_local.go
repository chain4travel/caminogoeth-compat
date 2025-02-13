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
	"time"

	_ "embed"

	"github.com/chain4travel/caminogoeth-compat/caminogo/units"
)

var LocalParams = Params{
	TxFeeConfig: TxFeeConfig{
		TxFee:                         units.MilliAvax,
		CreateAssetTxFee:              units.MilliAvax,
		CreateSubnetTxFee:             100 * units.MilliAvax,
		TransformSubnetTxFee:          100 * units.MilliAvax,
		CreateBlockchainTxFee:         100 * units.MilliAvax,
		AddPrimaryNetworkValidatorFee: 0,
		AddPrimaryNetworkDelegatorFee: 0,
		AddSubnetValidatorFee:         units.MilliAvax,
		AddSubnetDelegatorFee:         units.MilliAvax,
	},
	StakingConfig: StakingConfig{
		UptimeRequirement: .8, // 80%
		MinValidatorStake: 2 * units.KiloAvax,
		MaxValidatorStake: 3 * units.MegaAvax,
		MinDelegatorStake: 25 * units.Avax,
		MinDelegationFee:  20000, // 2%
		MinStakeDuration:  24 * time.Hour,
		MaxStakeDuration:  365 * 24 * time.Hour,
	},
}
