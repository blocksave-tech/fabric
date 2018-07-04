/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package container_test

import (
	"testing"

	"github.com/sinochem-tech/fabric/common/util"
	"github.com/sinochem-tech/fabric/core/container"
	pb "github.com/sinochem-tech/fabric/protos/peer"
	"github.com/stretchr/testify/assert"
)

func TestVM_GetChaincodePackageBytes(t *testing.T) {
	_, err := container.GetChaincodePackageBytes(nil)
	assert.Error(t, err,
		"GetChaincodePackageBytes did not return error when chaincode spec is nil")
	spec := &pb.ChaincodeSpec{ChaincodeId: nil}
	_, err = container.GetChaincodePackageBytes(spec)
	assert.Error(t, err, "Error expected when GetChaincodePackageBytes is called with nil chaincode ID")
	assert.Contains(t, err.Error(), "invalid chaincode spec")
	spec = &pb.ChaincodeSpec{Type: pb.ChaincodeSpec_GOLANG,
		ChaincodeId: nil,
		Input:       &pb.ChaincodeInput{Args: util.ToChaincodeArgs("f")}}
	_, err = container.GetChaincodePackageBytes(spec)
	assert.Error(t, err,
		"GetChaincodePackageBytes did not return error when chaincode ID is nil")
}
