// Code generated from QueryLanguage.g4 by ANTLR 4.7.1. DO NOT EDIT.

package language // QueryLanguage
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by QueryLanguageParser.
type QueryLanguageVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by QueryLanguageParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by QueryLanguageParser#matchClause.
	VisitMatchClause(ctx *MatchClauseContext) interface{}

	// Visit a parse tree produced by QueryLanguageParser#node.
	VisitNode(ctx *NodeContext) interface{}

	// Visit a parse tree produced by QueryLanguageParser#filter.
	VisitFilter(ctx *FilterContext) interface{}

	// Visit a parse tree produced by QueryLanguageParser#relationship.
	VisitRelationship(ctx *RelationshipContext) interface{}

	// Visit a parse tree produced by QueryLanguageParser#returnValue.
	VisitReturnValue(ctx *ReturnValueContext) interface{}

	// Visit a parse tree produced by QueryLanguageParser#aliasIdentity.
	VisitAliasIdentity(ctx *AliasIdentityContext) interface{}

	// Visit a parse tree produced by QueryLanguageParser#alias.
	VisitAlias(ctx *AliasContext) interface{}

	// Visit a parse tree produced by QueryLanguageParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by QueryLanguageParser#fieldValue.
	VisitFieldValue(ctx *FieldValueContext) interface{}

	// Visit a parse tree produced by QueryLanguageParser#nodeName.
	VisitNodeName(ctx *NodeNameContext) interface{}
}
