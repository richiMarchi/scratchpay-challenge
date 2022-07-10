package ports

import "github.com/richiMarchi/scratchpay-challenge/internal/core/domain"

type UsersService interface {
	Get(id uint) (domain.User, error)
	Create(id uint, name string) error
	List() ([]domain.User, error)
}
