package token 

type Token struct {
    Type uint8
    StringLiteral string
}

const (
    ILLEGAL = iota
    _
    EOF

    // Variables/Contants
    IDENTIFIER 
    DIGIT 

    // Operators
    ASSIGN 
    PLUS 
    MINUS
    BANG
    ASTERISK
    FORWARDSLASH
    LT
    GT

    // Separators
    COMMA 
    SEMICOLON 

    // Brackets
    LPAREN 
    RPAREN 
    LBRACE 
    RBRACE 

    // Multiword keywords
    FUNCTION 
    LET 
    TRUE
    FALSE
    IF
    ELSE
    RETURN

    EQUALS
    NOTEQUALS
)

var MultiWordKeywords = map[string]uint8 {
    "fn": FUNCTION,
    "let": LET,
    "true": TRUE,
    "false": FALSE,
    "if": IF,
    "else": ELSE,
    "return": RETURN,
}

var SingleWordKeywords = map[byte]uint8 {
    // Operators
    '=': ASSIGN,
    '+': PLUS,
    '-': MINUS,
    '!': BANG,
    '*': ASTERISK,
    '/': FORWARDSLASH,
    '<': LT,
    '>': GT,

    // Separators
    ';': SEMICOLON,
    ',': COMMA,

    // Brackets
    '(': LPAREN,
    ')': RPAREN,
    '{': LBRACE,
    '}': RBRACE,
}

func LookupIdent(ident string) uint8 { 
    if tok, ok := MultiWordKeywords[ident]; ok {
        return tok 
    }
    return IDENTIFIER 
}

