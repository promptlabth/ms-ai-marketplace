package generate

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type TestGenerateServiceSuite struct {
	suite.Suite
	ctrl                *gomock.Controller
	mockGenerateAdaptor *MockgenerateAdaptor
	agentStorage        *MockagentStorage
	stylepromptStorage  *MockstylepromptStorage
	frameworkStorage    *MockframeworkStorage
	roleStorage         *MockroleStorage
	historyStorage      *MockhistoryStorage
	storage             *MockStorage

	svc *GenerateService
}

func TestGenerateSuite(t *testing.T) {
	suite.Run(t, new(TestGenerateServiceSuite))
}

func (t *TestGenerateServiceSuite) SetupTest() {
	// Setup before each test
	t.ctrl = gomock.NewController(t.T())
	t.agentStorage = NewMockagentStorage(t.ctrl)
	t.stylepromptStorage = NewMockstylepromptStorage(t.ctrl)
	t.frameworkStorage = NewMockframeworkStorage(t.ctrl)
	t.historyStorage = NewMockhistoryStorage(t.ctrl)
	t.roleStorage = NewMockroleStorage(t.ctrl)
	t.mockGenerateAdaptor = NewMockgenerateAdaptor(t.ctrl)
	t.storage = NewMockStorage(t.ctrl)

	t.svc = NewService(
		t.mockGenerateAdaptor, t.agentStorage, t.stylepromptStorage, t.frameworkStorage, t.roleStorage, t.historyStorage, t.storage,
	)
}

func (s *TestGenerateServiceSuite) TearDownTest() {
	// Tear down after each test
	s.ctrl.Finish()
}

func (t *TestGenerateServiceSuite) TestGenerateGetAgentDetailFailed_ReturnErr() {
	// arrange
	t.agentStorage.EXPECT().GetAgentByID(gomock.Any(), gomock.Any()).Return(
		nil,
		errors.New("Failed to Get Agent"),
	)

	msg := GenerateRequest{}

	// act
	_, err := t.svc.Generate(context.Background(), msg,"th")

	// assert
	t.Error(err)
	t.EqualError(err, "Failed to Get Agent")
}
