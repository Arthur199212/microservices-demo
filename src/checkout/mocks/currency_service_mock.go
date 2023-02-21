// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Arthur199212/microservices-demo/gen/services/currency/v1 (interfaces: CurrencyServiceClient)

// Package mock_v1 is a generated GoMock package.
package mock_v1

import (
	context "context"
	reflect "reflect"

	currencyv1 "github.com/Arthur199212/microservices-demo/gen/services/currency/v1"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockCurrencyServiceClient is a mock of CurrencyServiceClient interface.
type MockCurrencyServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockCurrencyServiceClientMockRecorder
}

// MockCurrencyServiceClientMockRecorder is the mock recorder for MockCurrencyServiceClient.
type MockCurrencyServiceClientMockRecorder struct {
	mock *MockCurrencyServiceClient
}

// NewMockCurrencyServiceClient creates a new mock instance.
func NewMockCurrencyServiceClient(ctrl *gomock.Controller) *MockCurrencyServiceClient {
	mock := &MockCurrencyServiceClient{ctrl: ctrl}
	mock.recorder = &MockCurrencyServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCurrencyServiceClient) EXPECT() *MockCurrencyServiceClientMockRecorder {
	return m.recorder
}

// Convert mocks base method.
func (m *MockCurrencyServiceClient) Convert(arg0 context.Context, arg1 *currencyv1.ConvertRequest, arg2 ...grpc.CallOption) (*currencyv1.ConvertResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Convert", varargs...)
	ret0, _ := ret[0].(*currencyv1.ConvertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Convert indicates an expected call of Convert.
func (mr *MockCurrencyServiceClientMockRecorder) Convert(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Convert", reflect.TypeOf((*MockCurrencyServiceClient)(nil).Convert), varargs...)
}

// GetSupportedCurrencies mocks base method.
func (m *MockCurrencyServiceClient) GetSupportedCurrencies(arg0 context.Context, arg1 *currencyv1.GetSupportedCurrenciesRequest, arg2 ...grpc.CallOption) (*currencyv1.GetSupportedCurrenciesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSupportedCurrencies", varargs...)
	ret0, _ := ret[0].(*currencyv1.GetSupportedCurrenciesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSupportedCurrencies indicates an expected call of GetSupportedCurrencies.
func (mr *MockCurrencyServiceClientMockRecorder) GetSupportedCurrencies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupportedCurrencies", reflect.TypeOf((*MockCurrencyServiceClient)(nil).GetSupportedCurrencies), varargs...)
}