package main

import (
	"errors"
	"os"
	"os/exec"
)

func RunCmd(cmd []string, env Environment) (returnCode int) {
	prepareEnv(env)
	// #nosec G204
	command := exec.Command(cmd[0], cmd[1:]...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		return exitErr.ExitCode()
	}
	if err != nil {
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
