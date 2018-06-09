package nidc

import (
	"strings"
	"github.com/tyndyll/nidc/language"
)

type Query struct {
	Tables []*Table
}

func (q *Query) SQL() string {
	var sql strings.Builder
	sql.WriteString(`SELECT `)
	for _, fields := range q.Tables {
		sql.WriteString(fields.returnSelect())
	}
	sql.WriteString(` FROM `)
	for _, tables := range q.Tables {
		sql.WriteString(tables.returnFrom())
	}
	sql.WriteString(` WHERE `)

	whereClauses := []string{}
	for _, tables := range q.Tables {
		whereClauses = append(whereClauses, tables.returnWhere())
	}
	sql.WriteString(strings.Join(whereClauses, ` AND `))

	return sql.String()
}

func NewQuery(tree *language.QueryContext) (*Query, error) {
	visitor := &sqlQueryVisitor{}
	visitor.VisitQuery(tree)
	return visitor.Query, visitor.Error
}

