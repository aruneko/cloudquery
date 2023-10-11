// Code generated by MockGen. DO NOT EDIT.
// Source: monitors_api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	reflect "reflect"

	datadog "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	datadogV1 "github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	gomock "github.com/golang/mock/gomock"
)

// MockMonitorsAPIClient is a mock of MonitorsAPIClient interface.
type MockMonitorsAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockMonitorsAPIClientMockRecorder
}

// MockMonitorsAPIClientMockRecorder is the mock recorder for MockMonitorsAPIClient.
type MockMonitorsAPIClientMockRecorder struct {
	mock *MockMonitorsAPIClient
}

// NewMockMonitorsAPIClient creates a new mock instance.
func NewMockMonitorsAPIClient(ctrl *gomock.Controller) *MockMonitorsAPIClient {
	mock := &MockMonitorsAPIClient{ctrl: ctrl}
	mock.recorder = &MockMonitorsAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMonitorsAPIClient) EXPECT() *MockMonitorsAPIClientMockRecorder {
	return m.recorder
}

// ListMonitors mocks base method.
func (m *MockMonitorsAPIClient) ListMonitors(arg0 context.Context, arg1 ...datadogV1.ListMonitorsOptionalParameters) ([]datadogV1.Monitor, *http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMonitors", varargs...)
	ret0, _ := ret[0].([]datadogV1.Monitor)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListMonitors indicates an expected call of ListMonitors.
func (mr *MockMonitorsAPIClientMockRecorder) ListMonitors(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMonitors", reflect.TypeOf((*MockMonitorsAPIClient)(nil).ListMonitors), varargs...)
}

// ListMonitorsWithPagination mocks base method.
func (m *MockMonitorsAPIClient) ListMonitorsWithPagination(arg0 context.Context, arg1 ...datadogV1.ListMonitorsOptionalParameters) (<-chan datadog.PaginationResult[datadogV1.Monitor], func()) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMonitorsWithPagination", varargs...)
	ret0, _ := ret[0].(<-chan datadog.PaginationResult[datadogV1.Monitor])
	ret1, _ := ret[1].(func())
	return ret0, ret1
}

// ListMonitorsWithPagination indicates an expected call of ListMonitorsWithPagination.
func (mr *MockMonitorsAPIClientMockRecorder) ListMonitorsWithPagination(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMonitorsWithPagination", reflect.TypeOf((*MockMonitorsAPIClient)(nil).ListMonitorsWithPagination), varargs...)
}
