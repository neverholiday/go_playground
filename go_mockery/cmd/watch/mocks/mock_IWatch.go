// Code generated by mockery v2.48.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MockIWatch is an autogenerated mock type for the IWatch type
type MockIWatch struct {
	mock.Mock
}

type MockIWatch_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIWatch) EXPECT() *MockIWatch_Expecter {
	return &MockIWatch_Expecter{mock: &_m.Mock}
}

// GetCurrentTime provides a mock function with given fields:
func (_m *MockIWatch) GetCurrentTime() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetCurrentTime")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIWatch_GetCurrentTime_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCurrentTime'
type MockIWatch_GetCurrentTime_Call struct {
	*mock.Call
}

// GetCurrentTime is a helper method to define mock.On call
func (_e *MockIWatch_Expecter) GetCurrentTime() *MockIWatch_GetCurrentTime_Call {
	return &MockIWatch_GetCurrentTime_Call{Call: _e.mock.On("GetCurrentTime")}
}

func (_c *MockIWatch_GetCurrentTime_Call) Run(run func()) *MockIWatch_GetCurrentTime_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIWatch_GetCurrentTime_Call) Return(_a0 string, _a1 error) *MockIWatch_GetCurrentTime_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIWatch_GetCurrentTime_Call) RunAndReturn(run func() (string, error)) *MockIWatch_GetCurrentTime_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIWatch creates a new instance of MockIWatch. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIWatch(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIWatch {
	mock := &MockIWatch{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
