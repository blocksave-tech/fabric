/*
Copyright DTCC 2016 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package java

import (
	"archive/tar"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"errors"

	"io/ioutil"

	"github.com/golang/protobuf/proto"
	"github.com/sinochem-tech/fabric/common/flogging"
	"github.com/sinochem-tech/fabric/common/util"
	ccutil "github.com/sinochem-tech/fabric/core/chaincode/platforms/util"
	pb "github.com/sinochem-tech/fabric/protos/peer"
)

var logger = flogging.MustGetLogger("java/hash")

//collectChaincodeFiles collects chaincode files and generates hashcode for the
//package.
//NOTE: for dev mode, user builds and runs chaincode manually. The name provided
//by the user is equivalent to the path. This method will treat the name
//as codebytes and compute the hash from it. ie, user cannot run the chaincode
//with the same (name, input, args)
func collectChaincodeFiles(spec *pb.ChaincodeSpec, tw *tar.Writer) (string, error) {
	if spec == nil {
		return "", errors.New("Cannot collect chaincode files from nil spec")
	}

	chaincodeID := spec.ChaincodeId
	if chaincodeID == nil || chaincodeID.Path == "" {
		return "", errors.New("Cannot collect chaincode files from empty chaincode path")
	}

	codepath := chaincodeID.Path

	var err error
	if !strings.HasPrefix(codepath, "/") {
		wd := ""
		wd, err = os.Getwd()
		codepath = wd + "/" + codepath
	}

	if err != nil {
		return "", fmt.Errorf("Error getting code %s", err)
	}

	var hash []byte

	//install will not have inputs and we don't have to collect hash for it
	if spec.Input == nil || len(spec.Input.Args) == 0 {
		logger.Debugf("not using input for hash computation for %v ", chaincodeID)
	} else {
		inputbytes, err2 := proto.Marshal(spec.Input)
		if err2 != nil {
			return "", fmt.Errorf("Error marshalling constructor: %s", err)
		}
		hash = util.GenerateHashFromSignature(codepath, inputbytes)
	}

	buf, err := ioutil.ReadFile(codepath)
	if err != nil {
		logger.Errorf("Error reading %s", err)
		return "", err
	}

	//get the new hash from file contents
	hash = ccutil.ComputeHash(buf, hash)

	return hex.EncodeToString(hash[:]), nil
}
