package main_test

import (
	"github.com/gkarman/otus_go_home_work/hw08_envdir_tool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadDir(t *testing.T) {
	t.Run("dir no exit", func(t *testing.T) {
		dir := "/dir_no_exit"
		_, err := main.ReadDir(dir)
		require.Error(t, err)
	})

	t.Run("valid values", func(t *testing.T) {
		dir := "./testdata/env"
		expected := main.Environment{
			"BAR":   main.EnvValue{Value: "bar", NeedRemove: false},
			"EMPTY": main.EnvValue{Value: "", NeedRemove: false},
			"FOO":   main.EnvValue{Value: "   foo\nwith new line", NeedRemove: false},
			"HELLO": main.EnvValue{Value: "\"hello\"", NeedRemove: false},
			"UNSET": main.EnvValue{Value: "", NeedRemove: true},
		}

		envResult, err := main.ReadDir(dir)
		if len(envResult) != len(expected) {
			t.Errorf("Expected %d variables, got %d", len(expected), len(envResult))
		}

		require.NoError(t, err)

		for key, expectedVal := range expected {
			actualVal, exists := envResult[key]
			if !exists {
				t.Errorf("Expected variable %s not found", key)
				continue
			}
			assert.Equal(t, expectedVal.Value, actualVal.Value)
		}
	})
}
