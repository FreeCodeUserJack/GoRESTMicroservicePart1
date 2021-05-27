package services

import (
	"fmt"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/domain"
)


func UserService() {
	fmt.Println("user service")
	domain.UserDomain()
}