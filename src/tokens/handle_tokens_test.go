package tokens

import (
	"testing"
	"reflect"
)

func TestCreateToken(t *testing.T) {

	t.Run("Test Create a node", func(t *testing.T) {
		expected_token := &Token{Cmds: []string{"ls", "-l", "-a"}, Type: OpCmd}
		token := &Token{}
		token.CreateToken([]string{"ls", "-l", "-a"}, OpCmd)
		checkValues(t, expected_token, token)
	})
}

func TestAddTokenBack(t *testing.T) {
	token := &Token{}
	expected_token := &Token{Cmds: []string{"ls", "-l", "-a"}, Type: OpCmd, Next: &Token{Type: OpPipe}}
	token.CreateToken([]string{"ls", "-l", "-a"}, OpCmd)
	token.AddTokenBack([]string{}, OpPipe)
	checkValues(t, expected_token, token)
	checkValues(t, expected_token.Next, token.Next)
}

func checkValues(t *testing.T, expected *Token, token *Token) {
	t.Helper()
	if !reflect.DeepEqual(expected, token) {
		t.Fatalf("expected %v, result %v", expected, token)
	}
}