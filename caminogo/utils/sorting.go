// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"bytes"

	"golang.org/x/exp/slices"

	// "github.com/ava-labs/avalanchego/utils/hashing"
	"github.com/chain4travel/caminogoeth-compat/caminogo/hashing"
)

// TODO can we handle sorting where the Less function relies on a codec?

type Sortable[T any] interface {
	Compare(T) int
}

// Sorts the elements of [s].
func Sort[T Sortable[T]](s []T) {
	slices.SortFunc(s, func(i, j T) int {
		return i.Compare(j)
	})
}

// Sorts the elements of [s] based on their hashes.
func SortByHash[T ~[]byte](s []T) {
	slices.SortFunc(s, func(i, j T) int {
		iHash := hashing.ComputeHash256(i)
		jHash := hashing.ComputeHash256(j)
		return bytes.Compare(iHash, jHash)
	})
}

// Sorts a 2D byte slice.
// Each byte slice is not sorted internally; the byte slices are sorted relative
// to one another.
func SortBytes[T ~[]byte](arr []T) {
	slices.SortFunc(arr, func(i, j T) int {
		return bytes.Compare(i, j)
	})
}
