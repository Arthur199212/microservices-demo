// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Arthur199212/microservices-demo/gen/services/shipping/v1 (interfaces: ShippingServiceClient)

// Package mock_v1 is a generated GoMock package.
package mock_v1

import (
	context "context"
	reflect "reflect"

	shippingv1 "github.com/Arthur199212/microservices-demo/gen/services/shipping/v1"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockShippingServiceClient is a mock of ShippingServiceClient interface.
type MockShippingServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockShippingServiceClientMockRecorder
}

// MockShippingServiceClientMockRecorder is the mock recorder for MockShippingServiceClient.
type MockShippingServiceClientMockRecorder struct {
	mock *MockShippingServiceClient
}

// NewMockShippingServiceClient creates a new mock instance.
func NewMockShippingServiceClient(ctrl *gomock.Controller) *MockShippingServiceClient {
	mock := &MockShippingServiceClient{ctrl: ctrl}
	mock.recorder = &MockShippingServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShippingServiceClient) EXPECT() *MockShippingServiceClientMockRecorder {
	return m.recorder
}

// GetQuote mocks base method.
func (m *MockShippingServiceClient) GetQuote(arg0 context.Context, arg1 *shippingv1.GetQuoteRequest, arg2 ...grpc.CallOption) (*shippingv1.GetQuoteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQuote", varargs...)
	ret0, _ := ret[0].(*shippingv1.GetQuoteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuote indicates an expected call of GetQuote.
func (mr *MockShippingServiceClientMockRecorder) GetQuote(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuote", reflect.TypeOf((*MockShippingServiceClient)(nil).GetQuote), varargs...)
}

// ShipOrder mocks base method.
func (m *MockShippingServiceClient) ShipOrder(arg0 context.Context, arg1 *shippingv1.ShipOrderRequest, arg2 ...grpc.CallOption) (*shippingv1.ShipOrderResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ShipOrder", varargs...)
	ret0, _ := ret[0].(*shippingv1.ShipOrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShipOrder indicates an expected call of ShipOrder.
func (mr *MockShippingServiceClientMockRecorder) ShipOrder(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShipOrder", reflect.TypeOf((*MockShippingServiceClient)(nil).ShipOrder), varargs...)
}
