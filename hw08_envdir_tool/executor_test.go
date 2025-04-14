package main

import "testing"

func TestRunCmd(t *testing.T) {
	t.Run("check success code", func(t *testing.T) {
		cmd := []string{"echo", "hello"}
		env := Environment{}

		code := RunCmd(cmd, env)

		if code != 0 {
			t.Errorf("Expected exit code 0, got %d", code)
		}
	})

	t.Run("check error code", func(t *testing.T) {
		cmd := []string{"sh", "-c", "exit 42"}
		env := Environment{}

		code := RunCmd(cmd, env)

		if code != 42 {
			t.Errorf("Expected exit code 42, got %d", code)
		}
	})

}
