package lexer

import (
	"babel/pkg/token"
)

type Lexer struct {
    input string
    position int
    readPosition int
    ch byte
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}


func (l *Lexer) NextToken() token.Token {
    var t  token.Token

    l.skipWhitespace()
    
    ch := l.ch

    if ch == 0 {
        return token.Token{Type: token.EOF, StringLiteral: ""}
    }

    tkType, ok := token.SingleWordKeywords[ch]
    if ok {
        twoChars := string(ch) + string(l.peekChar()) 

        if twoChars == "==" {
            l.readChar()
            t = token.Token{Type: token.EQUALS, StringLiteral: string(ch) + string(l.ch)}
            l.readChar()
            return t
        }
        if twoChars == "!=" {
            l.readChar()
            t = token.Token{Type: token.NOTEQUALS, StringLiteral: string(ch) + string(l.ch)}
            l.readChar()
            return t
        }

        t = token.Token{Type: tkType, StringLiteral: string(ch)}
        l.readChar()
        return t
    }


    // Checks for multi length keywords, Digits and identifiers.
    if isLetter(l.ch) {
        t.StringLiteral = l.readIdentifier()
        t.Type = token.LookupIdent(t.StringLiteral)
        return t
    } else if isDigit(l.ch) {
        t.Type = token.DIGIT
        t.StringLiteral = l.readNumber()
        return t
    }         

    t = token.Token{token.ILLEGAL, string(ch)}
    l.readChar()
    return t
}

func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
    position := l.position
    for isLetter(l.ch) { 
        l.readChar()
    }
    return l.input[position:l.position] 
}

func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.ch) {
        l.readChar()
    }

    return l.input[position:l.position]
}

func isLetter(ch byte) bool {
    return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
    return ch <= '9' && ch >= '0'
}

func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
}

func (l *Lexer) peekChar() byte {
    if l.readPosition >= len(l.input) {
        return 0
    } 

    return l.input[l.readPosition]
}


