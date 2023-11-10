package model

import (
	"gitlab.et-ns.net/connect/graph-ql-api/internal/entity"
	"gitlab.et-ns.net/connect/graph-ql-api/internal/repository"
)

func CreateUser(t entity.User) (*entity.User, error) {
	// Create new User entity
	user := entity.NewUser(t.Email, t.Name, t.Password)

	// Save to database
	repository := repository.UserRepository{}

	log.Printf("Saving user... %v", user)
	user, err := repository.Create(user)

	return user, err
}

func ReadUser(t entity.User) (*entity.User, error) {
	// Save from database
	repository := repository.UserRepository{}

	user, err := repository.Read(&t)

	return user, err
}

func UpdateUser(t entity.User) (*entity.User, error) {
	// Update to database
	repository := repository.UserRepository{}

	log.Printf("Saving user... %v", t)
	user, err := repository.Update(&t)

	return user, err
}

func DeleteUser(t entity.User) (*entity.User, error) {
	// Delete from database
	repository := repository.UserRepository{}

	user, err := repository.Delete(&t)

	return user, err
}

func ListUser(
	limit int,
	offset int,
	options map[string]interface{}) ([]map[string]interface{}, error) {
	repository := repository.UserRepository{}

	users, err := repository.GetUsers(limit, offset)

	return users, err
}
