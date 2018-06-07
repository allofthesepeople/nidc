// Code generated from QueryLanguage.g4 by ANTLR 4.7.1. DO NOT EDIT.

package language // QueryLanguage
import "github.com/antlr/antlr4/runtime/Go/antlr"

// QueryLanguageListener is a complete listener for a parse tree produced by QueryLanguageParser.
type QueryLanguageListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterMatchClause is called when entering the matchClause production.
	EnterMatchClause(c *MatchClauseContext)

	// EnterNode is called when entering the node production.
	EnterNode(c *NodeContext)

	// EnterFilter is called when entering the filter production.
	EnterFilter(c *FilterContext)

	// EnterRelationship is called when entering the relationship production.
	EnterRelationship(c *RelationshipContext)

	// EnterReturnValue is called when entering the returnValue production.
	EnterReturnValue(c *ReturnValueContext)

	// EnterAliasIdentity is called when entering the aliasIdentity production.
	EnterAliasIdentity(c *AliasIdentityContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFieldValue is called when entering the fieldValue production.
	EnterFieldValue(c *FieldValueContext)

	// EnterNodeName is called when entering the nodeName production.
	EnterNodeName(c *NodeNameContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitMatchClause is called when exiting the matchClause production.
	ExitMatchClause(c *MatchClauseContext)

	// ExitNode is called when exiting the node production.
	ExitNode(c *NodeContext)

	// ExitFilter is called when exiting the filter production.
	ExitFilter(c *FilterContext)

	// ExitRelationship is called when exiting the relationship production.
	ExitRelationship(c *RelationshipContext)

	// ExitReturnValue is called when exiting the returnValue production.
	ExitReturnValue(c *ReturnValueContext)

	// ExitAliasIdentity is called when exiting the aliasIdentity production.
	ExitAliasIdentity(c *AliasIdentityContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFieldValue is called when exiting the fieldValue production.
	ExitFieldValue(c *FieldValueContext)

	// ExitNodeName is called when exiting the nodeName production.
	ExitNodeName(c *NodeNameContext)
}
