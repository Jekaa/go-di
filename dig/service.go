package dig

import "fmt"

// Интерфейс хранилища
type Store interface {
	GetUser(id int) string
}

// Реализация хранилища в памяти
type MemoryStore struct{}

func (m *MemoryStore) GetUser(id int) string {
	return fmt.Sprintf("User%d", id)
}

// Сервис, зависящий от Store
type UserService struct {
	store Store
}

// Конструктор для UserService
func NewUserService(store Store) *UserService {
	return &UserService{store: store}
}

func (s *UserService) GetUser(id int) string {
	return s.store.GetUser(id)
}
