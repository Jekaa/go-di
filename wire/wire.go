//go:build wireinject
// +build wireinject

package wire

import "github.com/google/wire"

func NewUserRepository() UserRepository {
	return &InMemoryUserRepository{}
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

var UserSet = wire.NewSet(
	NewUserRepository,
	NewUserService,
)

func InitializeUserService() *UserService {
	wire.Build(UserSet)
	return nil
}
