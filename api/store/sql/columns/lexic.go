package columns

// lexicColumn represents a text column
type lexicColumn struct {
	Name string
}

func (f lexicColumn) GetName() string {
	return f.Name
}

func (f lexicColumn) Equals(value string) Condition {
	return Predicate{
		Field: f.Name,
		Op:    "=",
		Value: value,
	}
}

func (f lexicColumn) Like(value string, caseInsensitive bool) Condition {
	operation := "LIKE"
	if caseInsensitive {
		operation = "ILIKE"
	}
	return Predicate{
		Field: f.Name,
		Op:    operation,
		Value: value,
	}
}

func (f lexicColumn) In(values ...string) Condition {
	return Predicate{
		Field: f.Name,
		Op:    "ANY",
		Value: values,
	}
}
