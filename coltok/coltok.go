package coltok

// TokenType : type of token
type TokenType int

// valid token types in COLON
const (
	INT TokenType = iota // INTEGER
	FLT                  // FLOAT
	STR                  // STRING
	BOL                  // BOOLEAN

	LPR // LEFT PARENTHESIS
	RPR // RIGHT PARENTHESIS

	PLS // PLUS
	MIN // MINUS
	PRD // PRODUCT
	DIV // DIVISION
	REM // REMAINDER
	POW // POWER

	EQL // EQUAL
	NEQ // NOT EQUAL
	GRT // GREATER
	LST // LESSER
	GRE // GEATER OR EQUAL
	LSE // LESS OR EQUAL

	ASN // ASSIGNMENT
	COM // COMMA

	LND // LOGICAL_AND
	LOR // LOGICAL_OR
	LNT // LOGICAL_NOT

	IDN // IDENTIFIER

	PRN // PRINT
	INP // INPUT
	VAR // VARIABLE
	SWB // SWITCH BEGIN
	SWE // SWITCH END
	SLP // SLEEP
	LPB // LOOP BEGIN
	LPE // LOOP END
	FNB // FUNCTION BEGIN
	FNE // FUNCTION END
	RET // RETURN
	BRK // BREAK
	CNT // CONTINUE

	EOL // END OF LINE
	EOF // END OF FILE
	ILG // ILLEGAL CHARACTER
)

// String [from tokentype]
func (t TokenType) String() string {
	switch t {
	case INT:
		return "INTEGER"
	case FLT:
		return "FLOATING"
	case STR:
		return "STRING"
	case BOL:
		return "BOOLEAN"
	case LPR:
		return "LEFT_PARENTHESES"
	case RPR:
		return "RIGHT_PARENTHESES"
	case PLS:
		return "PLUS"
	case MIN:
		return "MINUS"
	case PRD:
		return "PRODUCT"
	case DIV:
		return "DIVISION"
	case REM:
		return "REMAINDER"
	case POW:
		return "POWER"
	case EQL:
		return "EQUAL"
	case NEQ:
		return "NOT_EQUAL"
	case GRT:
		return "GREATER_THAN"
	case LST:
		return "LESS_THAN"
	case GRE:
		return "GREATER_EQUAL"
	case LSE:
		return "LESSER_EQUAL"
	case ASN:
		return "ASSIGNMENT"
	case LND:
		return "LOGICAL_AND"
	case LOR:
		return "LOGICAL_OR"
	case LNT:
		return "LOGICAL_NOT"
	case IDN:
		return "IDENTIFIER"
	case PRN:
		return "PRINT"
	case INP:
		return "INPUT"
	case VAR:
		return "VARIABLE"
	case SWB:
		return "SWB"
	case SWE:
		return "SWE"
	case SLP:
		return "SLP"
	case LPB:
		return "LPB"
	case LPE:
		return "LPE"
	case FNB:
		return "FNB"
	case FNE:
		return "FNE"
	case RET:
		return "RET"
	case BRK:
		return "BRK"
	case CNT:
		return "CNT"
	case COM:
		return "COM"
	case EOF:
		return "EOF"
	case EOL:
		return "EOL"
	}
	return "ILLEGAL_TOKEN"
}

// Keywords map
var Keywords map[string]TokenType = map[string]TokenType{
	"p:": PRN,
	"i:": INP,
	"v:": VAR,
	"s:": SWB,
	":s": SWE,
	"w:": SLP,
	"l:": LPB,
	":l": LPE,
	"f:": FNB,
	":f": FNE,
	"r:": RET,
	"b:": BRK,
	"c:": CNT,
}

// Token : properties
type Token struct {
	TokType TokenType
	Literal string
	Line    int
}

// NewToken : to assemble a token 'object'
func NewToken(tokenType TokenType, lit string, line int) Token {
	return Token{TokType: tokenType, Literal: lit, Line: line}
}

// IsDigit : to check if the current character is a digit
func IsDigit(val byte) bool {
	if val >= '0' && val <= '9' {
		return true
	}
	return false
}

// IsLetter : to check if the current character is a letter
func IsLetter(val byte) bool {
	if (val >= 'a' && val <= 'z') || (val >= 'A' && val <= 'Z') || (val == '_') || (val == ':') {
		return true
	}
	return false
}

// IsKeyword : to check if the given string is a keyword or not
func IsKeyword(word string) bool {
	if _, ok := Keywords[word]; ok {
		return true
	}
	return false
}