package internal

type StorageManager struct {
	data   map[string]string
	server *Server
}

func NewStorageManager(server *Server) *StorageManager {
	storageManager := &StorageManager{data: make(map[string]string), server: server}

	server.StorageManager = storageManager
	return storageManager
}

func (s *StorageManager) Get(key string) string {
	return s.data[key]
}

func (s *StorageManager) Set(key string, value string) {
	s.data[key] = value
}

func (s *StorageManager) Delete(key string) {
	delete(s.data, key)
}
