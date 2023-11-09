package entity

import (
	"encoding/json"
)

type User struct {
	UuidEntity
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(name string, email string, password string) *User {
	user := User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	return &user
}

func (user User) GetId() string {
	return user.UuidEntity.ID.String()
}

func (user User) GetAttributes() (map[string]string, error) {
	attr := make(map[string]string)

	return attr, nil
}

func (user User) GetData() ([]byte, error) {
	data, err := json.Marshal(user)

	if err != nil {
		log.Printf("Error parsing User into JSON")
		return []byte{}, err
	}

	return data, nil
}
