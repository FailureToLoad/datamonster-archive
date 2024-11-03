package columns

import (
	"fmt"
	"strings"
)

// Condition represents a single condition or a group of conditions
type Condition interface {
	Build(paramStart int) (string, []interface{}, int)
}

// Composite represents a group of conditions joined by an operator
type Composite struct {
	conditions []Condition
	operator   string
}

func (c Composite) Build(paramStart int) (string, []interface{}, int) {
	if len(c.conditions) == 0 {
		return "", nil, paramStart
	}

	clauses := make([]string, 0, len(c.conditions))
	var allArgs []interface{}
	currentParam := paramStart

	for _, cond := range c.conditions {
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

	result := strings.Join(clauses, fmt.Sprintf(" %s ", c.operator))
	if len(clauses) > 1 {
		result = fmt.Sprintf("(%s)", result)
	}

	return result, allArgs, currentParam
}

// And creates a composite AND condition
func And(conditions ...Condition) Condition {
	return Composite{
		conditions: conditions,
		operator:   "AND",
	}
}

// Or creates a composite OR condition
func Or(conditions ...Condition) Condition {
	return Composite{
		conditions: conditions,
		operator:   "OR",
	}
}
