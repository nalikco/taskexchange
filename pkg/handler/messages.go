package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"taskexchange"
	"taskexchange/pkg/service"
)

type getMessagesInput struct {
	ConversationId int `json:"conversation_id" binding:"required"`
}

type messagesResponse struct {
	Data []taskexchange.Message `json:"data"`
}

func (h *Handler) getMessages(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	var input getMessagesInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	conversation, err := h.services.Messages.GetConversationById(input.ConversationId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ifUserInConversation := false
	for _, member := range conversation.Members {
		if ifUserInConversation {
			break
		}

		if member.Id == user.Id {
			ifUserInConversation = true
			break
		}
	}

	messages, err := h.services.Messages.GetMessagesByConversation(conversation)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, messagesResponse{
		Data: messages,
	})
}

type conversationsResponse struct {
	Data []taskexchange.Conversation `json:"data"`
}

func (h *Handler) getConversations(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	conversations, err := h.services.Messages.GetUserConversations(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, conversationsResponse{
		Data: conversations,
	})
}

type sendMessageInput struct {
	RecipientId int    `json:"recipient_id" binding:"required"`
	Text        string `json:"text" binding:"required"`
}

func (h *Handler) sendMessage(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	var input sendMessageInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	recipient, err := h.services.Users.GetById(input.RecipientId, false)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Messages.SendMessageToRecipient(user, recipient, input.Text)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) CountUnViewedMessages(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	count, err := h.services.CountUserUnViewedMessages(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": count,
	})
}

type viewMessagesInput struct {
	ConversationId int `json:"conversation_id" binding:"required"`
}

func (h *Handler) ViewMessages(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	var input viewMessagesInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	conversation, err := h.services.Messages.GetConversationById(input.ConversationId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Messages.ViewConversation(conversation, user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

type oneMessageResponse struct {
	Data service.PollingMessage `json:"data"`
}

func (h *Handler) messagesPolling(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	newMessage, err := h.services.Messages.Polling(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, oneMessageResponse{
		Data: newMessage,
	})
}
