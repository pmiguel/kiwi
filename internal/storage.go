package internal

type Storage struct {
	data map[string]any
}

func NewStorage() *Storage {
	storage := Storage{data: make(map[string]any)}
	return &storage
}

func (s *Storage) Get(key string) interface{} {
	return s.data[key]
}

func (s *Storage) Set(key string, value any) {
	s.data[key] = value
}

func (s *Storage) Delete(key string) {
	delete(s.data, key)
}
