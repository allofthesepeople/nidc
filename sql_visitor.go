package nidc

import (
	"errors"
	"fmt"
	"github.com/tyndyll/nidc/language"
	"strings"
)

type Query struct {
	Tables []*Table
}

func (q *Query) SQL() (string) {
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

type Identity struct {
	Name   string
	Alias  string
	Field  []string
	Filter []*Filter
}

type Filter struct {
	Name  string
	Value interface{}
}

type Table struct {
	*Identity
}

func (table *Table) returnSelect() string {
	alias := table.Alias
	if alias == "" {
		alias = table.Name
	}

	if table.Field == nil || len(table.Field) == 0 {
		return fmt.Sprintf(`%s.*`, alias)
	}

	fields := make([]string, len(table.Field))
	for i, field := range table.Field {
		fields[i] = fmt.Sprintf(`%s.%s`, alias, field)
	}
	return strings.Join(fields, `, `)
}

func (table *Table) returnFrom() string {
	if table.Alias != "" {
		return fmt.Sprintf(`%s AS %s`, table.Name, table.Alias)
	}
	return fmt.Sprintf(`%s `, table.Name)
}

func (table *Table) returnWhere() string {
	if len(table.Filter) == 0 {
		return ""
	}

	alias := table.Alias
	if alias == "" {
		alias = table.Name
	}

	filters := make([]string, len(table.Filter))
	for i, filter := range table.Filter {
		filters[i] = fmt.Sprintf(`%s.%s = %s`, alias, filter.Name, filter.Value)
	}

	return strings.Join(filters, ` AND `)
}

type Relationship struct {
	*Identity
}

type sqlQueryVisitor struct {
	Query *Query
	Error error
}

func (visitor *sqlQueryVisitor) VisitQuery(ctx *language.QueryContext) interface{} {
	visitor.Query = &Query{}

	for _, child := range ctx.GetChildren() {
		switch child.(type) {
		case *language.MatchClauseContext:
			return visitor.VisitMatchClause(child.(*language.MatchClauseContext))
		}
	}
	visitor.Error = errors.New("No valid query type found")
	return nil
}

func (visitor *sqlQueryVisitor) VisitMatchClause(ctx *language.MatchClauseContext) interface{} {
	for _, child := range ctx.GetChildren() {
		switch child.(type) {
		case *language.NodeContext:
			return visitor.VisitNode(child.(*language.NodeContext))
		}
	}
	return nil
}

func (visitor *sqlQueryVisitor) VisitNode(ctx *language.NodeContext) interface{} {
	var identity *Identity

	for _, child := range ctx.GetChildren() {
		switch child.(type) {
		case *language.AliasIdentityContext:
			identity = visitor.VisitAliasIdentity(child.(*language.AliasIdentityContext)).(*Identity)
			if identity == nil {
				return nil
			}

		case *language.FilterContext:
			identity.Filter = append(identity.Filter, visitor.VisitFilter(child.(*language.FilterContext)).(*Filter))
		}
	}

	visitor.Query.Tables = append(visitor.Query.Tables, &Table{identity})
	return nil
}

func (visitor *sqlQueryVisitor) VisitAliasIdentity(ctx *language.AliasIdentityContext) interface{} {
	identity := &Identity{
		Filter: []*Filter{},
	}

	for _, child := range ctx.GetChildren() {
		switch child.(type) {
		case *language.AliasContext:
			identity.Alias = child.(*language.AliasContext).GetText()
		case *language.NodeNameContext:
			identity.Name = child.(*language.NodeNameContext).GetText()
		}
	}
	return identity
}

func (visitor *sqlQueryVisitor) VisitFilter(ctx *language.FilterContext) interface{} {
	filter := &Filter{}
	for _, child := range ctx.GetChildren() {
		switch child.(type) {
		case *language.FieldContext:
			filter.Name = child.(*language.FieldContext).GetText()
		case *language.FieldValueContext:
			filter.Value = child.(*language.FieldValueContext).GetText()
		}
	}
	return filter
}
