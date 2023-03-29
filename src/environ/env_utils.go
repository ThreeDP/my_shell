package env

import (
	"strings"
	"fmt"
)

type SysConfig struct {
	env		[]string
}

func (s *SysConfig) PrintEnv() {
	for i := 0; i < len(s.env); i++ {
		fmt.Println(s.env[i])
	}
}

func (s *SysConfig) Unset(key string) {
	if find, i := s.KeySearch(key); find == true {
		s.env = []string (append(s.env[:i], s.env[i+1:]...))
	}
}

func (s *SysConfig) Export(var_env string) {
	if find, i := s.KeySearch(var_env[:strings.IndexByte(var_env, '=')]); find == true {
		s.env[i] = var_env
		return
	}
	s.env = append(s.env, var_env)
}

func (s *SysConfig) KeySearch(key string) (bool, int) {
	for i := 0; i < len(s.env); i++ {
		if s.env[i][:strings.IndexByte(s.env[i], '=')] == key {
			return true, i
		}
	}
	return false, -1
}

func (s *SysConfig) CreateEnv(env []string) {
	s.env = env
}
