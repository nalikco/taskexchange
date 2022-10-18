package service

import (
	"taskexchange"
	"taskexchange/pkg/repository"
	"time"
)

type MessagesService struct {
	repo        repository.Messages
	newMessages []PollingMessage
}

type PollingMessage struct {
	ID             int    `json:"id"`
	SenderId       int    `json:"sender_id"`
	RecipientId    int    `json:"recipient_id"`
	ConversationId int    `json:"conversation_id"`
	Text           string `json:"text"`
}

func NewMessagesService(repo repository.Messages) *MessagesService {
	return &MessagesService{
		repo:        repo,
		newMessages: []PollingMessage{},
	}
}

func (s *MessagesService) GetUserConversations(user taskexchange.User) ([]taskexchange.Conversation, error) {
	return s.repo.GetUserConversations(user)
}

func (s *MessagesService) CreateConversation(members []taskexchange.User) (int, error) {
	return s.repo.CreateConversation(members)
}

func (s *MessagesService) GetConversationById(id int) (taskexchange.Conversation, error) {
	return s.repo.GetConversationById(id)
}

func (s *MessagesService) SendMessageToRecipient(sender taskexchange.User, recipient taskexchange.User, text string) (int, error) {
	id, err := s.repo.SendMessageToRecipient(sender, recipient, text)
	if err != nil {
		return 0, err
	}

	message, err := s.repo.GetMessageById(id)
	if err != nil {
		return 0, err
	}

	s.newMessages = append(s.newMessages, PollingMessage{
		ID:             message.ID,
		SenderId:       sender.Id,
		RecipientId:    recipient.Id,
		ConversationId: message.Conversation.ID,
		Text:           text,
	})

	return id, nil
}

func (s *MessagesService) GetMessagesByConversation(conversation taskexchange.Conversation) ([]taskexchange.Message, error) {
	return s.repo.GetMessagesByConversation(conversation)
}

func (s *MessagesService) CountUserUnViewedMessages(user taskexchange.User) (int, error) {
	return s.repo.CountUserUnViewedMessages(user)
}

func (s *MessagesService) ViewConversation(conversation taskexchange.Conversation, user taskexchange.User) error {
	return s.repo.ViewConversation(conversation, user)
}

func (s *MessagesService) Polling(user taskexchange.User) (PollingMessage, error) {
	for i := 0; i < 90; i++ {
		for i, message := range s.newMessages {
			if message.RecipientId == user.Id {
				s.newMessages = append(s.newMessages[:i], s.newMessages[i+1:]...)
				return message, nil
			}
		}
		time.Sleep(time.Millisecond * 100)
	}

	return PollingMessage{}, nil
}
