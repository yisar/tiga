package token

import (
	"strconv"
)

const (
	_ Token = iota

	ILLEGAL
	EOF
	WHITESPACE
	COMMENT
	KEYWORD

	STRING
	BOOLEAN
	NULL
	NUMBER
	IDENTIFIER

	PLUS      // +
	MINUS     // -
	MULTIPLY  // *
	SLASH     // /
	REMAINDER // %

	AND                  // &
	OR                   // |
	EXCLUSIVE_OR         // ^
	SHIFT_LEFT           // <<
	SHIFT_RIGHT          // >>
	UNSIGNED_SHIFT_RIGHT // >>>
	AND_NOT              // &^

	ADD_ASSIGN       // +=
	SUBTRACT_ASSIGN  // -=
	MULTIPLY_ASSIGN  // *=
	QUOTIENT_ASSIGN  // /=
	REMAINDER_ASSIGN // %=

	AND_ASSIGN                  // &=
	OR_ASSIGN                   // |=
	EXCLUSIVE_OR_ASSIGN         // ^=
	SHIFT_LEFT_ASSIGN           // <<=
	SHIFT_RIGHT_ASSIGN          // >>=
	UNSIGNED_SHIFT_RIGHT_ASSIGN // >>>=
	AND_NOT_ASSIGN              // &^=

	LOGICAL_AND // &&
	LOGICAL_OR  // ||
	INCREMENT   // ++
	DECREMENT   // --

	EQUAL        // ==
	STRICT_EQUAL // ===
	LESS         // <
	GREATER      // >
	ASSIGN       // =
	NOT          // !

	BITWISE_NOT // ~

	NOT_EQUAL        // !=
	STRICT_NOT_EQUAL // !==
	LESS_OR_EQUAL    // <=
	GREATER_OR_EQUAL // >=

	LEFT_PARENTHESIS // (
	LEFT_BRACKET     // [
	LEFT_BRACE       // {
	COMMA            // ,
	PERIOD           // .

	RIGHT_PARENTHESIS // )
	RIGHT_BRACKET     // ]
	RIGHT_BRACE       // }
	SEMICOLON         // ;
	COLON             // :
	QUESTION_MARK     // ?
	BACKSLASH

	firstKeyword
	IF
	IN
	DO

	VAR
	FOR
	NEW
	TRY

	THIS
	ELSE
	CASE
	VOID
	WITH

	WHILE
	BREAK
	CATCH
	THROW

	RETURN
	TYPEOF
	DELETE
	SWITCH

	DEFAULT
	FINALLY

	FUNCTION
	CONTINUE
	DEBUGGER

	INSTANCEOF
	lastKeyword
)

var token2string = [...]string{
	ILLEGAL:                     "ILLEGAL",
	EOF:                         "EOF",
	WHITESPACE:                  "WHITESPACE",
	COMMENT:                     "COMMENT",
	KEYWORD:                     "KEYWORD",
	STRING:                      "STRING",
	BOOLEAN:                     "BOOLEAN",
	NULL:                        "NULL",
	NUMBER:                      "NUMBER",
	IDENTIFIER:                  "IDENTIFIER",
	PLUS:                        "+",
	MINUS:                       "-",
	MULTIPLY:                    "*",
	SLASH:                       "/",
	REMAINDER:                   "%",
	AND:                         "&",
	OR:                          "|",
	EXCLUSIVE_OR:                "^",
	SHIFT_LEFT:                  "<<",
	SHIFT_RIGHT:                 ">>",
	UNSIGNED_SHIFT_RIGHT:        ">>>",
	AND_NOT:                     "&^",
	ADD_ASSIGN:                  "+=",
	SUBTRACT_ASSIGN:             "-=",
	MULTIPLY_ASSIGN:             "*=",
	QUOTIENT_ASSIGN:             "/=",
	REMAINDER_ASSIGN:            "%=",
	AND_ASSIGN:                  "&=",
	OR_ASSIGN:                   "|=",
	EXCLUSIVE_OR_ASSIGN:         "^=",
	SHIFT_LEFT_ASSIGN:           "<<=",
	SHIFT_RIGHT_ASSIGN:          ">>=",
	UNSIGNED_SHIFT_RIGHT_ASSIGN: ">>>=",
	AND_NOT_ASSIGN:              "&^=",
	LOGICAL_AND:                 "&&",
	LOGICAL_OR:                  "||",
	INCREMENT:                   "++",
	DECREMENT:                   "--",
	EQUAL:                       "==",
	STRICT_EQUAL:                "===",
	LESS:                        "<",
	GREATER:                     ">",
	ASSIGN:                      "=",
	NOT:                         "!",
	BITWISE_NOT:                 "~",
	NOT_EQUAL:                   "!=",
	STRICT_NOT_EQUAL:            "!==",
	LESS_OR_EQUAL:               "<=",
	GREATER_OR_EQUAL:            ">=",
	LEFT_PARENTHESIS:            "(",
	LEFT_BRACKET:                "[",
	LEFT_BRACE:                  "{",
	COMMA:                       ",",
	PERIOD:                      ".",
	RIGHT_PARENTHESIS:           ")",
	RIGHT_BRACKET:               "]",
	RIGHT_BRACE:                 "}",
	SEMICOLON:                   ";",
	COLON:                       ":",
	QUESTION_MARK:               "?",
	BACKSLASH:                   "\\",
	IF:                          "if",
	IN:                          "in",
	DO:                          "do",
	VAR:                         "var",
	FOR:                         "for",
	NEW:                         "new",
	TRY:                         "try",
	THIS:                        "this",
	ELSE:                        "else",
	CASE:                        "case",
	VOID:                        "void",
	WITH:                        "with",
	WHILE:                       "while",
	BREAK:                       "break",
	CATCH:                       "catch",
	THROW:                       "throw",
	RETURN:                      "return",
	TYPEOF:                      "typeof",
	DELETE:                      "delete",
	SWITCH:                      "switch",
	DEFAULT:                     "default",
	FINALLY:                     "finally",
	FUNCTION:                    "function",
	CONTINUE:                    "continue",
	DEBUGGER:                    "debugger",
	INSTANCEOF:                  "instanceof",
}

var keywordTable = map[string]_keyword{
	"if": _keyword{
		token: IF,
	},
	"in": _keyword{
		token: IN,
	},
	"do": _keyword{
		token: DO,
	},
	"var": _keyword{
		token: VAR,
	},
	"for": _keyword{
		token: FOR,
	},
	"new": _keyword{
		token: NEW,
	},
	"try": _keyword{
		token: TRY,
	},
	"this": _keyword{
		token: THIS,
	},
	"else": _keyword{
		token: ELSE,
	},
	"case": _keyword{
		token: CASE,
	},
	"void": _keyword{
		token: VOID,
	},
	"with": _keyword{
		token: WITH,
	},
	"while": _keyword{
		token: WHILE,
	},
	"break": _keyword{
		token: BREAK,
	},
	"catch": _keyword{
		token: CATCH,
	},
	"throw": _keyword{
		token: THROW,
	},
	"return": _keyword{
		token: RETURN,
	},
	"typeof": _keyword{
		token: TYPEOF,
	},
	"delete": _keyword{
		token: DELETE,
	},
	"switch": _keyword{
		token: SWITCH,
	},
	"default": _keyword{
		token: DEFAULT,
	},
	"finally": _keyword{
		token: FINALLY,
	},
	"function": _keyword{
		token: FUNCTION,
	},
	"continue": _keyword{
		token: CONTINUE,
	},
	"debugger": _keyword{
		token: DEBUGGER,
	},
	"instanceof": _keyword{
		token: INSTANCEOF,
	},
	"const": _keyword{
		token:         KEYWORD,
		futureKeyword: true,
	},
	"class": _keyword{
		token:         KEYWORD,
		futureKeyword: true,
	},
	"enum": _keyword{
		token:         KEYWORD,
		futureKeyword: true,
	},
	"export": _keyword{
		token:         KEYWORD,
		futureKeyword: true,
	},
	"extends": _keyword{
		token:         KEYWORD,
		futureKeyword: true,
	},
	"import": _keyword{
		token:         KEYWORD,
		futureKeyword: true,
	},
	"super": _keyword{
		token:         KEYWORD,
		futureKeyword: true,
	},
	"implements": _keyword{
		token:         KEYWORD,
		futureKeyword: true,
		strict:        true,
	},
	"interface": _keyword{
		token:         KEYWORD,
		futureKeyword: true,
		strict:        true,
	},
	"let": _keyword{
		token:         KEYWORD,
		futureKeyword: true,
		strict:        true,
	},
	"package": _keyword{
		token:         KEYWORD,
		futureKeyword: true,
		strict:        true,
	},
	"private": _keyword{
		token:         KEYWORD,
		futureKeyword: true,
		strict:        true,
	},
	"protected": _keyword{
		token:         KEYWORD,
		futureKeyword: true,
		strict:        true,
	},
	"public": _keyword{
		token:         KEYWORD,
		futureKeyword: true,
		strict:        true,
	},
	"static": _keyword{
		token:         KEYWORD,
		futureKeyword: true,
		strict:        true,
	},
}

// Token is the set of lexical tokens in JavaScript (ECMA5).
type Token int

// String returns the string corresponding to the token.
// For operators, delimiters, and keywords the string is the actual
// token string (e.g., for the token PLUS, the String() is
// "+"). For all other tokens the string corresponds to the token
// name (e.g. for the token IDENTIFIER, the string is "IDENTIFIER").
//
func (tkn Token) String() string {
	if 0 == tkn {
		return "UNKNOWN"
	}
	if tkn < Token(len(token2string)) {
		return token2string[tkn]
	}
	return "token(" + strconv.Itoa(int(tkn)) + ")"
}

// This is not used for anything
func (tkn Token) precedence(in bool) int {

	switch tkn {
	case LOGICAL_OR:
		return 1

	case LOGICAL_AND:
		return 2

	case OR, OR_ASSIGN:
		return 3

	case EXCLUSIVE_OR:
		return 4

	case AND, AND_ASSIGN, AND_NOT, AND_NOT_ASSIGN:
		return 5

	case EQUAL,
		NOT_EQUAL,
		STRICT_EQUAL,
		STRICT_NOT_EQUAL:
		return 6

	case LESS, GREATER, LESS_OR_EQUAL, GREATER_OR_EQUAL, INSTANCEOF:
		return 7

	case IN:
		if in {
			return 7
		}
		return 0

	case SHIFT_LEFT, SHIFT_RIGHT, UNSIGNED_SHIFT_RIGHT:
		fallthrough
	case SHIFT_LEFT_ASSIGN, SHIFT_RIGHT_ASSIGN, UNSIGNED_SHIFT_RIGHT_ASSIGN:
		return 8

	case PLUS, MINUS, ADD_ASSIGN, SUBTRACT_ASSIGN:
		return 9

	case MULTIPLY, SLASH, REMAINDER, MULTIPLY_ASSIGN, QUOTIENT_ASSIGN, REMAINDER_ASSIGN:
		return 11
	}
	return 0
}

type _keyword struct {
	token         Token
	futureKeyword bool
	strict        bool
}

// IsKeyword returns the keyword token if literal is a keyword, a KEYWORD token
// if the literal is a future keyword (const, let, class, super, ...), or 0 if the literal is not a keyword.
//
// If the literal is a keyword, IsKeyword returns a second value indicating if the literal
// is considered a future keyword in strict-mode only.
//
// 7.6.1.2 Future Reserved Words:
//
//       const
//       class
//       enum
//       export
//       extends
//       import
//       super
//
// 7.6.1.2 Future Reserved Words (strict):
//
//       implements
//       interface
//       let
//       package
//       private
//       protected
//       public
//       static
//
func IsKeyword(literal string) (Token, bool) {
	if keyword, exists := keywordTable[literal]; exists {
		if keyword.futureKeyword {
			return KEYWORD, keyword.strict
		}
		return keyword.token, false
	}
	return 0, false
}
