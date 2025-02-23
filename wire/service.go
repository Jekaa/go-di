package wire

import "fmt"

type UserRepository interface {
	GetUser(id int) string
}

type InMemoryUserRepository struct{}

func (r *InMemoryUserRepository) GetUser(id int) string {
	return fmt.Sprintf("User %d", id)
}

type UserService struct {
	repo UserRepository
}

func (s *UserService) GetUser(id int) string {
	return s.repo.GetUser(id)
}
