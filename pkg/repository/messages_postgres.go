package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
	"taskexchange"
	"time"
)

type MessagesPostgres struct {
	db *sqlx.DB
}

type conversationMembers struct {
	ID             int        `db:"id"`
	ConversationId int        `db:"conversation_id"`
	UserId         int        `db:"user_id"`
	CreatedAt      time.Time  `db:"created_at"`
	DeletedAt      *time.Time `db:"deleted_at"`
}

type messagesDb struct {
	ID             int        `db:"id"`
	ConversationId int        `db:"conversation_id"`
	SenderId       int        `db:"user_id"`
	Text           string     `db:"text"`
	ViewedAt       *time.Time `db:"viewed_at"`
	CreatedAt      time.Time  `db:"created_at"`
	DeletedAt      *time.Time `db:"deleted_at"`
}

func NewMessagesPostgres(db *sqlx.DB) *MessagesPostgres {
	return &MessagesPostgres{db: db}
}

func (r *MessagesPostgres) GetUserConversations(user taskexchange.User) ([]taskexchange.Conversation, error) {
	var conversations []taskexchange.Conversation
	var userConversationMembers []conversationMembers
	var conversationMembers []conversationMembers

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", messagesConversationMembersTable)
	err := r.db.Select(&userConversationMembers, query, user.Id)
	if err != nil {
		return conversations, err
	}

	var conversationIds []string
	for _, member := range userConversationMembers {
		conversationIds = append(conversationIds, strconv.Itoa(member.ConversationId))
	}

	if len(conversationIds) < 1 {
		return conversations, nil
	}

	conversationIdsString := strings.Join(conversationIds, ", ")

	query = fmt.Sprintf("SELECT * FROM %s WHERE id IN (%s)", messagesConversationsTable, conversationIdsString)
	err = r.db.Select(&conversations, query)
	if err != nil {
		return conversations, err
	}

	query = fmt.Sprintf("SELECT * FROM %s WHERE conversation_id IN (%s)", messagesConversationMembersTable, conversationIdsString)
	err = r.db.Select(&conversationMembers, query)
	if err != nil {
		return conversations, err
	}

	for i, conversation := range conversations {
		conversations[i].Messages, err = r.GetMessagesByConversation(conversation)
		if err != nil {
			return conversations, err
		}
		for _, conversationMember := range conversationMembers {

			if conversation.ID == conversationMember.ConversationId {
				var member taskexchange.User
				query = fmt.Sprintf("SELECT id,username,type,last_online,created_at FROM %s WHERE id=$1", usersTable)
				err = r.db.Get(&member, query, conversationMember.UserId)
				if err != nil {
					return conversations, err
				}

				conversations[i].Members = append(conversations[i].Members, member)
			}
		}
	}

	return conversations, nil
}

func (r *MessagesPostgres) CreateConversation(members []taskexchange.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (created_at) VALUES (now()) RETURNING id", messagesConversationsTable)
	row := r.db.QueryRow(query)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	for _, member := range members {
		query = fmt.Sprintf("INSERT INTO %s (conversation_id, user_id) VALUES ($1, $2) RETURNING id", messagesConversationMembersTable)
		_, err := r.db.Exec(query, id, member.Id)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (r *MessagesPostgres) GetConversationById(id int) (taskexchange.Conversation, error) {
	var conversation taskexchange.Conversation
	var conversationMembers []conversationMembers

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", messagesConversationsTable)
	err := r.db.Get(&conversation, query, id)

	query = fmt.Sprintf("SELECT * FROM %s WHERE conversation_id=$1", messagesConversationMembersTable)
	err = r.db.Select(&conversationMembers, query, conversation.ID)
	if err != nil {
		return conversation, err
	}

	for _, conversationMember := range conversationMembers {
		var member taskexchange.User

		query = fmt.Sprintf("SELECT id,username,type,last_online,created_at FROM %s WHERE id=$1", usersTable)
		err = r.db.Get(&member, query, conversationMember.UserId)
		if err != nil {
			return conversation, err
		}

		conversation.Members = append(conversation.Members, member)
	}

	return conversation, err
}

func (r *MessagesPostgres) GetMessageById(id int) (taskexchange.Message, error) {
	var message taskexchange.Message
	var messageDb messagesDb
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", messagesTable)
	err := r.db.Get(&messageDb, query, id)
	if err != nil {
		return message, err
	}

	query = fmt.Sprintf("SELECT id,username,type,last_online,created_at FROM %s WHERE id=$1", usersTable)
	err = r.db.Get(&message.Sender, query, messageDb.SenderId)
	if err != nil {
		return message, err
	}

	message.ID = messageDb.ID
	message.Conversation.ID = messageDb.ConversationId
	message.Text = messageDb.Text
	message.ViewedAt = messageDb.ViewedAt
	message.CreatedAt = messageDb.CreatedAt
	message.DeletedAt = messageDb.DeletedAt

	return message, nil
}

func (r *MessagesPostgres) SendMessageToRecipient(sender taskexchange.User, recipient taskexchange.User, text string) (int, error) {
	var id int

	userConversations, err := r.GetUserConversations(sender)
	if err != nil {
		return 0, err
	}

	var conversationWithRecipient taskexchange.Conversation

	for _, conversation := range userConversations {
		if conversationWithRecipient.ID != 0 {
			break
		}

		for _, member := range conversation.Members {
			if member.Id == recipient.Id {
				conversationWithRecipient = conversation
				break
			}
		}
	}

	if conversationWithRecipient.ID == 0 {
		newConversationId, err := r.CreateConversation([]taskexchange.User{
			sender,
			recipient,
		})
		if err != nil {
			return 0, err
		}

		conversationWithRecipient, err = r.GetConversationById(newConversationId)
	}

	query := fmt.Sprintf("INSERT INTO %s (conversation_id, user_id, text) VALUES ($1, $2, $3) RETURNING id", messagesTable)
	row := r.db.QueryRow(query, conversationWithRecipient.ID, sender.Id, text)
	if err = row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *MessagesPostgres) GetMessagesByConversation(conversation taskexchange.Conversation) ([]taskexchange.Message, error) {
	var messages []taskexchange.Message
	var messagesDb []messagesDb

	query := fmt.Sprintf("SELECT * FROM %s WHERE conversation_id=$1 ORDER BY created_at DESC LIMIT 500", messagesTable)
	err := r.db.Select(&messagesDb, query, conversation.ID)
	if err != nil {
		return messages, err
	}

	for _, messageDb := range messagesDb {
		var sender taskexchange.User

		query = fmt.Sprintf("SELECT id,username,type,last_online,created_at FROM %s WHERE id=$1", usersTable)
		err = r.db.Get(&sender, query, messageDb.SenderId)
		if err != nil {
			return messages, err
		}

		messages = append(messages, taskexchange.Message{
			ID:           messageDb.ID,
			Conversation: conversation,
			Sender:       sender,
			Text:         messageDb.Text,
			ViewedAt:     messageDb.ViewedAt,
			CreatedAt:    messageDb.CreatedAt,
			DeletedAt:    messageDb.DeletedAt,
		})
	}

	return messages, nil
}

func (r *MessagesPostgres) CountUserUnViewedMessages(user taskexchange.User) (int, error) {
	var count int
	var userConversationMembers []conversationMembers

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", messagesConversationMembersTable)
	err := r.db.Select(&userConversationMembers, query, user.Id)
	if err != nil {
		return count, err
	}

	var conversationIds []string
	for _, member := range userConversationMembers {
		conversationIds = append(conversationIds, strconv.Itoa(member.ConversationId))
	}

	if len(conversationIds) < 1 {
		return count, nil
	}

	conversationIdsString := strings.Join(conversationIds, ", ")

	query = fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE user_id != %d AND conversation_id IN (%s) AND viewed_at is null", messagesTable, user.Id, conversationIdsString)
	err = r.db.QueryRow(query).Scan(&count)
	if err != nil {
		return count, err
	}

	return count, nil
}

func (r *MessagesPostgres) ViewConversation(conversation taskexchange.Conversation, user taskexchange.User) error {
	query := fmt.Sprintf("UPDATE %s SET viewed_at=now() WHERE conversation_id=$1 AND user_id != $2", messagesTable)
	_, err := r.db.Exec(query, conversation.ID, user.Id)
	if err != nil {
		return err
	}

	return nil
}
