/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package kafka

import (
	"github.com/Shopify/sarama"
	localconfig "github.com/sinochem-tech/fabric/orderer/common/localconfig"
	"github.com/sinochem-tech/fabric/orderer/consensus"
	cb "github.com/sinochem-tech/fabric/protos/common"
	logging "github.com/op/go-logging"
)

// New creates a Kafka-based consenter. Called by orderer's main.go.
func New(config localconfig.Kafka) consensus.Consenter {
	if config.Verbose {
		logging.SetLevel(logging.DEBUG, saramaLogID)
	}
	brokerConfig := newBrokerConfig(config.TLS, config.Retry, config.Version, defaultPartition)
	return &consenterImpl{
		brokerConfigVal: brokerConfig,
		tlsConfigVal:    config.TLS,
		retryOptionsVal: config.Retry,
		kafkaVersionVal: config.Version,
	}
}

// consenterImpl holds the implementation of type that satisfies the
// consensus.Consenter interface --as the HandleChain contract requires-- and
// the commonConsenter one.
type consenterImpl struct {
	brokerConfigVal *sarama.Config
	tlsConfigVal    localconfig.TLS
	retryOptionsVal localconfig.Retry
	kafkaVersionVal sarama.KafkaVersion
}

// HandleChain creates/returns a reference to a consensus.Chain object for the
// given set of support resources. Implements the consensus.Consenter
// interface. Called by consensus.newChainSupport(), which is itself called by
// multichannel.NewManagerImpl() when ranging over the ledgerFactory's
// existingChains.
func (consenter *consenterImpl) HandleChain(support consensus.ConsenterSupport, metadata *cb.Metadata) (consensus.Chain, error) {
	lastOffsetPersisted, lastOriginalOffsetProcessed, lastResubmittedConfigOffset := getOffsets(metadata.Value, support.ChainID())
	return newChain(consenter, support, lastOffsetPersisted, lastOriginalOffsetProcessed, lastResubmittedConfigOffset)
}

// commonConsenter allows us to retrieve the configuration options set on the
// consenter object. These will be common across all chain objects derived by
// this consenter. They are set using using local configuration settings. This
// interface is satisfied by consenterImpl.
type commonConsenter interface {
	brokerConfig() *sarama.Config
	retryOptions() localconfig.Retry
}

func (consenter *consenterImpl) brokerConfig() *sarama.Config {
	return consenter.brokerConfigVal
}

func (consenter *consenterImpl) retryOptions() localconfig.Retry {
	return consenter.retryOptionsVal
}

// closeable allows the shut down of the calling resource.
type closeable interface {
	close() error
}
