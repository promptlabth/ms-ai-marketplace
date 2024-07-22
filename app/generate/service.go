package generate

import "context"

type GenerateService struct {
	generateAdaptor generateAdaptor
	agentStorage    agentStorage
	historyStorage  historyStorage
}

func NewService(
	generateAdaptor generateAdaptor,
	agentStorage agentStorage,
	historyStorage historyStorage,
) *GenerateService {
	return &GenerateService{
		generateAdaptor: generateAdaptor,
		agentStorage:    agentStorage,
		historyStorage:  historyStorage,
	}
}

func (s *GenerateService) Generate(ctx context.Context, generateRequest GenerateRequest) error {
	_, err := s.agentStorage.GetAgentByID(ctx, generateRequest.AgentID)
	if err != nil {
		return err
	}

	return nil
}
