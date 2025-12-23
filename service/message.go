package service

import (
	"project-portofolio/model"
	"project-portofolio/repository"
)

type MessageService struct {
	Repo repository.MessageRepositoryInterface
}

type MessageServiceInterface interface {
	CreateNewMessage(message *model.Message) error
}

// constructor
func NewMessageService(repo repository.MessageRepositoryInterface) MessageService {
	return MessageService{
		Repo: repo,
	}
}

func (s *MessageService) CreateMessage(message *model.Message) error {
	err := s.Repo.CreateMessage(message)
	if err != nil {
		return err
	}
	return nil
}
