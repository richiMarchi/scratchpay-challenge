package repositories

import (
	"encoding/json"
	"errors"

	"github.com/richiMarchi/scratchpay-challenge/internal/core/domain"
)

type memkvs struct {
	kvs map[uint][]byte
}

func NewMemKvs() *memkvs {
	return &memkvs{
		kvs: map[uint][]byte{},
	}
}

func (repo *memkvs) Get(id uint) (domain.User, error) {
	if value, ok := repo.kvs[id]; ok {
		user := domain.User{}
		err := json.Unmarshal(value, &user)
		if err != nil {
			return domain.User{}, err
		}

		return user, nil
	}

	return domain.User{}, errors.New("user not found in kvs")
}

func (repo *memkvs) GetAll() ([]domain.User, error) {
	var users []domain.User
	for _, elem := range repo.kvs {
		user := domain.User{}
		if err := json.Unmarshal(elem, &user); err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo *memkvs) Create(user domain.User) error {
	if _, ok := repo.kvs[user.Id]; ok {
		return errors.New("user already in kvs")
	}

	value, err := json.Marshal(user)
	if err == nil {
		repo.kvs[user.Id] = value
	}

	return err
}
