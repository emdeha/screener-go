// Code generated by counterfeiter. DO NOT EDIT.
package importerfakes

import (
	"sync"

	"github.com/emdeha/screener-go/internal/company/importer"
)

type FakeEDGARClient struct {
	GetBulkDataStub        func() []byte
	getBulkDataMutex       sync.RWMutex
	getBulkDataArgsForCall []struct {
	}
	getBulkDataReturns struct {
		result1 []byte
	}
	getBulkDataReturnsOnCall map[int]struct {
		result1 []byte
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEDGARClient) GetBulkData() []byte {
	fake.getBulkDataMutex.Lock()
	ret, specificReturn := fake.getBulkDataReturnsOnCall[len(fake.getBulkDataArgsForCall)]
	fake.getBulkDataArgsForCall = append(fake.getBulkDataArgsForCall, struct {
	}{})
	fake.recordInvocation("GetBulkData", []interface{}{})
	fake.getBulkDataMutex.Unlock()
	if fake.GetBulkDataStub != nil {
		return fake.GetBulkDataStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getBulkDataReturns
	return fakeReturns.result1
}

func (fake *FakeEDGARClient) GetBulkDataCallCount() int {
	fake.getBulkDataMutex.RLock()
	defer fake.getBulkDataMutex.RUnlock()
	return len(fake.getBulkDataArgsForCall)
}

func (fake *FakeEDGARClient) GetBulkDataCalls(stub func() []byte) {
	fake.getBulkDataMutex.Lock()
	defer fake.getBulkDataMutex.Unlock()
	fake.GetBulkDataStub = stub
}

func (fake *FakeEDGARClient) GetBulkDataReturns(result1 []byte) {
	fake.getBulkDataMutex.Lock()
	defer fake.getBulkDataMutex.Unlock()
	fake.GetBulkDataStub = nil
	fake.getBulkDataReturns = struct {
		result1 []byte
	}{result1}
}

func (fake *FakeEDGARClient) GetBulkDataReturnsOnCall(i int, result1 []byte) {
	fake.getBulkDataMutex.Lock()
	defer fake.getBulkDataMutex.Unlock()
	fake.GetBulkDataStub = nil
	if fake.getBulkDataReturnsOnCall == nil {
		fake.getBulkDataReturnsOnCall = make(map[int]struct {
			result1 []byte
		})
	}
	fake.getBulkDataReturnsOnCall[i] = struct {
		result1 []byte
	}{result1}
}

func (fake *FakeEDGARClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getBulkDataMutex.RLock()
	defer fake.getBulkDataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEDGARClient) recordInvocation(key string, args []interface{}) {
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

var _ importer.EDGARClient = new(FakeEDGARClient)