// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package runner

import (
	"sync"

	"github.com/balerter/balerter/internal/modules"
)

// dsManagerMock is a mock implementation of dsManager.
//
// 	func TestSomethingThatUsesdsManager(t *testing.T) {
//
// 		// make and configure a mocked dsManager
// 		mockeddsManager := &dsManagerMock{
// 			GetFunc: func() []modules.Module {
// 				panic("mock out the Get method")
// 			},
// 		}
//
// 		// use mockeddsManager in code that requires dsManager
// 		// and then make assertions.
//
// 	}
type dsManagerMock struct {
	// GetFunc mocks the Get method.
	GetFunc func() []modules.Module

	// calls tracks calls to the methods.
	calls struct {
		// Get holds details about calls to the Get method.
		Get []struct {
		}
	}
	lockGet sync.RWMutex
}

// Get calls GetFunc.
func (mock *dsManagerMock) Get() []modules.Module {
	if mock.GetFunc == nil {
		panic("dsManagerMock.GetFunc: method is nil but dsManager.Get was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc()
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockeddsManager.GetCalls())
func (mock *dsManagerMock) GetCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}