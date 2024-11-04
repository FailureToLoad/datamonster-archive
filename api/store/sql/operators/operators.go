package operators

import (
	"github.com/failuretoload/datamonster/store/sql/columns"
)

// And creates a composite AND condition
func And(conditions ...columns.Condition) columns.Condition {
	return columns.Composite{
		Conditions: conditions,
		Operator:   "AND",
	}
}

// Or creates a composite OR condition
func Or(conditions ...columns.Condition) columns.Condition {
	return columns.Composite{
		Conditions: conditions,
		Operator:   "OR",
	}
}
