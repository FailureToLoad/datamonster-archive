package sql

type Table struct {
	Name string
}

func (t Table) Delete() *DeleteCommand {
	return &DeleteCommand{
		TableName: t.Name,
	}
}

func (t Table) Insert() *InsertCommand {
	return &InsertCommand{
		TableName: t.Name,
	}
}

func (t Table) Select(cols ...string) *SelectQuery {
	if len(cols) == 0 {
		cols = []string{"*"}
	}
	return &SelectQuery{
		TableName:  t.Name,
		Selections: cols,
	}
}

func (t Table) Update() *UpdateQuery {
	return &UpdateQuery{
		TableName: t.Name,
	}
}
