package services

import (
	"errors"

	"github.com/richiMarchi/scratchpay-challenge/internal/core/domain"
	"github.com/richiMarchi/scratchpay-challenge/internal/core/ports"
)

type userService struct {
	userRepository ports.UsersRepository
}

func New(userRepo ports.UsersRepository) *userService {
	return &userService{
		userRepository: userRepo,
	}
}

func (svc *userService) List() ([]domain.User, error) {
	users, err := svc.userRepository.GetAll()
	if err != nil {
		return []domain.User{}, err
	}

	return append([]domain.User{}, users...), nil
}

func (svc *userService) Get(id uint) (domain.User, error) {
	user, err := svc.userRepository.Get(id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (svc *userService) Create(id uint, name string) error {
	if !svc.isIdPresent(id) {
		if err := svc.userRepository.Create(domain.NewUser(id, name)); err == nil {
			return err
		}

		return nil
	}

	return errors.New("id already present")
}

func (svc *userService) isIdPresent(id uint) bool {
	_, err := svc.userRepository.Get(id)
	return err == nil
}
