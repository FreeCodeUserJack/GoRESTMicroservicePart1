package services

import (
	"errors"
	"testing"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/domain"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/utils"
)


type SpyUserServiceImpl struct {
	getUserByIdCalls []uint64
}

func (s *SpyUserServiceImpl) GetUserById(id uint64) (*domain.User, error) {
		if id != uint64(1) {
			return nil, errors.New("user not found")
		}
		s.getUserByIdCalls = append(s.getUserByIdCalls, id)
		return nil, nil
}

func TestUserService(t *testing.T) {
	t.Run("get existing user", func(t *testing.T) {
		spyUserServiceImpl := SpyUserServiceImpl{}

		want := uint64(1)

		_, err := spyUserServiceImpl.GetUserById(want)

		utils.AssertNoError(t, err)

		utils.AssertUserId(t, spyUserServiceImpl.getUserByIdCalls[0], want)
	})

	t.Run("user not exists", func(t *testing.T) {
		spyUserServiceImpl := SpyUserServiceImpl{}

		want := uint64(2)

		_, err := spyUserServiceImpl.GetUserById(want)

		utils.AssertError(t, err)
	})
}