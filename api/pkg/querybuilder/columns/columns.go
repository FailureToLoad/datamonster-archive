package columns

type SmallInt struct{ numericColumn[int16] }
type Integer struct{ numericColumn[int32] }
type BigInt struct{ numericColumn[int64] }
type Serial struct{ numericColumn[int32] }
type BigSerial struct{ numericColumn[int64] }

type Text struct{ lexicColumn }
type VarChar struct{ lexicColumn }

func NewSmallIntColumn(name string) SmallInt {
	return SmallInt{numericColumn[int16]{Name: name}}
}

func NewIntegerColumn(name string) Integer {
	return Integer{numericColumn[int32]{Name: name}}
}

func NewBigIntColumn(name string) BigInt {
	return BigInt{numericColumn[int64]{Name: name}}
}

func NewSerialColumn(name string) Serial {
	return Serial{numericColumn[int32]{Name: name}}
}

func NewBigSerialColumn(name string) BigSerial {
	return BigSerial{numericColumn[int64]{Name: name}}
}

func NewTextColumn(name string) Text {
	return Text{lexicColumn{Name: name}}
}

func NewVarCharColumn(name string) VarChar {
	return VarChar{lexicColumn{Name: name}}
}
