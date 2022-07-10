package ports

import "github.com/richiMarchi/scratchpay-challenge/internal/core/domain"

type UsersRepository interface {
	Get(id uint) (domain.User, error)
	GetAll() ([]domain.User, error)
	Create(domain.User) error
}
