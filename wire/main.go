package wire

import "fmt"

func main() {
	service := InitializeUserService()

	fmt.Println(service.GetUser(1)) // Вывод: User 1
	fmt.Println(service.GetUser(2)) // Вывод: User 2
}
