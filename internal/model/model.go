package model

import (
	"strconv"

	"github.com/electivetechnology/utility-library-go/data"
	"github.com/electivetechnology/utility-library-go/logger"
)

var log logger.Logging

func init() {
	// Add generic logger
	log = logger.NewLogger("model")
}

func addFiltersFromOptions(filters map[string]data.Filter, options map[string]interface{}) map[string]data.Filter {
	// add visibility
	for k, f := range addVisibilityFiltersFromOptions(options) {
		filters[k] = f
	}

	// add deleted at
	for k, f := range addDeleteAtFiltersFromOptions(options) {
		filters[k] = f
	}

	return filters
}

func addVisibilityFiltersFromOptions(options map[string]interface{}) map[string]data.Filter {
	filters := make(map[string]data.Filter)

	// Add organisation filter
	if options["organisation"] != nil {
		// Convert value to string
		value := options["organisation"].(string)

		// Create new filters
		f1 := data.NewFilter()
		f2 := data.NewFilter()
		f2.Logic = data.FILTER_LOGIC_UNION

		// Protected and public visibility
		c1 := data.Criterion{
			Logic:   data.CRITERION_LOGIC_UNION,
			Key:     "visibility",
			Operand: data.CRITERION_OP_EQ,
			Type:    data.CRITERION_TYPE_VALUE,
			Value:   "protected",
		}

		c2 := data.Criterion{
			Logic:   data.CRITERION_LOGIC_UNION,
			Key:     "visibility",
			Operand: data.CRITERION_OP_EQ,
			Type:    data.CRITERION_TYPE_VALUE,
			Value:   "public",
		}

		// Private and organisation
		c3 := data.Criterion{
			Logic:   data.CRITERION_LOGIC_INTERSECTION,
			Key:     "organisation",
			Operand: data.CRITERION_OP_EQ,
			Type:    data.CRITERION_TYPE_VALUE,
			Value:   value,
		}

		c4 := data.Criterion{
			Logic:   data.CRITERION_LOGIC_INTERSECTION,
			Key:     "visibility",
			Operand: data.CRITERION_OP_EQ,
			Type:    data.CRITERION_TYPE_VALUE,
			Value:   "private",
		}

		// Add Criterion to list
		f1.Criterions = append(f1.Criterions, c1)
		f1.Criterions = append(f1.Criterions, c2)
		f2.Criterions = append(f2.Criterions, c3)
		f2.Criterions = append(f2.Criterions, c4)

		f1.Filters = make(map[string]*data.Filter)
		// Add private and org filter
		f1.Filters["visibility_2"] = f2

		// Add public and protected filter
		filters["visibility_1"] = *f1
	}

	return filters
}

func addDeleteAtFiltersFromOptions(options map[string]interface{}) map[string]data.Filter {
	filters := make(map[string]data.Filter)

	// Add deleted filter
	if options["deleted"] != nil {
		// Convert value to string
		value := options["deleted"].(bool)

		// to allow deleted we have to reverse this value
		if !value {
			// Create new filter
			f2 := data.NewFilter()

			// Create Criterion for DeletedAt
			c2 := data.Criterion{
				Logic:   data.CRITERION_LOGIC_INTERSECTION,
				Key:     "deleted_at",
				Operand: data.CRITERION_OP_BOOL,
				Type:    data.CRITERION_TYPE_VALUE,
				Value:   strconv.FormatBool(value),
			}

			// Add Criterion to list
			f2.Criterions = append(f2.Criterions, c2)

			// Add filter to the list
			filters["deleted"] = *f2
		}
	}

	return filters
}
