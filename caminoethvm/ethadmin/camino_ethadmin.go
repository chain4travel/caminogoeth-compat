// Copyright (C) 2022-2023, Chain4Travel AG. All rights reserved.

package ethadmin

const (
	NOT_VERIFIED = 0b0000 // 0
	KYC_VERIFIED = 0b0001 // 1, bit 0
	KYC_EXPIRED  = 0b0010 // 2, bit 1
	KYB_VERIFIED = 0b1000 // 8, bit 3

	VERIFIED = KYC_VERIFIED | KYB_VERIFIED
)
