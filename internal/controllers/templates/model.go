package templates

import (
	"github.com/electivetechnology/utility-library-go/data"
	"gitlab.et-ns.net/connect/graph-ql-api/internal/entity"
	"gitlab.et-ns.net/connect/graph-ql-api/internal/model"
)

func CreateTemplate(c CreateRequest) (*entity.Template, error) {
	t := entity.Template{
		Type:       c.Type,
		Visibility: c.Visibility,
		Content:    c.Content,
		Name:       c.Name,
	}

	template, err := model.CreateTemplate(t)

	return template, err
}

func ReadTemplate(templateId string) (*entity.Template, error) {
	t := entity.Template{}
	t.ID = templateId

	template, err := model.ReadTemplate(t)

	return template, err
}

func UpdateTemplate(templateId string, u UpdateRequest) (*entity.Template, error) {
	t := entity.Template{}
	t.ID = templateId

	if u.Type != "" {
		t.Type = u.Type
	}
	if u.Visibility != "" {
		t.Visibility = u.Visibility
	}
	if u.Content != nil {
		t.Content = u.Content
	}
	if u.Name != "" {
		t.Name = u.Name
	}

	template, err := model.UpdateTemplate(t)

	return template, err
}

func DeleteTemplate(templateId string) (*entity.Template, error) {
	t := entity.Template{}
	t.ID = templateId

	template, err := model.DeleteTemplate(t)

	return template, err
}

func ListTemplate(
	limit int,
	offset int,
	filters map[string]data.Filter,
	sorts map[string]data.Sort,
	displays map[string]data.Display) ([]map[string]interface{}, error) {

	// Generate options
	options := make(map[string]interface{})

	templates, err := model.ListTemplate(limit, offset, filters, sorts, displays, options)

	return templates, err
}
