package columns

import (
	"fmt"
)

// Predicate represents a WHERE clause constraint
type Predicate struct {
	Field string
	Op    string
	Value interface{}
}

func (c Predicate) Build(paramStart int) (string, []interface{}, int) {
	if c.Op == "ANY" {
		return fmt.Sprintf("%s = ANY($%d)", c.Field, paramStart),
			[]interface{}{c.Value},
			paramStart + 1
	}

	return fmt.Sprintf("%s %s $%d", c.Field, c.Op, paramStart),
		[]interface{}{c.Value},
		paramStart + 1
}
