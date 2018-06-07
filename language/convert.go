package language

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Parse takes a query string and attempts to lex and parse it. If successful it will return a *QueryContext AST. If
// it fails it will return nil and the error returned by the lexer
func Parse(query string) (*QueryContext, error) {
	input := antlr.NewInputStream(query)
	lexer := NewQueryLanguageLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	errorListener := &ErrorListener{}
	p := NewQueryLanguageParser(stream)

	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	queryAST := p.Query().(*QueryContext)

	if errorListener.Error != nil {
		return nil, errorListener.Error
	}

	return queryAST, nil
}

// ASTtoQuery takes a QueryContext and returns that context as a string, removing the <EOF> signifier from the end of
// the string
func ASTtoQuery(ctx *QueryContext) string {
	query := ctx.GetText()
	if strings.HasSuffix(query, `<EOF>`) {
		query = query[0 : len(query)-5]
	}
	return query
}
