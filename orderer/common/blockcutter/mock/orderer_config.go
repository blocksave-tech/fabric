// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"
	"time"

	"github.com/sinochem-tech/fabric/common/channelconfig"
	ab "github.com/sinochem-tech/fabric/protos/orderer"
)

type OrdererConfig struct {
	ConsensusTypeStub        func() string
	consensusTypeMutex       sync.RWMutex
	consensusTypeArgsForCall []struct{}
	consensusTypeReturns     struct {
		result1 string
	}
	consensusTypeReturnsOnCall map[int]struct {
		result1 string
	}
	BatchSizeStub        func() *ab.BatchSize
	batchSizeMutex       sync.RWMutex
	batchSizeArgsForCall []struct{}
	batchSizeReturns     struct {
		result1 *ab.BatchSize
	}
	batchSizeReturnsOnCall map[int]struct {
		result1 *ab.BatchSize
	}
	BatchTimeoutStub        func() time.Duration
	batchTimeoutMutex       sync.RWMutex
	batchTimeoutArgsForCall []struct{}
	batchTimeoutReturns     struct {
		result1 time.Duration
	}
	batchTimeoutReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	MaxChannelsCountStub        func() uint64
	maxChannelsCountMutex       sync.RWMutex
	maxChannelsCountArgsForCall []struct{}
	maxChannelsCountReturns     struct {
		result1 uint64
	}
	maxChannelsCountReturnsOnCall map[int]struct {
		result1 uint64
	}
	KafkaBrokersStub        func() []string
	kafkaBrokersMutex       sync.RWMutex
	kafkaBrokersArgsForCall []struct{}
	kafkaBrokersReturns     struct {
		result1 []string
	}
	kafkaBrokersReturnsOnCall map[int]struct {
		result1 []string
	}
	OrganizationsStub        func() map[string]channelconfig.Org
	organizationsMutex       sync.RWMutex
	organizationsArgsForCall []struct{}
	organizationsReturns     struct {
		result1 map[string]channelconfig.Org
	}
	organizationsReturnsOnCall map[int]struct {
		result1 map[string]channelconfig.Org
	}
	CapabilitiesStub        func() channelconfig.OrdererCapabilities
	capabilitiesMutex       sync.RWMutex
	capabilitiesArgsForCall []struct{}
	capabilitiesReturns     struct {
		result1 channelconfig.OrdererCapabilities
	}
	capabilitiesReturnsOnCall map[int]struct {
		result1 channelconfig.OrdererCapabilities
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *OrdererConfig) ConsensusType() string {
	fake.consensusTypeMutex.Lock()
	ret, specificReturn := fake.consensusTypeReturnsOnCall[len(fake.consensusTypeArgsForCall)]
	fake.consensusTypeArgsForCall = append(fake.consensusTypeArgsForCall, struct{}{})
	fake.recordInvocation("ConsensusType", []interface{}{})
	fake.consensusTypeMutex.Unlock()
	if fake.ConsensusTypeStub != nil {
		return fake.ConsensusTypeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.consensusTypeReturns.result1
}

func (fake *OrdererConfig) ConsensusTypeCallCount() int {
	fake.consensusTypeMutex.RLock()
	defer fake.consensusTypeMutex.RUnlock()
	return len(fake.consensusTypeArgsForCall)
}

func (fake *OrdererConfig) ConsensusTypeReturns(result1 string) {
	fake.ConsensusTypeStub = nil
	fake.consensusTypeReturns = struct {
		result1 string
	}{result1}
}

func (fake *OrdererConfig) ConsensusTypeReturnsOnCall(i int, result1 string) {
	fake.ConsensusTypeStub = nil
	if fake.consensusTypeReturnsOnCall == nil {
		fake.consensusTypeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.consensusTypeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *OrdererConfig) BatchSize() *ab.BatchSize {
	fake.batchSizeMutex.Lock()
	ret, specificReturn := fake.batchSizeReturnsOnCall[len(fake.batchSizeArgsForCall)]
	fake.batchSizeArgsForCall = append(fake.batchSizeArgsForCall, struct{}{})
	fake.recordInvocation("BatchSize", []interface{}{})
	fake.batchSizeMutex.Unlock()
	if fake.BatchSizeStub != nil {
		return fake.BatchSizeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.batchSizeReturns.result1
}

func (fake *OrdererConfig) BatchSizeCallCount() int {
	fake.batchSizeMutex.RLock()
	defer fake.batchSizeMutex.RUnlock()
	return len(fake.batchSizeArgsForCall)
}

func (fake *OrdererConfig) BatchSizeReturns(result1 *ab.BatchSize) {
	fake.BatchSizeStub = nil
	fake.batchSizeReturns = struct {
		result1 *ab.BatchSize
	}{result1}
}

func (fake *OrdererConfig) BatchSizeReturnsOnCall(i int, result1 *ab.BatchSize) {
	fake.BatchSizeStub = nil
	if fake.batchSizeReturnsOnCall == nil {
		fake.batchSizeReturnsOnCall = make(map[int]struct {
			result1 *ab.BatchSize
		})
	}
	fake.batchSizeReturnsOnCall[i] = struct {
		result1 *ab.BatchSize
	}{result1}
}

func (fake *OrdererConfig) BatchTimeout() time.Duration {
	fake.batchTimeoutMutex.Lock()
	ret, specificReturn := fake.batchTimeoutReturnsOnCall[len(fake.batchTimeoutArgsForCall)]
	fake.batchTimeoutArgsForCall = append(fake.batchTimeoutArgsForCall, struct{}{})
	fake.recordInvocation("BatchTimeout", []interface{}{})
	fake.batchTimeoutMutex.Unlock()
	if fake.BatchTimeoutStub != nil {
		return fake.BatchTimeoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.batchTimeoutReturns.result1
}

func (fake *OrdererConfig) BatchTimeoutCallCount() int {
	fake.batchTimeoutMutex.RLock()
	defer fake.batchTimeoutMutex.RUnlock()
	return len(fake.batchTimeoutArgsForCall)
}

func (fake *OrdererConfig) BatchTimeoutReturns(result1 time.Duration) {
	fake.BatchTimeoutStub = nil
	fake.batchTimeoutReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *OrdererConfig) BatchTimeoutReturnsOnCall(i int, result1 time.Duration) {
	fake.BatchTimeoutStub = nil
	if fake.batchTimeoutReturnsOnCall == nil {
		fake.batchTimeoutReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.batchTimeoutReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *OrdererConfig) MaxChannelsCount() uint64 {
	fake.maxChannelsCountMutex.Lock()
	ret, specificReturn := fake.maxChannelsCountReturnsOnCall[len(fake.maxChannelsCountArgsForCall)]
	fake.maxChannelsCountArgsForCall = append(fake.maxChannelsCountArgsForCall, struct{}{})
	fake.recordInvocation("MaxChannelsCount", []interface{}{})
	fake.maxChannelsCountMutex.Unlock()
	if fake.MaxChannelsCountStub != nil {
		return fake.MaxChannelsCountStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.maxChannelsCountReturns.result1
}

func (fake *OrdererConfig) MaxChannelsCountCallCount() int {
	fake.maxChannelsCountMutex.RLock()
	defer fake.maxChannelsCountMutex.RUnlock()
	return len(fake.maxChannelsCountArgsForCall)
}

func (fake *OrdererConfig) MaxChannelsCountReturns(result1 uint64) {
	fake.MaxChannelsCountStub = nil
	fake.maxChannelsCountReturns = struct {
		result1 uint64
	}{result1}
}

func (fake *OrdererConfig) MaxChannelsCountReturnsOnCall(i int, result1 uint64) {
	fake.MaxChannelsCountStub = nil
	if fake.maxChannelsCountReturnsOnCall == nil {
		fake.maxChannelsCountReturnsOnCall = make(map[int]struct {
			result1 uint64
		})
	}
	fake.maxChannelsCountReturnsOnCall[i] = struct {
		result1 uint64
	}{result1}
}

func (fake *OrdererConfig) KafkaBrokers() []string {
	fake.kafkaBrokersMutex.Lock()
	ret, specificReturn := fake.kafkaBrokersReturnsOnCall[len(fake.kafkaBrokersArgsForCall)]
	fake.kafkaBrokersArgsForCall = append(fake.kafkaBrokersArgsForCall, struct{}{})
	fake.recordInvocation("KafkaBrokers", []interface{}{})
	fake.kafkaBrokersMutex.Unlock()
	if fake.KafkaBrokersStub != nil {
		return fake.KafkaBrokersStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.kafkaBrokersReturns.result1
}

func (fake *OrdererConfig) KafkaBrokersCallCount() int {
	fake.kafkaBrokersMutex.RLock()
	defer fake.kafkaBrokersMutex.RUnlock()
	return len(fake.kafkaBrokersArgsForCall)
}

func (fake *OrdererConfig) KafkaBrokersReturns(result1 []string) {
	fake.KafkaBrokersStub = nil
	fake.kafkaBrokersReturns = struct {
		result1 []string
	}{result1}
}

func (fake *OrdererConfig) KafkaBrokersReturnsOnCall(i int, result1 []string) {
	fake.KafkaBrokersStub = nil
	if fake.kafkaBrokersReturnsOnCall == nil {
		fake.kafkaBrokersReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.kafkaBrokersReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *OrdererConfig) Organizations() map[string]channelconfig.Org {
	fake.organizationsMutex.Lock()
	ret, specificReturn := fake.organizationsReturnsOnCall[len(fake.organizationsArgsForCall)]
	fake.organizationsArgsForCall = append(fake.organizationsArgsForCall, struct{}{})
	fake.recordInvocation("Organizations", []interface{}{})
	fake.organizationsMutex.Unlock()
	if fake.OrganizationsStub != nil {
		return fake.OrganizationsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.organizationsReturns.result1
}

func (fake *OrdererConfig) OrganizationsCallCount() int {
	fake.organizationsMutex.RLock()
	defer fake.organizationsMutex.RUnlock()
	return len(fake.organizationsArgsForCall)
}

func (fake *OrdererConfig) OrganizationsReturns(result1 map[string]channelconfig.Org) {
	fake.OrganizationsStub = nil
	fake.organizationsReturns = struct {
		result1 map[string]channelconfig.Org
	}{result1}
}

func (fake *OrdererConfig) OrganizationsReturnsOnCall(i int, result1 map[string]channelconfig.Org) {
	fake.OrganizationsStub = nil
	if fake.organizationsReturnsOnCall == nil {
		fake.organizationsReturnsOnCall = make(map[int]struct {
			result1 map[string]channelconfig.Org
		})
	}
	fake.organizationsReturnsOnCall[i] = struct {
		result1 map[string]channelconfig.Org
	}{result1}
}

func (fake *OrdererConfig) Capabilities() channelconfig.OrdererCapabilities {
	fake.capabilitiesMutex.Lock()
	ret, specificReturn := fake.capabilitiesReturnsOnCall[len(fake.capabilitiesArgsForCall)]
	fake.capabilitiesArgsForCall = append(fake.capabilitiesArgsForCall, struct{}{})
	fake.recordInvocation("Capabilities", []interface{}{})
	fake.capabilitiesMutex.Unlock()
	if fake.CapabilitiesStub != nil {
		return fake.CapabilitiesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.capabilitiesReturns.result1
}

func (fake *OrdererConfig) CapabilitiesCallCount() int {
	fake.capabilitiesMutex.RLock()
	defer fake.capabilitiesMutex.RUnlock()
	return len(fake.capabilitiesArgsForCall)
}

func (fake *OrdererConfig) CapabilitiesReturns(result1 channelconfig.OrdererCapabilities) {
	fake.CapabilitiesStub = nil
	fake.capabilitiesReturns = struct {
		result1 channelconfig.OrdererCapabilities
	}{result1}
}

func (fake *OrdererConfig) CapabilitiesReturnsOnCall(i int, result1 channelconfig.OrdererCapabilities) {
	fake.CapabilitiesStub = nil
	if fake.capabilitiesReturnsOnCall == nil {
		fake.capabilitiesReturnsOnCall = make(map[int]struct {
			result1 channelconfig.OrdererCapabilities
		})
	}
	fake.capabilitiesReturnsOnCall[i] = struct {
		result1 channelconfig.OrdererCapabilities
	}{result1}
}

func (fake *OrdererConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.consensusTypeMutex.RLock()
	defer fake.consensusTypeMutex.RUnlock()
	fake.batchSizeMutex.RLock()
	defer fake.batchSizeMutex.RUnlock()
	fake.batchTimeoutMutex.RLock()
	defer fake.batchTimeoutMutex.RUnlock()
	fake.maxChannelsCountMutex.RLock()
	defer fake.maxChannelsCountMutex.RUnlock()
	fake.kafkaBrokersMutex.RLock()
	defer fake.kafkaBrokersMutex.RUnlock()
	fake.organizationsMutex.RLock()
	defer fake.organizationsMutex.RUnlock()
	fake.capabilitiesMutex.RLock()
	defer fake.capabilitiesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *OrdererConfig) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
