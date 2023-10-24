package websocket_conn_pool

import (
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type WSUserConnPool struct {
	Pool sync.Map
}

type Client struct {
	conn       GWebSocket
	PlatformID int    `json:"platformID"`
	UserID     string `json:"userID"`
}

type GWebSocket struct {
	protocolType     int
	conn             *websocket.Conn
	handshakeTimeout time.Duration
}

// Set StoreUserConn store user conn
func (p *WSUserConnPool) Set(userID string, v *Client) {
	p.Pool.Store(userID, v)
}

// Get specific user conn pool
func (p *WSUserConnPool) Get(userID string) (v *Client) {
	client, ok := p.Pool.Load(userID)
	if !ok {
		return nil
	}
	return client.(*Client)
}

// Delete user specific conn
func (p *WSUserConnPool) Delete(userID string) {
	p.Pool.Delete(userID)
}
