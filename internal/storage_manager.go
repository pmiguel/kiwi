package internal

import (
	"fmt"
	"strings"
	"sync"
)

type StorageManager struct {
	data   sync.Map
	server *Server
}

func NewStorageManager(server *Server) *StorageManager {
	storageManager := &StorageManager{
		server: server,
	}

	server.StorageManager = storageManager
	return storageManager
}

func (sm *StorageManager) Get(key string) any {
	value, ok := sm.data.Load(key)

	if ok {
		return value
	}

	return nil
}

func (sm *StorageManager) Set(key string, value string) {
	sm.data.Store(key, value)
}

func (sm *StorageManager) Delete(key string) {
	sm.data.Delete(key)
}

func (sm *StorageManager) Keys() string {
	var keys []string
	sm.data.Range(func(key, value any) bool {
		keys = append(keys, fmt.Sprintf("%v", key))
		return true
	})

	return strings.Join(keys[:], ",")
}
