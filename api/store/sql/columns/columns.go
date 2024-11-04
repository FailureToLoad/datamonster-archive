package columns

type SmallInt struct{ numericColumn[int16] }
type Integer struct{ numericColumn[int32] }
type BigInt struct{ numericColumn[int64] }
type Serial struct{ numericColumn[int32] }
type BigSerial struct{ numericColumn[int64] }

type Text struct{ lexicColumn }
type VarChar struct{ lexicColumn }

func NewIntegerColumn(name string) Integer {
	return Integer{numericColumn[int32]{Name: name}}
}

func NewTextColumn(name string) Text {
	return Text{lexicColumn{Name: name}}
}
