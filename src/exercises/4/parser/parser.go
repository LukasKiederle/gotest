package parser

import (
	"fmt"

	"github.com/LukasKiederle/gotest/src/exercises/4/lexer"
)

type node interface {
	Eval(vars map[string]bool) bool
}

type val struct {
	name string
}

type or struct {
	lhs node
	rhs node
}

type and struct {
	lhs node
	rhs node
}

type not struct {
	factor node
}

type Parser struct {
	rootNode node
	lexer    *lexer.Lexer
	token    string
}

func NewParser(lexer *lexer.Lexer) *Parser {
	b := Parser{lexer: lexer}
	b.parse()
	return &b
}

func (p *Parser) parse() {
	p.expression()
}

func (p *Parser) term() {
	p.factor()

	if p.token == "&" { //maybe the term is followed 'or' and another term
		lhs := p.rootNode
		p.factor()
		rhs := p.rootNode
		p.rootNode = &and{lhs, rhs}
	}

}

func (p *Parser) factor() {
	token := p.lexer.NextToken()

	switch token {
	case "":
		return
	case "!":
		p.factor()
		p.rootNode = &not{p.rootNode}
		break
	case "(":
		p.expression()
		p.token = p.lexer.NextToken()
		break
	default:
		p.rootNode = &val{token}
		p.token = p.lexer.NextToken()
	}
}

func (p *Parser) Eval(vars map[string]bool) bool {
	return p.rootNode.Eval(vars)
}

func (p *Parser) String() string {
	return fmt.Sprintf("%v", p.rootNode)
}

func (o *or) Eval(vars map[string]bool) bool {
	return o.lhs.Eval(vars) || o.rhs.Eval(vars)
}

func (a *and) Eval(vars map[string]bool) bool {
	return a.lhs.Eval(vars) && a.rhs.Eval(vars)
}

func (n *not) Eval(vars map[string]bool) bool {
	return !n.factor.Eval(vars)
}

func (v *val) Eval(vars map[string]bool) bool {
	return vars[v.name] // missing vars will be evaluated to false!
}

func (p *Parser) expression() {
	p.term()             //an expression always has a term
	for p.token == "|" { //maybe the term is followed 'or' and another term
		lhs := p.rootNode
		p.term()
		rhs := p.rootNode
		p.rootNode = &or{lhs, rhs}
	}
}
