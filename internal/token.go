package internal

type TokenType int

const (
	TokenTypeVariable TokenType = iota + 1
	TokenTypeFunction
)

type Token interface {
	GetType() TokenType
	GetName() string
	GetDefinition() interface{}
	Print()
}
