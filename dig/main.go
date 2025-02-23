package dig

import (
	"fmt"
	"go.uber.org/dig"
)

func main() {
	// Создаем контейнер Dig
	container := dig.New()

	// Регистрируем провайдеры
	container.Provide(func() Store {
		return &MemoryStore{}
	})
	container.Provide(NewUserService)

	// Извлекаем сервис и используем его
	err := container.Invoke(func(service *UserService) {
		fmt.Println(service.GetUser(1)) // User1
	})

	if err != nil {
		panic(err)
	}
}
