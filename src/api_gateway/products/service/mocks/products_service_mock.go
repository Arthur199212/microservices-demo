// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Arthur199212/microservices-demo/src/api_gateway/products/service (interfaces: ProductsService)

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	productsv1 "github.com/Arthur199212/microservices-demo/gen/services/products/v1"
	service "github.com/Arthur199212/microservices-demo/src/api_gateway/products/service"
	gomock "github.com/golang/mock/gomock"
)

// MockProductsService is a mock of ProductsService interface.
type MockProductsService struct {
	ctrl     *gomock.Controller
	recorder *MockProductsServiceMockRecorder
}

// MockProductsServiceMockRecorder is the mock recorder for MockProductsService.
type MockProductsServiceMockRecorder struct {
	mock *MockProductsService
}

// NewMockProductsService creates a new mock instance.
func NewMockProductsService(ctrl *gomock.Controller) *MockProductsService {
	mock := &MockProductsService{ctrl: ctrl}
	mock.recorder = &MockProductsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductsService) EXPECT() *MockProductsServiceMockRecorder {
	return m.recorder
}

// GetProductById mocks base method.
func (m *MockProductsService) GetProductById(arg0 context.Context, arg1 service.GetProductByIdArgs) (*productsv1.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductById", arg0, arg1)
	ret0, _ := ret[0].(*productsv1.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductById indicates an expected call of GetProductById.
func (mr *MockProductsServiceMockRecorder) GetProductById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductById", reflect.TypeOf((*MockProductsService)(nil).GetProductById), arg0, arg1)
}

// ListProducts mocks base method.
func (m *MockProductsService) ListProducts(arg0 context.Context, arg1 service.ListProductsArgs) ([]*productsv1.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts", arg0, arg1)
	ret0, _ := ret[0].([]*productsv1.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockProductsServiceMockRecorder) ListProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockProductsService)(nil).ListProducts), arg0, arg1)
}
