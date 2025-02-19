// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package corestorage

import (
	"sync"
)

// CoreStorageMock is a mock implementation of CoreStorage.
//
// 	func TestSomethingThatUsesCoreStorage(t *testing.T) {
//
// 		// make and configure a mocked CoreStorage
// 		mockedCoreStorage := &CoreStorageMock{
// 			AlertFunc: func() Alert {
// 				panic("mock out the Alert method")
// 			},
// 			KVFunc: func() KV {
// 				panic("mock out the KV method")
// 			},
// 			NameFunc: func() string {
// 				panic("mock out the Name method")
// 			},
// 			StopFunc: func() error {
// 				panic("mock out the Stop method")
// 			},
// 		}
//
// 		// use mockedCoreStorage in code that requires CoreStorage
// 		// and then make assertions.
//
// 	}
type CoreStorageMock struct {
	// AlertFunc mocks the Alert method.
	AlertFunc func() Alert

	// KVFunc mocks the KV method.
	KVFunc func() KV

	// NameFunc mocks the Name method.
	NameFunc func() string

	// StopFunc mocks the Stop method.
	StopFunc func() error

	// calls tracks calls to the methods.
	calls struct {
		// Alert holds details about calls to the Alert method.
		Alert []struct {
		}
		// KV holds details about calls to the KV method.
		KV []struct {
		}
		// Name holds details about calls to the Name method.
		Name []struct {
		}
		// Stop holds details about calls to the Stop method.
		Stop []struct {
		}
	}
	lockAlert sync.RWMutex
	lockKV    sync.RWMutex
	lockName  sync.RWMutex
	lockStop  sync.RWMutex
}

// Alert calls AlertFunc.
func (mock *CoreStorageMock) Alert() Alert {
	if mock.AlertFunc == nil {
		panic("CoreStorageMock.AlertFunc: method is nil but CoreStorage.Alert was just called")
	}
	callInfo := struct {
	}{}
	mock.lockAlert.Lock()
	mock.calls.Alert = append(mock.calls.Alert, callInfo)
	mock.lockAlert.Unlock()
	return mock.AlertFunc()
}

// AlertCalls gets all the calls that were made to Alert.
// Check the length with:
//     len(mockedCoreStorage.AlertCalls())
func (mock *CoreStorageMock) AlertCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockAlert.RLock()
	calls = mock.calls.Alert
	mock.lockAlert.RUnlock()
	return calls
}

// KV calls KVFunc.
func (mock *CoreStorageMock) KV() KV {
	if mock.KVFunc == nil {
		panic("CoreStorageMock.KVFunc: method is nil but CoreStorage.KV was just called")
	}
	callInfo := struct {
	}{}
	mock.lockKV.Lock()
	mock.calls.KV = append(mock.calls.KV, callInfo)
	mock.lockKV.Unlock()
	return mock.KVFunc()
}

// KVCalls gets all the calls that were made to KV.
// Check the length with:
//     len(mockedCoreStorage.KVCalls())
func (mock *CoreStorageMock) KVCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockKV.RLock()
	calls = mock.calls.KV
	mock.lockKV.RUnlock()
	return calls
}

// Name calls NameFunc.
func (mock *CoreStorageMock) Name() string {
	if mock.NameFunc == nil {
		panic("CoreStorageMock.NameFunc: method is nil but CoreStorage.Name was just called")
	}
	callInfo := struct {
	}{}
	mock.lockName.Lock()
	mock.calls.Name = append(mock.calls.Name, callInfo)
	mock.lockName.Unlock()
	return mock.NameFunc()
}

// NameCalls gets all the calls that were made to Name.
// Check the length with:
//     len(mockedCoreStorage.NameCalls())
func (mock *CoreStorageMock) NameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockName.RLock()
	calls = mock.calls.Name
	mock.lockName.RUnlock()
	return calls
}

// Stop calls StopFunc.
func (mock *CoreStorageMock) Stop() error {
	if mock.StopFunc == nil {
		panic("CoreStorageMock.StopFunc: method is nil but CoreStorage.Stop was just called")
	}
	callInfo := struct {
	}{}
	mock.lockStop.Lock()
	mock.calls.Stop = append(mock.calls.Stop, callInfo)
	mock.lockStop.Unlock()
	return mock.StopFunc()
}

// StopCalls gets all the calls that were made to Stop.
// Check the length with:
//     len(mockedCoreStorage.StopCalls())
func (mock *CoreStorageMock) StopCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockStop.RLock()
	calls = mock.calls.Stop
	mock.lockStop.RUnlock()
	return calls
}
