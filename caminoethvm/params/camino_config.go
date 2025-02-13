// Copyright (C) 2022-2025, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package params

import (
	"math/big"

	"github.com/chain4travel/caminogoeth-compat/caminoethvm/utils"
)

// Gas Price
const (
	SunrisePhase0ExtraDataSize        = 0
	SunrisePhase0BaseFee       uint64 = 200_000_000_000
)

// Camino ChainIDs
var (
	// CaminoChainID ...
	CaminoChainID = big.NewInt(500)
	// CaminoChainID ...
	ColumbusChainID = big.NewInt(501)
	// KopernikusChainID ...
	KopernikusChainID = big.NewInt(502)
)

// IsSunrisePhase0 returns whether [blockTimestamp] represents a block
// with a timestamp after the Sunrise Phase 0 upgrade time.
func (c *ChainConfig) IsSunrisePhase0(time *big.Int) bool {
	return utils.IsForked(c.SunrisePhase0BlockTimestamp, time)
}

// IsBerlin returns whether [time] represents a block
// with a timestamp after the Berlin upgrade time.
func (c *ChainConfig) IsBerlin(time *big.Int) bool {
	return utils.IsForked(c.BerlinBlockTimestamp, time)
}
