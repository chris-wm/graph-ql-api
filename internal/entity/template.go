package entity

import (
	"encoding/json"

	"github.com/electivetechnology/utility-library-go/validation"
)

type Template struct {
	BaseEntity
	Type         string  `json:"type"`
	Name         string  `json:"name"`
	Visibility   string  `json:"visibility"`
	Organisation string  `json:"organisation"`
	Content      *string `json:"content"`
}

func NewTemplate(templateType string, name string, visibility string, organisation string, content *string) *Template {
	template := Template{
		Type:         templateType,
		Name:         name,
		Visibility:   visibility,
		Organisation: organisation,
		Content:      content,
	}

	return &template
}

func (template Template) GetAttributes() (map[string]string, error) {
	attr := make(map[string]string)

	return attr, nil
}

func (template Template) GetData() ([]byte, error) {
	data, err := json.Marshal(template)

	if err != nil {
		log.Printf("Error parsing Template into JSON")
		return []byte{}, err
	}

	return data, nil
}

func (template Template) GetClassFilters() []string {
	t := []string{
		"id",
		"type",
		"name",
		"visibility",
		"organisation",
		"content",
		"created_at",
		"updated_at",
		"deleted_at",
	}

	return t
}

func (template Template) GetAliasClassFilters() []string {
	t := []string{
		"template.id",
		"template.type",
		"template.name",
		"template.visibility",
		"template.organisation",
		"template.content",
		"template.created_at",
		"template.updated_at",
		"template.deleted_at",
	}

	return t
}

func (template Template) GetValidFilters() validation.ValidatorRequirements {
	r := validation.Requirements{}
	r.Fields = template.GetClassFilters()

	return r
}

func (template Template) GetValidSorts() validation.ValidatorRequirements {
	// Use the same fields as for filters

	return template.GetValidFilters()
}

func (template Template) GetValidDisplays() validation.ValidatorRequirements {
	// Use the same fields as for filters

	return template.GetValidFilters()
}
