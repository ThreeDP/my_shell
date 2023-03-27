package env

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestReloadEnv(t *testing.T) {
	t.Run("return the copy of the env", func(t *testing.T){
		expected := os.Environ()
		s := initSys(expected)
		if !reflect.DeepEqual(expected, s.env) {
			t.Fatalf("expected %v, result %v", expected, s.env)
		}
	})
}

func TestKeySearch(t *testing.T) {
	t.Run("fetches the first item of the env", func(t *testing.T) {
		i := 0
		s := initSys(os.Environ())
		evar := os.Environ()[i]
		key := evar[:strings.IndexByte(evar, '=')]
		s.compareResult(t, i, true, key)
	})

	t.Run("fetches the last item from the env", func(t *testing.T) {
		i := len(os.Environ()) -1
		s := initSys(os.Environ())
		evar := os.Environ()[i]
		key := evar[:strings.IndexByte(evar, '=')]
		s.compareResult(t, i, true, key)
	})

	t.Run("looking for a key that doesn't exist", func(t *testing.T) {
		s := initSys(os.Environ())
		key := "YOURSHADOWNOTPASS"
		s.compareResult(t, -1, false, key)
	})
}

func TestUnset(t *testing.T) {

	t.Run("unset a middle variable", func(t *testing.T){
		key := "STATUS"
		env := []string{"USER=dapaulin", "KEY=1997", "STATUS=www.google.com", "TEST=value"}
		expected := []string{"USER=dapaulin", "KEY=1997", "TEST=value"}
		assertUnset(t, key, expected, env)
	})

	t.Run("unset the last variable", func(t *testing.T){
		key := "TEST"
		env := []string{"USER=dapaulin", "KEY=1997", "STATUS=www.google.com", "TEST=value"}
		expected := []string{"USER=dapaulin", "KEY=1997", "STATUS=www.google.com"}
		assertUnset(t, key, expected, env)
	})

	t.Run("unset the first variable", func(t *testing.T){
		key := "USER"
		env := []string{"USER=dapaulin", "KEY=1997", "STATUS=www.google.com", "TEST=value"}
		expected := []string{"KEY=1997", "STATUS=www.google.com", "TEST=value"}
		assertUnset(t, key, expected, env)
	})

	t.Run("unset a non-existent variable", func(t *testing.T){
		key := "IRMAO_DO_JOREL"
		env := []string{"USER=dapaulin", "KEY=1997", "STATUS=www.google.com", "TEST=value"}
		expected := []string{"USER=dapaulin", "KEY=1997", "STATUS=www.google.com", "TEST=value"}
		assertUnset(t, key, expected, env)
	})
}

func assertUnset(t *testing.T, key string, expected, env []string) {
	t.Helper()
	s := initSys(env)
	s.Unset(key)
	if !reflect.DeepEqual(expected, s.env) {
		t.Fatalf("expected %v, result %v", expected, s.env)
	}
}

func initSys(env []string) SysConfig {
	s := SysConfig{}
	s.CreateEnv(env)
	return s
}

func (s SysConfig) compareResult(t *testing.T, exp_i int, exp bool, key string) {
	t.Helper()
	result, i := s.KeySearch(key)
	if exp != result {
		t.Error("return expected is true but return false")
	}
	if exp_i != i {
		t.Errorf("expected Index '%d', result index '%d'", exp_i, i)
	}
}
