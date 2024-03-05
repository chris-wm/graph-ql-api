package repository

import (
	"log"

	"github.com/electivetechnology/utility-library-go/data"
	"github.com/electivetechnology/utility-library-go/db/sql"
	"gitlab.et-ns.net/connect/graph-ql-api/internal/adapter"
	"gitlab.et-ns.net/connect/graph-ql-api/internal/entity"
)

type TemplateRepository struct {
}

func (r TemplateRepository) Create(template *entity.Template) (*entity.Template, error) {
	result := adapter.GetDb().Create(&template)

	return template, result.Error
}

func (r TemplateRepository) Delete(template *entity.Template) (*entity.Template, error) {
	result := adapter.GetDb().Delete(&template)

	return template, result.Error
}

func (r TemplateRepository) Update(template *entity.Template) (*entity.Template, error) {
	db := GetAdaptor()

	result := db.Updates(template).First(&template)

	return template, result.Error
}

func (r TemplateRepository) Read(template *entity.Template) (*entity.Template, error) {
	db := GetAdaptor()

	result := db.Where(template).First(&template)

	return template, result.Error
}

func (r TemplateRepository) GetTemplates(
	limit int,
	offset int,
	filters map[string]data.Filter,
	sorts map[string]data.Sort,
	displays map[string]data.Display) ([]map[string]interface{}, error) {
	// Prepare storage for results
	var results []map[string]interface{}

	// Create new query for selected table
	q := sql.NewQuery("templates")

	// Crete and add new filed Map
	fieldMap := map[string]string{
		"*": "templates",
	}
	q.FieldMap = fieldMap

	// Apply Limit and Offset
	q.Limit = limit
	q.Offset = offset

	// Apply filters
	q.Filters = filters

	// Apply Sorts
	q.Sorts = sorts

	// Apply Displays
	q.Displays = displays

	// Prepare query
	q.Prepare()

	// Fetch results
	log.Printf("Running query: '%s'", q.GetSql())
	adapter.GetDb().Raw(q.GetSql()).Find(&results)

	return results, nil
}
