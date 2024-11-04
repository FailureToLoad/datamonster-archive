package columns

import (
	"fmt"
	"strings"
)

// Condition represents a single condition or a group of Conditions
type Condition interface {
	Build(paramStart int) (string, []interface{}, int)
}

// Composite represents a group of Conditions joined by an Operator
type Composite struct {
	Conditions []Condition
	Operator   string
}

func (c Composite) Build(paramStart int) (string, []interface{}, int) {
	if len(c.Conditions) == 0 {
		return "", nil, paramStart
	}

	clauses := make([]string, 0, len(c.Conditions))
	var allArgs []interface{}
	currentParam := paramStart

	for _, cond := range c.Conditions {
		clause, args, nextParam := cond.Build(currentParam)
		if clause != "" {
			clauses = append(clauses, clause)
			allArgs = append(allArgs, args...)
			currentParam = nextParam
		}
	}

	if len(clauses) == 0 {
		return "", nil, paramStart
	}

	result := strings.Join(clauses, fmt.Sprintf(" %s ", c.Operator))
	if len(clauses) > 1 {
		result = fmt.Sprintf("(%s)", result)
	}

	return result, allArgs, currentParam
}

// And creates a composite AND condition
func And(conditions ...Condition) Condition {
	return Composite{
		Conditions: conditions,
		Operator:   "AND",
	}
}

// Or creates a composite OR condition
func Or(conditions ...Condition) Condition {
	return Composite{
		Conditions: conditions,
		Operator:   "OR",
	}
}
