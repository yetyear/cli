// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	client "github.com/cosmos/cosmos-sdk/client"

	mock "github.com/stretchr/testify/mock"

	tx "github.com/cosmos/cosmos-sdk/client/tx"
)

// Signer is an autogenerated mock type for the Signer type
type Signer struct {
	mock.Mock
}

type Signer_Expecter struct {
	mock *mock.Mock
}

func (_m *Signer) EXPECT() *Signer_Expecter {
	return &Signer_Expecter{mock: &_m.Mock}
}

// Sign provides a mock function with given fields: ctx, txf, name, txBuilder, overwriteSig
func (_m *Signer) Sign(ctx context.Context, txf tx.Factory, name string, txBuilder client.TxBuilder, overwriteSig bool) error {
	ret := _m.Called(ctx, txf, name, txBuilder, overwriteSig)

	if len(ret) == 0 {
		panic("no return value specified for Sign")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, tx.Factory, string, client.TxBuilder, bool) error); ok {
		r0 = rf(ctx, txf, name, txBuilder, overwriteSig)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Signer_Sign_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sign'
type Signer_Sign_Call struct {
	*mock.Call
}

// Sign is a helper method to define mock.On call
//   - ctx context.Context
//   - txf tx.Factory
//   - name string
//   - txBuilder client.TxBuilder
//   - overwriteSig bool
func (_e *Signer_Expecter) Sign(ctx interface{}, txf interface{}, name interface{}, txBuilder interface{}, overwriteSig interface{}) *Signer_Sign_Call {
	return &Signer_Sign_Call{Call: _e.mock.On("Sign", ctx, txf, name, txBuilder, overwriteSig)}
}

func (_c *Signer_Sign_Call) Run(run func(ctx context.Context, txf tx.Factory, name string, txBuilder client.TxBuilder, overwriteSig bool)) *Signer_Sign_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(tx.Factory), args[2].(string), args[3].(client.TxBuilder), args[4].(bool))
	})
	return _c
}

func (_c *Signer_Sign_Call) Return(_a0 error) *Signer_Sign_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Signer_Sign_Call) RunAndReturn(run func(context.Context, tx.Factory, string, client.TxBuilder, bool) error) *Signer_Sign_Call {
	_c.Call.Return(run)
	return _c
}

// NewSigner creates a new instance of Signer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSigner(t interface {
	mock.TestingT
	Cleanup(func())
}) *Signer {
	mock := &Signer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
