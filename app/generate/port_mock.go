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
	"github.com/promptlabth/ms-ai-marketplace/app/framework"
	history "github.com/promptlabth/ms-ai-marketplace/app/history"
	"github.com/promptlabth/ms-ai-marketplace/app/role"
	styleprompt "github.com/promptlabth/ms-ai-marketplace/app/style_prompt"
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
func (mr *MockagentStorageMockRecorder) GetAgentByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentByID", reflect.TypeOf((*MockagentStorage)(nil).GetAgentByID), arg0, arg1)
}

// UpdateAgentDetail mocks base method.
func (m *MockagentStorage) UpdateAgentDetail(arg0 context.Context, arg1 agentdetail.AgentDetailEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAgentDetail", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAgentDetail indicates an expected call of UpdateAgentDetail.
func (mr *MockagentStorageMockRecorder) UpdateAgentDetail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAgentDetail", reflect.TypeOf((*MockagentStorage)(nil).UpdateAgentDetail), arg0, arg1)
}

type MockstylepromptStorage struct {
	ctrl     *gomock.Controller
	recorder *MockstylepromptStorageMockRecorder
}

// MockstylepromptStorageMockRecorder is the mock recorder for MockstylepromptStorage.
type MockstylepromptStorageMockRecorder struct {
	mock *MockstylepromptStorage
}

// NewMockstylepromptStorage creates a new mock instance.
func NewMockstylepromptStorage(ctrl *gomock.Controller) *MockstylepromptStorage {
	mock := &MockstylepromptStorage{ctrl: ctrl}
	mock.recorder = &MockstylepromptStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockstylepromptStorage) EXPECT() *MockstylepromptStorageMockRecorder {
	return m.recorder
}

// GetStylePromptByID mocks base method.
func (m *MockstylepromptStorage) GetStylePromptByID(arg0 context.Context, arg1 int) (*styleprompt.StylePromptEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStylePromptByID", arg0, arg1)
	ret0, _ := ret[0].(*styleprompt.StylePromptEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStylePromptByID indicates an expected call of GetStylePromptByID.
func (mr *MockstylepromptStorageMockRecorder) GetStylePromptByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStylePromptByID", reflect.TypeOf((*MockstylepromptStorage)(nil).GetStylePromptByID), arg0, arg1)
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

// MockStorage is a mock of the storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Generate mocks the base method
func (m *MockStorage) Generate(ctx context.Context, prompt string, model string) (string, int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", ctx, prompt, model)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(int)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Generate indicates an expected call of Generate
func (mr *MockStorageMockRecorder) Generate(ctx, prompt, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockStorage)(nil).Generate), ctx, prompt, model)
}

type MockframeworkStorage struct {
	ctrl     *gomock.Controller
	recorder *MockframeworkStorageMockRecorder
}

// MockframeworkStorageMockRecorder is the mock recorder for MockframeworkStorage
type MockframeworkStorageMockRecorder struct {
	mock *MockframeworkStorage
}

// NewMockframeworkStorage creates a new mock instance
func NewMockframeworkStorage(ctrl *gomock.Controller) *MockframeworkStorage {
	mock := &MockframeworkStorage{ctrl: ctrl}
	mock.recorder = &MockframeworkStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockframeworkStorage) EXPECT() *MockframeworkStorageMockRecorder {
	return m.recorder
}

// GetframeworkyID mocks the base method
func (m *MockframeworkStorage) GetFrameworkByID(ctx context.Context, id int) (*framework.FrameworkEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFrameworkByID", ctx, id)
	ret0, _ := ret[0].(*framework.FrameworkEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetframeworkyID indicates an expected call of GetframeworkyID
func (mr *MockframeworkStorageMockRecorder) GetFrameworkByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFrameworkByID", reflect.TypeOf((*MockframeworkStorage)(nil).GetFrameworkByID), ctx, id)
}

// MockroleStorage is a mock of roleStorage
type MockroleStorage struct {
	ctrl     *gomock.Controller
	recorder *MockroleStorageMockRecorder
}

// MockroleStorageMockRecorder is the mock recorder for MockroleStorage
type MockroleStorageMockRecorder struct {
	mock *MockroleStorage
}

// NewMockroleStorage creates a new mock instance
func NewMockroleStorage(ctrl *gomock.Controller) *MockroleStorage {
	mock := &MockroleStorage{ctrl: ctrl}
	mock.recorder = &MockroleStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockroleStorage) EXPECT() *MockroleStorageMockRecorder {
	return m.recorder
}

// GetRoleByID mocks the base method
func (m *MockroleStorage) GetRoleByID(ctx context.Context, id int) (*role.RoleEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleByID", ctx, id)
	ret0, _ := ret[0].(*role.RoleEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleByID indicates an expected call of GetRoleByID
func (mr *MockroleStorageMockRecorder) GetRoleByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleByID", reflect.TypeOf((*MockroleStorage)(nil).GetRoleByID), ctx, id)
}
