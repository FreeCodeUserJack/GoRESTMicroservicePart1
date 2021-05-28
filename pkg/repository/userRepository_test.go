package repository

import (
	"errors"
	"testing"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/domain"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/utils"
)


type SpyUserRepository struct {
	getUserByIdCalls []uint64
}

func (s *SpyUserRepository) GetUserById(userId uint64) (*domain.User, error) {
	if userId != uint64(1) {
		return nil, errors.New("userId not found")
	}
	s.getUserByIdCalls = append(s.getUserByIdCalls, userId)
	return nil, nil
}

func TestUserRepository(t *testing.T) {
	t.Run("returns user for an existing user", func(t *testing.T) {
		spyUserRepository := &SpyUserRepository{}

		want := uint64(1)

		_, err := spyUserRepository.GetUserById(want)

		utils.AssertNoError(t, err)

		utils.AssertUserId(t, spyUserRepository.getUserByIdCalls[0], want)
	})

	t.Run("returns err when user is not found", func(t *testing.T) {
		spyUserRepository := &SpyUserRepository{}

		want := uint64(2)

		_, err := spyUserRepository.GetUserById(want)

		utils.AssertError(t, err)
	})
}