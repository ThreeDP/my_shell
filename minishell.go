package main

import (
	"fmt"
	"bufio"
	"strings"
	"os/exec"
	"os"
	"errors"
	"./src/environ"
)

func execInput(input string, s *env.SysConfig) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}
		return os.Chdir(args[1])
	case "unset":
		s.Unset(args[1])
	case "export":
		s.Export(args[1])
	case "env":
		s.PrintEnv();
	case "exit":
		os.Exit(0)
	}
	cmd := exec.Command(args[0], args[1:]...)
	//cmd.Env = s.env
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := env.SysConfig{}
	s.CreateEnv(os.Environ())
	for {
		fmt.Print("3DPShell> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err = execInput(input, &s); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
