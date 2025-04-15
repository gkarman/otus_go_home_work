package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("dir no exit", func(t *testing.T) {
		dir := "/dir_no_exit"
		_, err := ReadDir(dir)
		require.Error(t, err)
	})

	t.Run("valid values", func(t *testing.T) {
		dir := "./testdata/env"
		expected := Environment{
			"BAR":   EnvValue{Value: "bar", NeedRemove: false},
			"EMPTY": EnvValue{Value: "", NeedRemove: false},
			"FOO":   EnvValue{Value: "   foo\nwith new line", NeedRemove: false},
			"HELLO": EnvValue{Value: "\"hello\"", NeedRemove: false},
			"UNSET": EnvValue{Value: "", NeedRemove: true},
		}

		envResult, err := ReadDir(dir)
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
