package service

import (
	"fmt"
	"github.com/hertz-contrib/websocket"
	"sync"
)

// ClientConnMap 线程安全的客户端映射
type ClientConnMap struct {
	mu    sync.RWMutex
	store map[string]*websocket.Conn
}

var clientConnMap *ClientConnMap

var once sync.Once

func NewClientConnMap() *ClientConnMap {
	once.Do(func() {
		clientConnMap = &ClientConnMap{store: make(map[string]*websocket.Conn)}
	})
	return clientConnMap
}

// Add 添加客户端连接
func (tsm *ClientConnMap) Add(clientID string, conn *websocket.Conn) {
	tsm.mu.Lock()
	defer tsm.mu.Unlock()
	tsm.store[clientID] = conn
}

// Remove 删除客户端连接
func (tsm *ClientConnMap) Remove(clientID string) {
	tsm.mu.Lock()
	defer tsm.mu.Unlock()
	delete(tsm.store, clientID)
}

// Get 获取客户端连接
func (tsm *ClientConnMap) Get(clientID string) (*websocket.Conn, bool) {
	tsm.mu.RLock()
	defer tsm.mu.RUnlock()
	conn, exists := tsm.store[clientID]
	return conn, exists
}

// SendMessage 给特定客户端发送消息
func (tsm *ClientConnMap) SendMessage(clientID string, message string) error {
	tsm.mu.RLock()
	defer tsm.mu.RUnlock()

	conn, exists := tsm.store[clientID]
	if !exists {
		return fmt.Errorf("client %s not found", clientID)
	}

	// 发送消息到客户端
	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send message to client %s: %v", clientID, err)
	}

	return nil
}