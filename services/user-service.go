package services

import (
	"fmt"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/domain"
)

func PrintUser() {
	fmt.Println("user service print")
	domain.UserDomain()
}