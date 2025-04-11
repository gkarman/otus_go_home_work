package main

import (
	"os"
	"os/exec"
)

func RunCmd(cmd []string, env Environment) (returnCode int) {
	prepareEnv(env)
	command := exec.Command(cmd[0], cmd[1:]...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return exitErr.ExitCode()
		}
		return 1
	}

	return 0
}

func prepareEnv(envFromDir Environment) {
	for keyEnv, itemEnv := range envFromDir {
		_ = os.Unsetenv(keyEnv)
		if itemEnv.NeedRemove {
			continue
		}
		_ = os.Setenv(keyEnv, itemEnv.Value)
	}
}
