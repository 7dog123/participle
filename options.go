package participle

import (
	"github.com/7dog123/participle/lexer"
)

// An Option to modify the behaviour of the Parser.
type Option func(p *Parser) error

// Lexer is an Option that sets the lexer to use with the given grammar.
func Lexer(def lexer.Definition) Option {
	return func(p *Parser) error {
		p.lex = def
		return nil
	}
}

// UseLookahead allows branch lookahead up to "n" tokens.
//
// If parsing cannot be disambiguated before "n" tokens of lookahead, parsing will fail.
//
// Note that increasing lookahead has a minor performance impact, but also
// reduces the accuracy of error reporting.
func UseLookahead(n int) Option {
	return func(p *Parser) error {
		p.useLookahead = n
		return nil
	}
}

// CaseInsensitive allows the specified token types to be matched case-insensitively.
func CaseInsensitive(tokens ...string) Option {
	return func(p *Parser) error {
		for _, token := range tokens {
			p.caseInsensitive[token] = true
		}
		return nil
	}
}

// ParseOption modifies how an individual parse is applied.
type ParseOption func(p *parseContext)

// AllowTrailing tokens without erroring.
//
// That is, do not error if a full parse completes but additional tokens remain.
func AllowTrailing(ok bool) ParseOption {
	return func(p *parseContext) {
		p.allowTrailing = ok
	}
}
