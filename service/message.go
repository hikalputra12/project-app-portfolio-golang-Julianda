package service

import (
	"project-portofolio/model"
	"project-portofolio/repository"

	"go.uber.org/zap"
)

type MessageService struct {
	Repo   repository.MessageRepositoryInterface
	Logger *zap.Logger
}

type MessageServiceInterface interface {
	CreateNewMessage(message *model.Message) error
}

// constructor
func NewMessageService(repo repository.MessageRepositoryInterface, logger *zap.Logger) MessageService {
	return MessageService{
		Repo:   repo,
		Logger: logger,
	}
}

func (s *MessageService) CreateNewMessage(message *model.Message) error {
	err := s.Repo.CreateMessage(message)
	if err != nil {
		s.Logger.Error("error  CreateNewMessage service ", zap.Error(err))
		return err
	}
	return nil
}
