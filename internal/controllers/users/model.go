package users

import (
	"github.com/google/uuid"
	"gitlab.et-ns.net/connect/graph-ql-api/internal/entity"
	"gitlab.et-ns.net/connect/graph-ql-api/internal/model"
)

func CreateUser(c CreateRequest) (*entity.User, error) {
	t := entity.User{
		Email:    c.Email,
		Password: c.Password,
		Name:     c.Name,
	}

	user, err := model.CreateUser(t)

	return user, err
}

func ReadUser(userId uuid.UUID) (*entity.User, error) {
	t := entity.User{}
	t.ID = userId

	user, err := model.ReadUser(t)

	return user, err
}

func UpdateUser(userId uuid.UUID, u UpdateRequest) (*entity.User, error) {
	t := entity.User{}
	t.ID = userId

	if u.Email != "" {
		t.Email = u.Email
	}
	if u.Password != "" {
		t.Password = u.Password
	}
	if u.Name != "" {
		t.Name = u.Name
	}

	user, err := model.UpdateUser(t)

	return user, err
}

func DeleteUser(userId uuid.UUID) (*entity.User, error) {
	t := entity.User{}
	t.ID = userId

	user, err := model.DeleteUser(t)

	return user, err
}

func ListUser(
	limit int,
	offset int) ([]map[string]interface{}, error) {

	// Generate options
	options := make(map[string]interface{})

	users, err := model.ListUser(limit, offset, options)

	return users, err
}
