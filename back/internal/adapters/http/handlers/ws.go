package handlers

import (
	"magnifin/internal/adapters/http/middlewares"
	"magnifin/internal/adapters/notifier"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WSHandler struct {
	Notifier *notifier.Notifier
}

func NewWSHandler(notifier *notifier.Notifier) *WSHandler {
	return &WSHandler{
		Notifier: notifier,
	}
}

func (h *WSHandler) Listen(gctx *gin.Context) {
	user := middlewares.GetUser(gctx.Request.Context())
	if user == nil {
		gctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	c, err := upgrader.Upgrade(gctx.Writer, gctx.Request, nil)
	if err != nil {
		gctx.JSON(400, gin.H{"error": "unable to upgrade connection in websocket"})
		return
	}

	client := h.Notifier.RegisterConnection(user.ID, c)
	ticker := time.NewTicker(3 * time.Second)
	defer func() {
		ticker.Stop()
		_ = client.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-client.Send:
			_ = client.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				// The hub closed the channel.
				_ = client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, _ = w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(client.Send)
			for i := 0; i < n; i++ {
				_, _ = w.Write(<-client.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			_ = client.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
