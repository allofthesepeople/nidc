package nidc

import (
	"errors"
	"fmt"
	"github.com/tyndyll/nidc/language"
)


type sqlQueryVisitor struct {
	Query *Query
	Error error
}

func (visitor *sqlQueryVisitor) VisitQuery(ctx *language.QueryContext) interface{} {
	visitor.Query = &Query{}

	for _, child := range ctx.GetChildren() {
		switch child.(type) {
		case *language.MatchClauseContext:
			visitor.VisitMatchClause(child.(*language.MatchClauseContext))
		}

		if visitor.Error != nil {
			return visitor.Error
		}
	}
	return nil
}

func (visitor *sqlQueryVisitor) VisitMatchClause(ctx *language.MatchClauseContext) interface{} {
	for _, child := range ctx.GetChildren() {
		switch child.(type) {
		case *language.NodeContext:
			visitor.VisitNode(child.(*language.NodeContext))
		case *language.ReturnValueContext:
			visitor.VisitReturnValue(child.(*language.ReturnValueContext))
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

func (visitor *sqlQueryVisitor) VisitReturnValue(ctx *language.ReturnValueContext) interface{} {
	var alias string
	field := `*`
	for _, child := range ctx.GetChildren() {
		switch child.(type) {
		case *language.AliasContext:
			alias = child.(*language.AliasContext).GetText()
		case *language.FieldContext:
			field = child.(*language.FieldContext).GetText()
		}
	}

	for _, table := range visitor.Query.Tables {
		if table.Alias == alias {
			table.Field = append(table.Field, field)
			return nil
		}
	}
	// The alias wasn't found. Raise an error
	visitor.Error = errors.New(fmt.Sprintf(`Alias %s not found in query`, alias))
	return nil
}
