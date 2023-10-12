package services

import (
	"github.com/google/uuid"
	"github.com/muathendirangu/hex/internal/core/domain"
	"github.com/muathendirangu/hex/internal/core/ports"
)

type MessageService struct {
	repo ports.MessageRepository
}

func NewMessageService(repo ports.MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) SaveMessage(message domain.Message) error {
	message.ID = uuid.New().String()
	return s.repo.SaveMessage(message)
}

func (s *MessageService) ReadMessage(id string) (*domain.Message, error) {
	return s.repo.ReadMessage(id)
}

func (s *MessageService) ReadMessages() ([]*domain.Message, error) {
	return s.repo.ReadMessages()

}
