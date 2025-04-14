package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRunCmd(t *testing.T) {
	t.Run("check success code", func(t *testing.T) {
		cmd := []string{"echo", "hello"}
		env := Environment{}

		code := RunCmd(cmd, env)

		require.Equal(t, code, 0)
	})

	t.Run("check error code", func(t *testing.T) {
		cmd := []string{"sh", "-c", "exit 42"}
		env := Environment{}

		code := RunCmd(cmd, env)

		require.Equal(t, code, 42)
	})

}
