package lexer

import (
    "babel/pkg/token"
    "testing"
)

func TestNextToken(t *testing.T) {
    input := `let five = 5;
    let ten = 10;
    let add = fn(x, y) {
        x + y;
    };
    let result = add(five, ten);
    !-/*5;
    5 < 10 > 5;
    10 == 10; 10 != 9; `

    tests := []struct {
        expectedType uint8 
        expectedLiteral string
    }{
        {token.LET, "let"},
        {token.IDENTIFIER, "five"},
        {token.ASSIGN, "="},
        {token.DIGIT, "5"},
        {token.SEMICOLON, ";"},
        {token.LET, "let"},
        {token.IDENTIFIER, "ten"},
        {token.ASSIGN, "="},
        {token.DIGIT, "10"},
        {token.SEMICOLON, ";"},
        {token.LET, "let"},
        {token.IDENTIFIER, "add"},
        {token.ASSIGN, "="},
        {token.FUNCTION, "fn"},
        {token.LPAREN, "("},
        {token.IDENTIFIER, "x"},
        {token.COMMA, ","},
        {token.IDENTIFIER, "y"},
        {token.RPAREN, ")"},
        {token.LBRACE, "{"},
        {token.IDENTIFIER, "x"},
        {token.PLUS, "+"},
        {token.IDENTIFIER, "y"},
        {token.SEMICOLON, ";"},
        {token.RBRACE, "}"},
        {token.SEMICOLON, ";"},
        {token.LET, "let"},
        {token.IDENTIFIER, "result"},
        {token.ASSIGN, "="},
        {token.IDENTIFIER, "add"},
        {token.LPAREN, "("},
        {token.IDENTIFIER, "five"},
        {token.COMMA, ","},
        {token.IDENTIFIER, "ten"},
        {token.RPAREN, ")"},
        {token.SEMICOLON, ";"},
        {token.BANG, "!"},
        {token.MINUS, "-"},
        {token.FORWARDSLASH, "/"},
        {token.ASTERISK, "*"},
        {token.DIGIT, "5"},
        {token.SEMICOLON, ";"},
        {token.DIGIT, "5"},
        {token.LT, "<"},
        {token.DIGIT, "10"},
        {token.GT, ">"},
        {token.DIGIT, "5"},
        {token.SEMICOLON, ";"},
        {token.DIGIT, "10"},
        {token.EQUALS, "=="},
        {token.DIGIT, "10"},
        {token.SEMICOLON, ";"},
        {token.DIGIT, "10"},
        {token.NOTEQUALS, "!="},
        {token.DIGIT, "9"},
        {token.SEMICOLON, ";"},
        {token.EOF, ""},
    }

    l := New(input)

    for i, tt := range tests {
        tok := l.NextToken()

        if tok.StringLiteral != tt.expectedLiteral {
            t.Fatalf("tests[%d] = tokentype wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.StringLiteral)
        }

        if tok.Type != tt.expectedType {
            t.Fatalf("tests[%d] = tokentype wrong. expected=%+v, got=%+v", i, tt.expectedType, tok.Type)
        }
    }
}


