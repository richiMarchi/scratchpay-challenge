package domain

type User struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func NewUser(id uint, name string) User {
	return User{
		Id:   id,
		Name: name,
	}
}
