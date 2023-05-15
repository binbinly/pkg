package orm

import (
	"sync"

	"gorm.io/gorm"
)

// manager db连接管理器
var manager *Manager

// Manager define a db manager
type Manager struct {
	mu sync.RWMutex

	clients map[string]*gorm.DB
}

// InitClient init a redis client
func (m *Manager) InitClient(name string, c *Config) {
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

	m.clients[name] = NewMySQL(c)
}

// GetClient get a redis client
func (m *Manager) GetClient(name string) *gorm.DB {
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
		clients: make(map[string]*gorm.DB),
	}
	manager = m
	return m
}

// GetManager get a redis manager
func GetManager() *Manager {
	return manager
}

// GetClient get a redis client
func GetClient(name string) *gorm.DB {
	return manager.GetClient(name)
}
