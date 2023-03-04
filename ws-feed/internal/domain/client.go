package domain

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Client struct {
	ID     uuid.UUID
	WS     *websocket.Conn
	SendCh chan []byte
	logger *zap.Logger
}

func NewClient(userID uuid.UUID, conn *websocket.Conn) *Client {
	return &Client{
		ID:     userID,
		WS:     conn,
		SendCh: make(chan []byte, maxMessageSize),
	}
}

func (c *Client) ReadFromWS(unregister chan<- *Client, broadcast chan<- []byte) {
	defer func() {
		unregister <- c

		err := c.WS.Close()
		if err != nil {
			c.logger.Error(fmt.Sprintf("can't close read ws: %v", err))
		}
	}()

	c.WS.SetReadLimit(maxMessageSize)

	err := c.WS.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		c.logger.Error(fmt.Sprintf("can't set read deadline: %v", err))
	}

	c.WS.SetPongHandler(func(string) error {
		return c.WS.SetReadDeadline(time.Now().Add(pongWait))
	})

	for {
		_, message, err := c.WS.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}

			break
		}

		message = bytes.TrimSpace(bytes.ReplaceAll(message, newline, space))
		broadcast <- message
	}
}

//nolint:gocognit // ok
func (c *Client) WriteToWS() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()

		err := c.WS.Close()
		if err != nil {
			c.logger.Error(fmt.Sprintf("can't close write ws: %v", err))
		}
	}()

	for {
		select {
		case message, ok := <-c.SendCh:
			err := c.WS.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				c.logger.Error(fmt.Sprintf("can't set write deadline: %v", err))
			}

			if !ok {
				err = c.WS.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					c.logger.Error(fmt.Sprintf("can't send close msg: %v", err))
				}

				return
			}

			w, err := c.WS.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			_, err = w.Write(message)
			if err != nil {
				c.logger.Error(fmt.Sprintf("can't write ws msg: %v", err))
				return
			}

			n := len(c.SendCh)
			for i := 0; i < n; i++ {
				_, err = w.Write(newline)
				if err != nil {
					c.logger.Error(fmt.Sprintf("can't write ws msg newline: %v", err))
					return
				}

				_, err = w.Write(<-c.SendCh)
				if err != nil {
					c.logger.Error(fmt.Sprintf("can't write ws msg next: %v", err))
					return
				}
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			err := c.WS.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				c.logger.Error(fmt.Sprintf("can't set waiting write deadline: %v", err))
			}

			if err := c.WS.WriteMessage(websocket.PingMessage, nil); err != nil {
				c.logger.Error(fmt.Sprintf("can't send ping: %v", err))
				return
			}
		}
	}
}
