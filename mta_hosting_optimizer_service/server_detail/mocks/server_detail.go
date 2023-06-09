// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/sahil/go/src/github.com/mta-hosting-optimizer/mta_hosting_optimizer_service/server_detail/server_detail.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	server_detail "github.com/mta-hosting-optimizer/mta_hosting_optimizer_service/server_detail"
)

// MockIServersDetail is a mock of IServersDetail interface.
type MockIServersDetail struct {
	ctrl     *gomock.Controller
	recorder *MockIServersDetailMockRecorder
}

// MockIServersDetailMockRecorder is the mock recorder for MockIServersDetail.
type MockIServersDetailMockRecorder struct {
	mock *MockIServersDetail
}

// NewMockIServersDetail creates a new mock instance.
func NewMockIServersDetail(ctrl *gomock.Controller) *MockIServersDetail {
	mock := &MockIServersDetail{ctrl: ctrl}
	mock.recorder = &MockIServersDetailMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIServersDetail) EXPECT() *MockIServersDetailMockRecorder {
	return m.recorder
}

// GetServersDetail mocks base method.
func (m *MockIServersDetail) GetServersDetail(ctx *gin.Context, request *server_detail.GetServersDetailRequest) (*server_detail.GetServersDetailResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServersDetail", ctx, request)
	ret0, _ := ret[0].(*server_detail.GetServersDetailResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServersDetail indicates an expected call of GetServersDetail.
func (mr *MockIServersDetailMockRecorder) GetServersDetail(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServersDetail", reflect.TypeOf((*MockIServersDetail)(nil).GetServersDetail), ctx, request)
}
