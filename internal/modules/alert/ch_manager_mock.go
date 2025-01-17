// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package alert

import (
	"sync"

	"github.com/balerter/balerter/internal/alert"
)

// chManagerMock is a mock implementation of chManager.
//
// 	func TestSomethingThatUseschManager(t *testing.T) {
//
// 		// make and configure a mocked chManager
// 		mockedchManager := &chManagerMock{
// 			SendFunc: func(a *alert.Alert, text string, options *alert.Options)  {
// 				panic("mock out the Send method")
// 			},
// 		}
//
// 		// use mockedchManager in code that requires chManager
// 		// and then make assertions.
//
// 	}
type chManagerMock struct {
	// SendFunc mocks the Send method.
	SendFunc func(a *alert.Alert, text string, options *alert.Options)

	// calls tracks calls to the methods.
	calls struct {
		// Send holds details about calls to the Send method.
		Send []struct {
			// A is the a argument value.
			A *alert.Alert
			// Text is the text argument value.
			Text string
			// Options is the options argument value.
			Options *alert.Options
		}
	}
	lockSend sync.RWMutex
}

// Send calls SendFunc.
func (mock *chManagerMock) Send(a *alert.Alert, text string, options *alert.Options) {
	if mock.SendFunc == nil {
		panic("chManagerMock.SendFunc: method is nil but chManager.Send was just called")
	}
	callInfo := struct {
		A       *alert.Alert
		Text    string
		Options *alert.Options
	}{
		A:       a,
		Text:    text,
		Options: options,
	}
	mock.lockSend.Lock()
	mock.calls.Send = append(mock.calls.Send, callInfo)
	mock.lockSend.Unlock()
	mock.SendFunc(a, text, options)
}

// SendCalls gets all the calls that were made to Send.
// Check the length with:
//     len(mockedchManager.SendCalls())
func (mock *chManagerMock) SendCalls() []struct {
	A       *alert.Alert
	Text    string
	Options *alert.Options
} {
	var calls []struct {
		A       *alert.Alert
		Text    string
		Options *alert.Options
	}
	mock.lockSend.RLock()
	calls = mock.calls.Send
	mock.lockSend.RUnlock()
	return calls
}
