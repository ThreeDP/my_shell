package tokens

import (
	"testing"
)

func TestTokenNew(t *testing.T) {
	exp_cmds := []string {"ls", "-l", "-a"}
	exp_type := OpCmd
	exp_next := nil
	t := Token{}
	t.CreateToken(exp_cmds, exp_type, exp_next)

}

func (t Token) checkValues(t *testing.T, cmds []string, token TypeToken, next *Token) {
	t.Helper()
	if 
}