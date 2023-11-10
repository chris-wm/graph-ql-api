package repository

import (
	"log"

	"github.com/electivetechnology/utility-library-go/db/sql"
	"gitlab.et-ns.net/connect/graph-ql-api/internal/adapter"
	"gitlab.et-ns.net/connect/graph-ql-api/internal/entity"
)

type UserRepository struct {
}

func (r UserRepository) Create(user *entity.User) (*entity.User, error) {
	result := adapter.GetDb().Create(&user)

	return user, result.Error
}

func (r UserRepository) Delete(user *entity.User) (*entity.User, error) {
	result := adapter.GetDb().Delete(&user)

	return user, result.Error
}

func (r UserRepository) Update(user *entity.User) (*entity.User, error) {
	db := GetAdaptor()

	result := db.Updates(user).First(&user)

	return user, result.Error
}

func (r UserRepository) Read(user *entity.User) (*entity.User, error) {
	db := GetAdaptor()

	result := db.Where(user).First(&user)

	return user, result.Error
}

func (r UserRepository) GetUsers(
	limit int,
	offset int) ([]map[string]interface{}, error) {
	// Prepare storage for results
	var results []map[string]interface{}

	// Create new query for selected table
	q := sql.NewQuery("users")

	// Crete and add new filed Map
	fieldMap := map[string]string{
		"*": "users",
	}
	q.FieldMap = fieldMap

	// Apply Limit and Offset
	q.Limit = limit
	q.Offset = offset

	// Prepare query
	q.Prepare()

	// Fetch results
	log.Printf("Running query: '%s'", q.GetSql())
	adapter.GetDb().Raw(q.GetSql()).Find(&results)

	return results, nil
}
