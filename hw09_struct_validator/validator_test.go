package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:1"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

//nolint:funlen,gocognit // функция слишком длинная, но намеренно
func TestValidate(t *testing.T) {
	tests := []struct {
		structForValidate interface{}
		expectedErr       error
	}{
		{
			"somethingString",
			ErrMustBeStruct,
		},
		{
			User{
				"1",
				"Ivan",
				18,
				"ivan@mail.ru",
				"admin",
				[]string{"79111111111", "79222222222"},
				[]byte{1},
			},
			nil,
		},
		{
			User{
				"qwertyuiopoiuytrewqwertyuioiuytre",
				"Ivan",
				18,
				"ivan@mail.ru",
				"admin",
				[]string{"79111111111", "79222222222"},
				[]byte{1},
			},
			ErrLenString,
		},
		{
			User{
				"1",
				"Ivan",
				17,
				"ivan@mail.ru",
				"admin",
				[]string{"79111111111", "79222222222"},
				[]byte{1},
			},
			ErrValueTooSmall,
		},
		{
			User{
				"1",
				"Ivan",
				90,
				"ivan@mail.ru",
				"admin",
				[]string{"79111111111", "79222222222"},
				[]byte{1},
			},
			ErrValueTooBig,
		},
		{
			User{
				"1",
				"Ivan",
				18,
				"noemail",
				"admin",
				[]string{"79111111111", "79222222222"},
				[]byte{1},
			},
			ErrRegexpString,
		},
		{
			User{
				"1",
				"Ivan",
				18,
				"ivan@mail.ru",
				"admin11",
				[]string{"79111111111", "79222222222"},
				[]byte{1},
			},
			ErrValueNotInList,
		},
		{
			User{
				"1",
				"Ivan",
				18,
				"ivan@mail.ru",
				"admin",
				[]string{"79", "79"},
				[]byte{1},
			},
			ErrLenString,
		},
		{
			App{
				"12345",
			},
			nil,
		},
		{
			App{
				"1234",
			},
			ErrLenString,
		},
		{
			Token{
				[]byte{1, 2, 3},
				[]byte{1, 2, 3},
				[]byte{1, 2, 3},
			},
			nil,
		},
		{
			Response{
				200,
				"somestring",
			},
			nil,
		},
		{
			Response{
				211,
				"somestring",
			},
			ErrValueNotInList,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			err := Validate(tt.structForValidate)
			if tt.expectedErr == nil {
				if err != nil {
					t.Errorf("ожидалась ошибка <nil>, но получено: %v", err)
				}
				return
			}

			if errors.Is(err, tt.expectedErr) {
				return
			}
			if tt.expectedErr == nil && err != nil {
				t.Errorf("получили ошибку, хотя не должны были %v: ", err)
			}

			var valErrs ValidationErrors
			if errors.As(err, &valErrs) {
				found := false
				for _, vErr := range valErrs {
					if errors.Is(vErr.Err, tt.expectedErr) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("ожидалась ошибка %v, но в списке валидации её нет: %v", tt.expectedErr, err)
				}
				return
			}
			t.Errorf("ожидалась ошибка %v, но получено: %v", tt.expectedErr, err)
		})
	}
}
