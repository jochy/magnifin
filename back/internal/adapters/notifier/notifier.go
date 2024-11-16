package notifier

import (
	"encoding/json"
	"magnifin/internal/adapters"
	"magnifin/internal/app/model"

	"github.com/gorilla/websocket"
)

type Client struct {
	UserID int32
	Conn   *websocket.Conn
	Send   chan []byte
}

type Mapper interface {
	ToPublicFormat(trs *model.Transaction) *adapters.EnrichedTransaction
}

type Notifier struct {
	Mapper  Mapper
	Clients map[int32]*Client
}

func NewNotifier(mapper Mapper) *Notifier {
	return &Notifier{
		Mapper:  mapper,
		Clients: make(map[int32]*Client),
	}
}

func (n *Notifier) Close() error {
	return nil
}

func (n *Notifier) Notify(userID int32, trs *model.Transaction) {
	if n.Clients[userID] == nil {
		return
	}

	t := n.Mapper.ToPublicFormat(trs)
	bytes, err := json.Marshal(t)
	if err != nil {
		return
	}

	n.Clients[userID].Send <- bytes
}

func (n *Notifier) RegisterConnection(userID int32, c *websocket.Conn) *Client {
	if n.Clients[userID] != nil {
		_ = n.Clients[userID].Conn.Close()
	}

	n.Clients[userID] = &Client{
		UserID: userID,
		Conn:   c,
		Send:   make(chan []byte),
	}

	return n.Clients[userID]
}
