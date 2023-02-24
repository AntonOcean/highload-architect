package formatter

import (
	"fmt"
	"strings"
	"time"

	"chat/internal/domain"
)

type Text struct {
	Text string `json:"text" example:"Lorem ipsum dolor sit amet" binding:"required"`
}

func (t *Text) ToDomain() (string, error) {
	text := strings.TrimSpace(t.Text)
	if text == "" {
		return "", fmt.Errorf("%w %s", ErrInvalidData, "Поле <Текст> не может быть пустым")
	}

	// валидация именно на размер
	if len(text) > 200 {
		return "", fmt.Errorf("%w %s", ErrInvalidData, "Поле <Текст> не может быть больше 200 символов")
	}

	return text, nil
}

type GetMessage struct {
	From DomainIDType `json:"author_id"`
	To   DomainIDType `json:"to"`
	Text
	Created time.Time `json:"created"`
}

func CreateChatResp(msg *domain.Message) *GetMessage {
	if msg == nil {
		return nil
	}

	return &GetMessage{
		From:    DomainIDType(msg.SenderID.String()),
		To:      DomainIDType(msg.ReceiverID.String()),
		Text:    Text{msg.Text},
		Created: msg.Created,
	}
}

func CreateChatListResp(msgs []*domain.Message) []*GetMessage {
	response := make([]*GetMessage, len(msgs))

	for i := range msgs {
		response[i] = CreateChatResp(msgs[i])
	}

	return response
}
