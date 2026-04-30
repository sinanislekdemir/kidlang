package interpreter

import "testing"

func TestTokenizeSyntaxHighlightsFunctions(t *testing.T) {
	keywords := map[string]bool{
		"FUNCTION": true,
		"RETURN":   true,
		"BOX":      true,
	}

	tokens := tokenizeSyntax("function cheer(box name)", keywords)

	if len(tokens) < 4 {
		t.Fatalf("tokenizeSyntax() returned too few tokens: %#v", tokens)
	}
	if tokens[0].Kind != syntaxKeyword || tokens[0].Text != "function" {
		t.Fatalf("first token = %#v, want function keyword", tokens[0])
	}
	if tokens[2].Kind != syntaxFunction || tokens[2].Text != "cheer" {
		t.Fatalf("function name token = %#v, want function highlight", tokens[2])
	}

	callTokens := tokenizeSyntax("print cheer(box name)", map[string]bool{
		"PRINT": true,
		"BOX":   true,
	})

	foundFunctionCall := false
	for _, token := range callTokens {
		if token.Text == "cheer" && token.Kind == syntaxFunction {
			foundFunctionCall = true
			break
		}
	}
	if !foundFunctionCall {
		t.Fatalf("tokenizeSyntax() did not highlight function call: %#v", callTokens)
	}
}

func TestTokenizeSyntaxHighlightsREMComments(t *testing.T) {
	tokens := tokenizeSyntax("  REM explain the game", map[string]bool{})

	if len(tokens) != 2 {
		t.Fatalf("tokenizeSyntax() returned %#v, want indentation and REM comment", tokens)
	}
	if tokens[0].Kind != syntaxNormal || tokens[0].Text != "  " {
		t.Fatalf("indent token = %#v, want leading spaces", tokens[0])
	}
	if tokens[1].Kind != syntaxComment || tokens[1].Text != "REM explain the game" {
		t.Fatalf("comment token = %#v, want REM comment", tokens[1])
	}
}
