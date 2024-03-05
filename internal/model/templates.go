package model

import (
	"github.com/electivetechnology/utility-library-go/data"
	"gitlab.et-ns.net/connect/graph-ql-api/internal/entity"
	"gitlab.et-ns.net/connect/graph-ql-api/internal/repository"
)

func CreateTemplate(t entity.Template) (*entity.Template, error) {
	// Create new Template entity
	template := entity.NewTemplate(t.Type, t.Name, t.Visibility, t.Organisation, t.Content)

	// Save to database
	repository := repository.TemplateRepository{}

	log.Printf("Saving template... %v", template)
	template, err := repository.Create(template)

	return template, err
}

func ReadTemplate(t entity.Template) (*entity.Template, error) {
	// Save from database
	repository := repository.TemplateRepository{}

	template, err := repository.Read(&t)

	return template, err
}

func UpdateTemplate(t entity.Template) (*entity.Template, error) {
	// Update to database
	repository := repository.TemplateRepository{}

	log.Printf("Saving template... %v", t)
	template, err := repository.Update(&t)

	return template, err
}

func DeleteTemplate(t entity.Template) (*entity.Template, error) {
	// Delete from database
	repository := repository.TemplateRepository{}

	template, err := repository.Delete(&t)

	return template, err
}

func ListTemplate(
	limit int,
	offset int,
	filters map[string]data.Filter,
	sorts map[string]data.Sort,
	displays map[string]data.Display,
	options map[string]interface{}) ([]map[string]interface{}, error) {
	repository := repository.TemplateRepository{}

	// Add organisation, visibility & deleted filters
	filters = addFiltersFromOptions(filters, options)

	templates, err := repository.GetTemplates(limit, offset, filters, sorts, displays)

	return templates, err
}
