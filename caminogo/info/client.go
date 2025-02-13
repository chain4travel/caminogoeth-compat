// Copyright (C) 2022-2024, Chain4Travel AG. All rights reserved.
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

package info

import (
	"context"

	// "github.com/ava-labs/avalanchego/utils/rpc"
	"github.com/chain4travel/caminogoeth-compat/caminogo/rpc"
)

var _ Client = (*client)(nil)

// Client interface for an Info API Client.
// See also AwaitBootstrapped.
type Client interface {
	IsBootstrapped(context.Context, string, ...rpc.Option) (bool, error)
}

// Client implementation for an Info API Client
type client struct {
	requester rpc.EndpointRequester
}

// NewClient returns a new Info API Client
func NewClient(uri string) Client {
	return &client{requester: rpc.NewEndpointRequester(
		uri + "/ext/info",
	)}
}

func (c *client) IsBootstrapped(ctx context.Context, chainID string, options ...rpc.Option) (bool, error) {
	res := &IsBootstrappedResponse{}
	err := c.requester.SendRequest(ctx, "info.isBootstrapped", &IsBootstrappedArgs{
		Chain: chainID,
	}, res, options...)
	return res.IsBootstrapped, err
}

// IsBootstrappedArgs are the arguments for calling IsBootstrapped
type IsBootstrappedArgs struct {
	// Alias of the chain
	// Can also be the string representation of the chain's ID
	Chain string `json:"chain"`
}

// IsBootstrappedResponse are the results from calling IsBootstrapped
type IsBootstrappedResponse struct {
	// True iff the chain exists and is done bootstrapping
	IsBootstrapped bool `json:"isBootstrapped"`
}
