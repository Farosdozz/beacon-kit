// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package components

import (
	"cosmossdk.io/depinject"
	"github.com/berachain/beacon-kit/mod/config"
	"github.com/berachain/beacon-kit/mod/log"
	payloadbuilder "github.com/berachain/beacon-kit/mod/payload/pkg/builder"
	"github.com/berachain/beacon-kit/mod/payload/pkg/cache"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
)

// LocalBuilderInput is an input for the dep inject framework.
type LocalBuilderInput[
	BeaconStateT any,
	LoggerT log.AdvancedLogger[any, LoggerT],
] struct {
	depinject.In
	AttributesFactory AttributesFactory[BeaconStateT, *PayloadAttributes]
	Cfg               *config.Config
	ChainSpec         common.ChainSpec
	ExecutionEngine   *ExecutionEngine
	Logger            LoggerT
}

// ProvideLocalBuilder provides a local payload builder for the
// depinject framework.
func ProvideLocalBuilder[
	BeaconBlockHeaderT any,
	BeaconStateT BeaconState[
		BeaconStateT, BeaconBlockHeaderT, BeaconStateMarshallableT,
		*Eth1Data, *ExecutionPayloadHeader, *Fork, KVStoreT, *Validator,
		Validators, *Withdrawal,
	],
	BeaconStateMarshallableT any,
	KVStoreT any,
	LoggerT log.AdvancedLogger[any, LoggerT],
](
	in LocalBuilderInput[BeaconStateT, LoggerT],
) *payloadbuilder.PayloadBuilder[
	BeaconStateT, *ExecutionPayload, *ExecutionPayloadHeader,
	*PayloadAttributes, PayloadID, *Withdrawal,
] {
	return payloadbuilder.New[
		BeaconStateT, *ExecutionPayload, *ExecutionPayloadHeader,
		*PayloadAttributes, PayloadID, *Withdrawal,
	](
		&in.Cfg.PayloadBuilder,
		in.ChainSpec,
		in.Logger.With("service", "payload-builder"),
		in.ExecutionEngine,
		cache.NewPayloadIDCache[
			PayloadID,
			[32]byte, math.Slot,
		](),
		in.AttributesFactory,
	)
}