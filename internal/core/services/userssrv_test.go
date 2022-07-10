package services_test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/richiMarchi/scratchpay-challenge/internal/core/domain"
	"github.com/richiMarchi/scratchpay-challenge/internal/core/services"
	"github.com/richiMarchi/scratchpay-challenge/internal/repositories"
)

func TestUserService(t *testing.T) {
	usersRepository := repositories.NewMemKvs()
	usersService := services.New(usersRepository)

	t.Run("List users, empty list, OK", func(t *testing.T) {
		expectedResponse := []domain.User{}
		response, err := usersService.List()

		assert.Equal(t, expectedResponse, response)
		assert.Equal(t, nil, err)
	})

	t.Run("Add user with id = 0, OK", func(t *testing.T) {
		user := domain.NewUser(0, "pippo")
		err := usersService.Create(user.Id, user.Name)
		list, _ := usersService.List()

		assert.Equal(t, nil, err)
		assert.Equal(t, len(list), 1)
		assert.Equal(t, true, contains(list, user.Id))
	})

	t.Run("Add user with id = 0 again, KO", func(t *testing.T) {
		user := domain.NewUser(0, "pluto")
		err := usersService.Create(user.Id, user.Name)
		list, _ := usersService.List()

		assert.NotEqual(t, nil, err)
		assert.Equal(t, len(list), 1)
		assert.Equal(t, true, contains(list, user.Id))
	})

	t.Run("Add user with id = 1, OK", func(t *testing.T) {
		user := domain.NewUser(1, "paperino")
		err := usersService.Create(user.Id, user.Name)
		list, _ := usersService.List()

		assert.Equal(t, nil, err)
		assert.Equal(t, len(list), 2)
		assert.Equal(t, true, contains(list, user.Id))
	})

	t.Run("Get user with id = 1, OK", func(t *testing.T) {
		user := domain.NewUser(1, "paperino")
		retrieved, err := usersService.Get(user.Id)

		assert.Equal(t, nil, err)
		assert.Equal(t, user, retrieved)
	})

	t.Run("Get user with id = 12, KO", func(t *testing.T) {
		user := domain.NewUser(12, "paperone")
		retrieved, err := usersService.Get(user.Id)

		assert.NotEqual(t, nil, err)
		assert.Equal(t, domain.User{}, retrieved)
	})
}

func contains(list []domain.User, id uint) bool {
	for _, item := range list {
			if item.Id == id {
					return true
			}
	}
	return false
}
