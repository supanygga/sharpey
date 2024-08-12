// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// TwitchClient is an autogenerated mock type for the TwitchClient type
type TwitchClient struct {
	mock.Mock
}

type TwitchClient_Expecter struct {
	mock *mock.Mock
}

func (_m *TwitchClient) EXPECT() *TwitchClient_Expecter {
	return &TwitchClient_Expecter{mock: &_m.Mock}
}

// Join provides a mock function with given fields: channels
func (_m *TwitchClient) Join(channels ...string) error {
	_va := make([]interface{}, len(channels))
	for _i := range channels {
		_va[_i] = channels[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Join")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...string) error); ok {
		r0 = rf(channels...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TwitchClient_Join_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Join'
type TwitchClient_Join_Call struct {
	*mock.Call
}

// Join is a helper method to define mock.On call
//   - channels ...string
func (_e *TwitchClient_Expecter) Join(channels ...interface{}) *TwitchClient_Join_Call {
	return &TwitchClient_Join_Call{Call: _e.mock.On("Join",
		append([]interface{}{}, channels...)...)}
}

func (_c *TwitchClient_Join_Call) Run(run func(channels ...string)) *TwitchClient_Join_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *TwitchClient_Join_Call) Return(_a0 error) *TwitchClient_Join_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TwitchClient_Join_Call) RunAndReturn(run func(...string) error) *TwitchClient_Join_Call {
	_c.Call.Return(run)
	return _c
}

// NewTwitchClient creates a new instance of TwitchClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTwitchClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *TwitchClient {
	mock := &TwitchClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
