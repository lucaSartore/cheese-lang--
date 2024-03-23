package test

import (
	"testing"

	"github.com/lucaSartore/cheese-lang/internal/tokenizer"
)

func TestTokenizer1(t *testing.T) {
	s := "\\\\this is a comment\n == = != && ^{    }\t	\n\n \t  -="
	tokens, err := tokenizer.Tokenize(s)

	if err != nil {
		t.Error(err)
		return
	}

	expected_tokens := []tokenizer.Token{
		tokenizer.MakeTokenWithMessage(tokenizer.Comment, "this is a comment"),
		tokenizer.MakeToken(tokenizer.EqualOperator),
		tokenizer.MakeToken(tokenizer.AssignOperator),
		tokenizer.MakeToken(tokenizer.UnEqualOperator),
		tokenizer.MakeToken(tokenizer.AndOperator),
		tokenizer.MakeToken(tokenizer.ExorOperator),
		tokenizer.MakeToken(tokenizer.OpenCurlyBracket),
		tokenizer.MakeToken(tokenizer.CloseCurlyBracket),
		tokenizer.MakeToken(tokenizer.SubOperator),
		tokenizer.MakeToken(tokenizer.AssignOperator),
	}

	for i, token := range tokens {
		if token.TokenType != expected_tokens[i].TokenType {
			t.Errorf("Expected token type %v, got %v", expected_tokens[i].TokenType, token.TokenType)
		}
		if token.TokenVal != expected_tokens[i].TokenVal {
			t.Errorf("Expected token value %v, got %v", expected_tokens[i].TokenVal, token.TokenVal)
		}
	}
}