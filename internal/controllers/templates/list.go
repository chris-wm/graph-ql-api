package templates

import (
	"net/http"

	context "github.com/electivetechnology/utility-library-go/request/context"
	"github.com/gin-gonic/gin"
	"gitlab.et-ns.net/connect/graph-ql-api/internal/entity"
)

func List(c *gin.Context) {
	// Reserve space for errors
	var errors []error

	// Get valid filters, sorts and displays
	validFilters := entity.Template.GetValidFilters(entity.Template{})
	validSorts := entity.Template.GetValidSorts(entity.Template{})
	validDisplays := entity.Template.GetValidDisplays(entity.Template{})

	// Get Filters
	filters, err := context.GetFilters(c, validFilters)
	if err != nil {
		errors = append(errors, err)
	}
	log.Printf("Got filters for request: %v", filters.GetFilters())

	// Get Sorts
	sorts, err := context.GetSorts(c, validSorts)
	if err != nil {
		errors = append(errors, err)
	}
	log.Printf("Got sorts for request: %v", sorts.GetSorts())

	// Get Displays
	displays, err := context.GetDisplays(c, validDisplays)
	if err != nil {
		errors = append(errors, err)
	}
	log.Printf("Got displays for request: %v", displays)

	// Get Limit
	limit, err := context.GetLimit(c)
	if err != nil {
		errors = append(errors, err)
	}
	log.Printf("Got limit for request: %v", limit)

	// Get Offset
	offset, err := context.GetOffset(c)
	if err != nil {
		errors = append(errors, err)
	}
	log.Printf("Got offset for request: %v", offset)

	// Get Format
	format := context.GetFormat(c)
	log.Printf("Got format for request: %v", format)

	// Output errors (if any) and terminate
	if len(errors) > 0 {
		var list []string
		for _, e := range errors {
			list = append(list, e.Error())
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": list})
		return
	}

	// Get templates
	templates, err := ListTemplate(
		limit.GetLimit(),
		offset.GetOffset(),
		filters.GetDataFilters(),
		sorts.GetDataSorts(),
		displays.GetDataDisplays())

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, templates)
}
