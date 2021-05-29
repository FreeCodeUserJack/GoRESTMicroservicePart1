package services

import (
	"testing"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/domain"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/utils"
)

// type SpyUserServiceImpl struct {
// 	getUserByIdCalls []uint64
// }

// func (s *SpyUserServiceImpl) GetUserById(id uint64) (*domain.User, error) {
// 		if id != uint64(1) {
// 			return nil, errors.New("user not found")
// 		}
// 		s.getUserByIdCalls = append(s.getUserByIdCalls, id)
// 		return nil, nil
// }

func TestUserService(t *testing.T) {
	userServiceImpl := NewUserServiceImpl(utils.TestUserInMemoryRepository{})

	t.Run("get existing user", func(t *testing.T) {

		want := uint64(1)
		wantUser := &domain.User{Id: 1, FirstName: "John", LastName: "Doe", Email: "johndoe@example.com"}

		user, err := userServiceImpl.GetUserById(want)

		utils.AssertNoApplicationError(t, err)

		utils.AssertUserId(t, user.Id, want)

		utils.AssertEqualInstance(t, user, wantUser)
	})

	t.Run("user not exists", func(t *testing.T) {

		want := uint64(2)

		_, err := userServiceImpl.GetUserById(want)

		utils.AssertError(t, err)
	})
}