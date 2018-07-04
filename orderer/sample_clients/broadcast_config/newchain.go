// Copyright IBM Corp. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/sinochem-tech/fabric/common/localmsp"
	"github.com/sinochem-tech/fabric/common/tools/configtxgen/encoder"
	genesisconfig "github.com/sinochem-tech/fabric/common/tools/configtxgen/localconfig"
	cb "github.com/sinochem-tech/fabric/protos/common"
)

func newChainRequest(consensusType, creationPolicy, newChannelID string) *cb.Envelope {
	env, err := encoder.MakeChannelCreationTransaction(newChannelID, localmsp.NewSigner(), nil, genesisconfig.Load(genesisconfig.SampleSingleMSPChannelProfile))
	if err != nil {
		panic(err)
	}
	return env
}
