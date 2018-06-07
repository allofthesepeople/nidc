// Code generated from QueryLanguage.g4 by ANTLR 4.7.1. DO NOT EDIT.

package language // QueryLanguage
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseQueryLanguageListener is a complete listener for a parse tree produced by QueryLanguageParser.
type BaseQueryLanguageListener struct{}

var _ QueryLanguageListener = &BaseQueryLanguageListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseQueryLanguageListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseQueryLanguageListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseQueryLanguageListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseQueryLanguageListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseQueryLanguageListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseQueryLanguageListener) ExitQuery(ctx *QueryContext) {}

// EnterMatchClause is called when production matchClause is entered.
func (s *BaseQueryLanguageListener) EnterMatchClause(ctx *MatchClauseContext) {}

// ExitMatchClause is called when production matchClause is exited.
func (s *BaseQueryLanguageListener) ExitMatchClause(ctx *MatchClauseContext) {}

// EnterNode is called when production node is entered.
func (s *BaseQueryLanguageListener) EnterNode(ctx *NodeContext) {}

// ExitNode is called when production node is exited.
func (s *BaseQueryLanguageListener) ExitNode(ctx *NodeContext) {}

// EnterFilter is called when production filter is entered.
func (s *BaseQueryLanguageListener) EnterFilter(ctx *FilterContext) {}

// ExitFilter is called when production filter is exited.
func (s *BaseQueryLanguageListener) ExitFilter(ctx *FilterContext) {}

// EnterRelationship is called when production relationship is entered.
func (s *BaseQueryLanguageListener) EnterRelationship(ctx *RelationshipContext) {}

// ExitRelationship is called when production relationship is exited.
func (s *BaseQueryLanguageListener) ExitRelationship(ctx *RelationshipContext) {}

// EnterReturnValue is called when production returnValue is entered.
func (s *BaseQueryLanguageListener) EnterReturnValue(ctx *ReturnValueContext) {}

// ExitReturnValue is called when production returnValue is exited.
func (s *BaseQueryLanguageListener) ExitReturnValue(ctx *ReturnValueContext) {}

// EnterAliasIdentity is called when production aliasIdentity is entered.
func (s *BaseQueryLanguageListener) EnterAliasIdentity(ctx *AliasIdentityContext) {}

// ExitAliasIdentity is called when production aliasIdentity is exited.
func (s *BaseQueryLanguageListener) ExitAliasIdentity(ctx *AliasIdentityContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseQueryLanguageListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseQueryLanguageListener) ExitAlias(ctx *AliasContext) {}

// EnterField is called when production field is entered.
func (s *BaseQueryLanguageListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseQueryLanguageListener) ExitField(ctx *FieldContext) {}

// EnterFieldValue is called when production fieldValue is entered.
func (s *BaseQueryLanguageListener) EnterFieldValue(ctx *FieldValueContext) {}

// ExitFieldValue is called when production fieldValue is exited.
func (s *BaseQueryLanguageListener) ExitFieldValue(ctx *FieldValueContext) {}

// EnterNodeName is called when production nodeName is entered.
func (s *BaseQueryLanguageListener) EnterNodeName(ctx *NodeNameContext) {}

// ExitNodeName is called when production nodeName is exited.
func (s *BaseQueryLanguageListener) ExitNodeName(ctx *NodeNameContext) {}
