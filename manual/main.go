package manual

import "fmt"

func main() {
	// Ручное создание зависимостей
	store := &MemoryStore{}
	service := NewUserService(store)

	fmt.Println(service.GetUser(1)) // User1
}
