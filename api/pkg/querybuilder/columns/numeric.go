package columns

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// numericColumn is a generic struct that works with any numeric type
type numericColumn[T Numeric] struct {
	Name string
}

func (f numericColumn[T]) GetName() string {
	return f.Name
}

func (f numericColumn[T]) Equals(value T) Condition {
	return Predicate{
		Field: f.Name,
		Op:    "=",
		Value: value,
	}
}

func (f numericColumn[T]) GreaterThan(value T) Condition {
	return Predicate{
		Field: f.Name,
		Op:    ">",
		Value: value,
	}
}

func (f numericColumn[T]) LessThan(value T) Condition {
	return Predicate{
		Field: f.Name,
		Op:    "<",
		Value: value,
	}
}

func (f numericColumn[T]) GreaterOrEqual(value T) Condition {
	return Predicate{
		Field: f.Name,
		Op:    ">=",
		Value: value,
	}
}

func (f numericColumn[T]) LessOrEqual(value T) Condition {
	return Predicate{
		Field: f.Name,
		Op:    "<=",
		Value: value,
	}
}

func (f numericColumn[T]) In(values ...T) Condition {
	return Predicate{
		Field: f.Name,
		Op:    "ANY",
		Value: values,
	}
}
