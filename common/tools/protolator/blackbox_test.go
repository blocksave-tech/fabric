/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package protolator_test

import (
	"bytes"
	"testing"

	"github.com/sinochem-tech/fabric/common/tools/configtxgen/configtxgentest"
	"github.com/sinochem-tech/fabric/common/tools/configtxgen/encoder"
	genesisconfig "github.com/sinochem-tech/fabric/common/tools/configtxgen/localconfig"
	. "github.com/sinochem-tech/fabric/common/tools/protolator"
	cb "github.com/sinochem-tech/fabric/protos/common"
	"github.com/sinochem-tech/fabric/protos/msp"
	"github.com/sinochem-tech/fabric/protos/utils"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func bidirectionalMarshal(t *testing.T, doc proto.Message) {
	var buffer bytes.Buffer

	assert.NoError(t, DeepMarshalJSON(&buffer, doc))

	newRoot := proto.Clone(doc)
	newRoot.Reset()
	assert.NoError(t, DeepUnmarshalJSON(bytes.NewReader(buffer.Bytes()), newRoot))

	// Note, we cannot do an equality check between newRoot and sampleDoc
	// because of the nondeterministic nature of binary proto marshaling
	// So instead we re-marshal to JSON which is a deterministic marshaling
	// and compare equality there instead

	//t.Log(doc)
	//t.Log(newRoot)

	var remarshaled bytes.Buffer
	assert.NoError(t, DeepMarshalJSON(&remarshaled, newRoot))
	assert.Equal(t, string(buffer.Bytes()), string(remarshaled.Bytes()))
	//t.Log(string(buffer.Bytes()))
	//t.Log(string(remarshaled.Bytes()))
}

func TestConfigUpdate(t *testing.T) {
	cg, err := encoder.NewChannelGroup(configtxgentest.Load(genesisconfig.SampleSingleMSPSoloProfile))
	assert.NoError(t, err)

	bidirectionalMarshal(t, &cb.ConfigUpdateEnvelope{
		ConfigUpdate: utils.MarshalOrPanic(&cb.ConfigUpdate{
			ReadSet:  cg,
			WriteSet: cg,
		}),
	})
}

func TestIdemix(t *testing.T) {
	bidirectionalMarshal(t, &msp.MSPConfig{
		Type: 1,
		Config: utils.MarshalOrPanic(&msp.IdemixMSPConfig{
			Name: "fooo",
		}),
	})
}

func TestGenesisBlock(t *testing.T) {
	p := encoder.New(configtxgentest.Load(genesisconfig.SampleSingleMSPSoloProfile))
	gb := p.GenesisBlockForChannel("foo")

	bidirectionalMarshal(t, gb)
}
