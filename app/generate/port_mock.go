// Code generated by MockGen. DO NOT EDIT.
// Source: port.go
//
// Generated by this command:
//
//	mockgen -source=port.go -package=generate -destination=port_mock.go
//

// Package generate is a generated GoMock package.
package generate

import (
	context "context"
	reflect "reflect"

	agentdetail "github.com/promptlabth/ms-ai-marketplace/app/agent_detail"
	history "github.com/promptlabth/ms-ai-marketplace/app/history"
	gomock "go.uber.org/mock/gomock"
)

// MockgenerateAdaptor is a mock of generateAdaptor interface.
type MockgenerateAdaptor struct {
	ctrl     *gomock.Controller
	recorder *MockgenerateAdaptorMockRecorder
}

// MockgenerateAdaptorMockRecorder is the mock recorder for MockgenerateAdaptor.
type MockgenerateAdaptorMockRecorder struct {
	mock *MockgenerateAdaptor
}

// NewMockgenerateAdaptor creates a new mock instance.
func NewMockgenerateAdaptor(ctrl *gomock.Controller) *MockgenerateAdaptor {
	mock := &MockgenerateAdaptor{ctrl: ctrl}
	mock.recorder = &MockgenerateAdaptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockgenerateAdaptor) EXPECT() *MockgenerateAdaptorMockRecorder {
	return m.recorder
}

// GenerateAdaptor mocks base method.
func (m *MockgenerateAdaptor) GenerateAdaptor(ctx context.Context, prompt, model string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAdaptor", ctx, prompt, model)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAdaptor indicates an expected call of GenerateAdaptor.
func (mr *MockgenerateAdaptorMockRecorder) GenerateAdaptor(ctx, prompt, model any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAdaptor", reflect.TypeOf((*MockgenerateAdaptor)(nil).GenerateAdaptor), ctx, prompt, model)
}

// MockagentStorage is a mock of agentStorage interface.
type MockagentStorage struct {
	ctrl     *gomock.Controller
	recorder *MockagentStorageMockRecorder
}

// MockagentStorageMockRecorder is the mock recorder for MockagentStorage.
type MockagentStorageMockRecorder struct {
	mock *MockagentStorage
}

// NewMockagentStorage creates a new mock instance.
func NewMockagentStorage(ctrl *gomock.Controller) *MockagentStorage {
	mock := &MockagentStorage{ctrl: ctrl}
	mock.recorder = &MockagentStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockagentStorage) EXPECT() *MockagentStorageMockRecorder {
	return m.recorder
}

// GetAgentByID mocks base method.
func (m *MockagentStorage) GetAgentByID(arg0 context.Context, arg1 int) (*agentdetail.AgentDetailEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentByID", arg0, arg1)
	ret0, _ := ret[0].(*agentdetail.AgentDetailEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentByID indicates an expected call of GetAgentByID.
func (mr *MockagentStorageMockRecorder) GetAgentByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentByID", reflect.TypeOf((*MockagentStorage)(nil).GetAgentByID), arg0, arg1)
}

// MockhistoryStorage is a mock of historyStorage interface.
type MockhistoryStorage struct {
	ctrl     *gomock.Controller
	recorder *MockhistoryStorageMockRecorder
}

// MockhistoryStorageMockRecorder is the mock recorder for MockhistoryStorage.
type MockhistoryStorageMockRecorder struct {
	mock *MockhistoryStorage
}

// NewMockhistoryStorage creates a new mock instance.
func NewMockhistoryStorage(ctrl *gomock.Controller) *MockhistoryStorage {
	mock := &MockhistoryStorage{ctrl: ctrl}
	mock.recorder = &MockhistoryStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockhistoryStorage) EXPECT() *MockhistoryStorageMockRecorder {
	return m.recorder
}

// CreateHistory mocks base method.
func (m *MockhistoryStorage) CreateHistory(ctx context.Context, history history.HistoryEntity) (*int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHistory", ctx, history)
	ret0, _ := ret[0].(*int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHistory indicates an expected call of CreateHistory.
func (mr *MockhistoryStorageMockRecorder) CreateHistory(ctx, history any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHistory", reflect.TypeOf((*MockhistoryStorage)(nil).CreateHistory), ctx, history)
}
