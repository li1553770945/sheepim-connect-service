package service

import (
	"fmt"
	"github.com/hertz-contrib/websocket"
	"sync"
	"time"
)

// ClientConnMap 线程安全的客户端映射
type ClientConnMap struct {
	mu              sync.RWMutex
	store           map[string]*websocket.Conn
	lastPingTimeMap map[string]time.Time // 新增：记录客户端的最后心跳时间

}

var clientConnMap *ClientConnMap

var once sync.Once

func NewClientConnMap() *ClientConnMap {
	once.Do(func() {
		clientConnMap = &ClientConnMap{
			store:           make(map[string]*websocket.Conn),
			lastPingTimeMap: make(map[string]time.Time), // 初始化心跳时间记录
		}
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

// UpdateLastPing 更新客户端最后心跳时间
func (tsm *ClientConnMap) UpdateLastPing(clientID string) {
	tsm.mu.Lock()
	defer tsm.mu.Unlock()
	tsm.lastPingTimeMap[clientID] = time.Now()
}

// GetLastPing 获取客户端最后心跳时间
func (tsm *ClientConnMap) GetLastPing(clientID string) (time.Time, bool) {
	tsm.mu.RLock()
	defer tsm.mu.RUnlock()
	lastPing, exists := tsm.lastPingTimeMap[clientID]
	return lastPing, exists
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
