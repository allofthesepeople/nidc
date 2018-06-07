// Code generated from QueryLanguage.g4 by ANTLR 4.7.1. DO NOT EDIT.

package language // QueryLanguage
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseQueryLanguageVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseQueryLanguageVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLanguageVisitor) VisitMatchClause(ctx *MatchClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLanguageVisitor) VisitNode(ctx *NodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLanguageVisitor) VisitFilter(ctx *FilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLanguageVisitor) VisitRelationship(ctx *RelationshipContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLanguageVisitor) VisitReturnValue(ctx *ReturnValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLanguageVisitor) VisitAliasIdentity(ctx *AliasIdentityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLanguageVisitor) VisitAlias(ctx *AliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLanguageVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLanguageVisitor) VisitFieldValue(ctx *FieldValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseQueryLanguageVisitor) VisitNodeName(ctx *NodeNameContext) interface{} {
	return v.VisitChildren(ctx)
}
