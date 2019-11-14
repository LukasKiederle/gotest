package parser

import (
	"fmt"
	"testing"

	"github.com/LukasKiederle/gotest/src/exercises/4/lexer"
)

func TestParser_Eval(t *testing.T) {
	// create parser
	p := NewParser(lexer.NewLexer("a & b & !c"))

	// set vars
	vars := map[string]bool{
		"a": true,
		"b": true,
		"c": false,
	}
	if p.Eval(vars) != true {
		t.Error(fmt.Sprintf("Wrong result detected"))
	}

	// set vars
	vars = map[string]bool{
		"a": true,
		"b": true,
		"c": true,
	}

	if p.Eval(vars) != false {
		t.Error(fmt.Sprintf("Wrong result detected"))
	}

	// set vars
	vars = map[string]bool{
		"a": false,
		"b": false,
		"c": false,
	}
	if p.Eval(vars) != false {
		t.Error(fmt.Sprintf("Wrong result detected"))
	}

	p = NewParser(lexer.NewLexer("a & (b | c & b) & d"))

	// set vars
	vars = map[string]bool{
		"a": true,
		"b": true,
		"c": false,
		"d": true,
	}
	if p.Eval(vars) != true {
		t.Error(fmt.Sprintf("Wrong result detected"))
	}

	// test string support
	if p.String() != "&(&('a',|('b',&('c','b'))),'d')" {
		t.Error(fmt.Sprintf("Wrong string representation: %v", p.String()))
	}

	// p = NewParser(NewLexer("a &"))
	// if p.String() == "&('a','a')" {
	// 	t.Error(fmt.Sprintf("The string \"a &\" parses to: %s", p.String()))
	// }
}
