package redis

import (
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
)

// manager redis连接管理器
var manager *Manager

// Manager define a redis manager
type Manager struct {
	mu sync.RWMutex

	clients map[string]*redis.Client
}

// NewClient new a redis client
func (m *Manager) NewClient(name string, c *Config) {
	if name == "" {
		name = "default"
	}
	m.mu.RLock()
	if _, ok := m.clients[name]; ok {
		m.mu.RUnlock()
		return
	}
	m.mu.RUnlock()

	m.mu.Lock()
	defer m.mu.Unlock()

	rdb, err := NewClient(c)
	if err != nil {
		log.Fatalf("init redis client err:%v", err)
	}
	m.clients[name] = rdb
}

// GetClient get a redis client
func (m *Manager) GetClient(name string) *redis.Client {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if c, ok := m.clients[name]; ok {
		return c
	}
	return nil
}

// NewRManager create a redis manager
func NewRManager() *Manager {
	m := &Manager{
		clients: make(map[string]*redis.Client),
	}
	manager = m
	return m
}

// GetManager get a redis manager
func GetManager() *Manager {
	return manager
}

// GetClient get a redis client
func GetClient(name string) *redis.Client {
	return manager.GetClient(name)
}
