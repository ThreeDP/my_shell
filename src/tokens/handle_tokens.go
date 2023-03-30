package tokens

type TypeToken int8

/* Types of operators handled in the shell */
const (
	OpDefault TypeToken = iota
	OpCmd TypeToken = iota
	OpBuildIn
	OpInput
	OpUntil
	OpOutput
	OpAppend
	OpAnd
	OpOr
	OpPipe
)

type Token struct {
	Cmds	[]string
	Type	TypeToken
	Next	*Token
}

func (token *Token) CreateToken(cmds []string, tt TypeToken) {
	token.Cmds = cmds
	token.Type = tt
}