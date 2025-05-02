package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var ErrMustBeStruct = errors.New("только для структуры")
var ErrInvalidTag = errors.New("тег невалидный")
var ErrValidationRuleDoesntSupportTypeField = errors.New("тип не поддерживается")
var ErrValueTooSmall = errors.New("значение меньше минимального")
var ErrValueTooBig = errors.New("значение выше максимального")
var ErrLenString = errors.New("строка не той длинны")
var ErrRegexpString = errors.New("значение не нужного формата")
var ErrValueNotInList = errors.New("значение не входит в допустимый список")

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

type ValidateFunc func(ValidationFieldData) (ValidationErrors, error)

type ValidationFieldData struct {
	Field     reflect.StructField
	Value     reflect.Value
	ruleValue string
}

func (v ValidationErrors) Error() string {
	if len(v) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString("Ошибки валидации:\n")

	for _, err := range v {
		sb.WriteString(fmt.Sprintf(" - Поле '%s': %s\n", err.Field, err.Err))
	}

	return sb.String()
}

func Validate(structForValidate interface{}) error {
	structType := reflect.TypeOf(structForValidate)
	if structType.Kind() != reflect.Struct {
		return ErrMustBeStruct
	}

	var allValidationErrors ValidationErrors
	structValues := reflect.ValueOf(structForValidate)
	for i := 0; i < structType.NumField(); i++ {
		structField := structType.Field(i)
		valueField := structValues.Field(i)

		fieldValidationErrors, err := processValidateField(structField, valueField)
		if err != nil {
			return ErrInvalidTag
		}
		if fieldValidationErrors != nil {
			allValidationErrors = append(allValidationErrors, fieldValidationErrors...)
		}
	}

	if allValidationErrors == nil {
		return nil
	}

	return allValidationErrors
}

func processValidateField(structField reflect.StructField, valueField reflect.Value) (ValidationErrors, error) {
	if !structField.IsExported() {
		return nil, nil
	}

	validateTag := structField.Tag.Get("validate")
	if validateTag == "" {
		return nil, nil
	}

	var fieldValidationErrors ValidationErrors

	validateRules := strings.Split(validateTag, "|")
	for _, validateRule := range validateRules {
		validateRuleParts := strings.Split(validateRule, ":")
		if len(validateRuleParts) != 2 {
			return nil, ErrInvalidTag
		}
		validateRuleName := validateRuleParts[0]
		validateRuleValue := validateRuleParts[1]

		functionForValidate, err := getFunctionForValidateByRuleName(validateRuleName)
		if err != nil {
			return nil, fmt.Errorf("Тег валидации не поддерживается %w", err)
		}

		validationFieldData := ValidationFieldData{structField, valueField, validateRuleValue}
		return functionForValidate(validationFieldData)

	}
	return fieldValidationErrors, nil

}

func getFunctionForValidateByRuleName(ruleName string) (ValidateFunc, error) {
	validateFunctions := map[string]ValidateFunc{
		"len":    processValidateLen,
		"min":    processValidateMin,
		"max":    processValidateMax,
		"regexp": processValidateRegexp,
		"in":     processValidateIn,
	}

	function, ok := validateFunctions[ruleName]
	if !ok {
		return nil, fmt.Errorf("<UNK> <UNK> <UNK> %s", ruleName)
	}

	return function, nil
}

func processValidateLen(vfd ValidationFieldData) (ValidationErrors, error) {
	var fieldValidationErrors ValidationErrors

	ruleConstraint, err := strconv.Atoi(vfd.ruleValue)
	if err != nil {
		return fieldValidationErrors, ErrInvalidTag
	}

	isOneValue := vfd.Value.Kind() == reflect.String
	isSlice := vfd.Field.Type.Kind() == reflect.Slice && vfd.Field.Type.Elem().Kind() == reflect.String

	if !isOneValue && !isSlice {
		return fieldValidationErrors, ErrValidationRuleDoesntSupportTypeField
	}

	itemsForValidate := []string{}
	if isOneValue {
		itemsForValidate = append(itemsForValidate, vfd.Value.String())
	}

	if isSlice {
		for i := 0; i < vfd.Value.Len(); i++ {
			itemsForValidate = append(itemsForValidate, vfd.Value.Index(i).String())
		}
	}

	for i := 0; i < len(itemsForValidate); i++ {
		if err := checkLenString(itemsForValidate[i], ruleConstraint); err != nil {
			fieldValidationErrors = append(fieldValidationErrors, ValidationError{vfd.Field.Name, err})
		}
	}

	return fieldValidationErrors, nil
}

func processValidateMin(vfd ValidationFieldData) (ValidationErrors, error) {
	var fieldValidationErrors ValidationErrors

	ruleConstraint, err := strconv.Atoi(vfd.ruleValue)
	if err != nil {
		return fieldValidationErrors, ErrInvalidTag
	}

	isOneValue := vfd.Value.Kind() == reflect.Int
	isSlice := vfd.Field.Type.Kind() == reflect.Slice && vfd.Field.Type.Elem().Kind() == reflect.Int

	if !isOneValue && !isSlice {
		return fieldValidationErrors, ErrValidationRuleDoesntSupportTypeField
	}

	itemsForValidate := []int{}
	if isOneValue {
		itemsForValidate = append(itemsForValidate, int(vfd.Value.Int()))
	}

	if isSlice {
		for i := 0; i < vfd.Value.Len(); i++ {
			itemsForValidate = append(itemsForValidate, int(vfd.Value.Index(i).Int()))
		}
	}

	for i := 0; i < len(itemsForValidate); i++ {
		if err := checkMinInt(itemsForValidate[i], ruleConstraint); err != nil {
			fieldValidationErrors = append(fieldValidationErrors, ValidationError{vfd.Field.Name, err})
		}
	}

	return fieldValidationErrors, nil
}

func processValidateMax(vfd ValidationFieldData) (ValidationErrors, error) {
	var fieldValidationErrors ValidationErrors

	ruleConstraint, err := strconv.Atoi(vfd.ruleValue)
	if err != nil {
		return fieldValidationErrors, ErrInvalidTag
	}

	isOneValue := vfd.Value.Kind() == reflect.Int
	isSlice := vfd.Field.Type.Kind() == reflect.Slice && vfd.Field.Type.Elem().Kind() == reflect.Int

	if !isOneValue && !isSlice {
		return fieldValidationErrors, ErrValidationRuleDoesntSupportTypeField
	}

	itemsForValidate := []int{}
	if isOneValue {
		itemsForValidate = append(itemsForValidate, int(vfd.Value.Int()))
	}

	if isSlice {
		for i := 0; i < vfd.Value.Len(); i++ {
			itemsForValidate = append(itemsForValidate, int(vfd.Value.Index(i).Int()))
		}
	}

	for i := 0; i < len(itemsForValidate); i++ {
		if err := checkMaxInt(itemsForValidate[i], ruleConstraint); err != nil {
			fieldValidationErrors = append(fieldValidationErrors, ValidationError{vfd.Field.Name, err})
		}
	}

	return fieldValidationErrors, nil
}

func processValidateRegexp(vfd ValidationFieldData) (ValidationErrors, error) {
	var fieldValidationErrors ValidationErrors

	ruleConstraint, err := regexp.Compile(vfd.ruleValue)
	if err != nil {
		return fieldValidationErrors, ErrInvalidTag
	}

	isOneValue := vfd.Value.Kind() == reflect.String
	isSlice := vfd.Field.Type.Kind() == reflect.Slice && vfd.Field.Type.Elem().Kind() == reflect.String

	if !isOneValue && !isSlice {
		return fieldValidationErrors, ErrValidationRuleDoesntSupportTypeField
	}

	itemsForValidate := []string{}
	if isOneValue {
		itemsForValidate = append(itemsForValidate, vfd.Value.String())
	}

	if isSlice {
		for i := 0; i < vfd.Value.Len(); i++ {
			itemsForValidate = append(itemsForValidate, vfd.Value.Index(i).String())
		}
	}

	for i := 0; i < len(itemsForValidate); i++ {
		if err := checkRegexp(itemsForValidate[i], ruleConstraint); err != nil {
			fieldValidationErrors = append(fieldValidationErrors, ValidationError{vfd.Field.Name, err})
		}
	}

	return fieldValidationErrors, nil
}

func processValidateIn(vfd ValidationFieldData) (ValidationErrors, error) {
	var fieldValidationErrors ValidationErrors
	ruleConstraints := strings.Split(vfd.ruleValue, ",")

	isOneValueString := vfd.Value.Kind() == reflect.String
	isSliceStrings := vfd.Field.Type.Kind() == reflect.Slice && vfd.Field.Type.Elem().Kind() == reflect.String
	isOneValueInt := vfd.Value.Kind() == reflect.Int
	isSliceInts := vfd.Field.Type.Kind() == reflect.Slice && vfd.Field.Type.Elem().Kind() == reflect.Int

	if !isOneValueString && !isSliceStrings && !isOneValueInt && !isSliceInts {
		return fieldValidationErrors, ErrValidationRuleDoesntSupportTypeField
	}

	if isOneValueString || isSliceStrings {
		itemsForValidate := []string{}
		if isOneValueString {
			itemsForValidate = append(itemsForValidate, vfd.Value.String())
		}

		if isSliceStrings {
			for i := 0; i < vfd.Value.Len(); i++ {
				itemsForValidate = append(itemsForValidate, vfd.Value.Index(i).String())
			}
		}

		for i := 0; i < len(itemsForValidate); i++ {
			if err := checkIn(itemsForValidate[i], ruleConstraints); err != nil {
				fieldValidationErrors = append(fieldValidationErrors, ValidationError{vfd.Field.Name, err})
			}
		}
	}

	if isOneValueInt || isSliceInts {
		itemsForValidateInt := []int{}
		if isOneValueInt {
			itemsForValidateInt = append(itemsForValidateInt, int(vfd.Value.Int()))
		}

		if isSliceInts {
			for i := 0; i < vfd.Value.Len(); i++ {
				itemsForValidateInt = append(itemsForValidateInt, int(vfd.Value.Index(i).Int()))
			}
		}
		ruleConstraintsInt := []int{}
		for _, ruleConstraintString := range ruleConstraints {
			ruleConstraintInt, err := strconv.Atoi(ruleConstraintString)
			if err != nil {
				return fieldValidationErrors, ErrInvalidTag
			}
			ruleConstraintsInt = append(ruleConstraintsInt, ruleConstraintInt)
		}

		for i := 0; i < len(itemsForValidateInt); i++ {
			if err := checkIn(itemsForValidateInt[i], ruleConstraintsInt); err != nil {
				fieldValidationErrors = append(fieldValidationErrors, ValidationError{vfd.Field.Name, err})
			}
		}
	}

	return fieldValidationErrors, nil
}

func checkLenString(s string, maxLen int) error {
	if utf8.RuneCountInString(s) != maxLen {
		return ErrLenString
	}
	return nil
}

func checkMinInt(v int, vMin int) error {
	if v < vMin {
		return ErrValueTooSmall
	}
	return nil
}

func checkMaxInt(v int, vMax int) error {
	if v > vMax {
		return ErrValueTooBig
	}
	return nil
}

func checkRegexp(v string, regexp *regexp.Regexp) error {
	if !regexp.MatchString(v) {
		return ErrRegexpString
	}
	return nil
}

func checkIn[T interface{ string | int }](v T, ruleConstraints []T) error {
	for _, ruleConstraint := range ruleConstraints {
		if v == ruleConstraint {
			return nil
		}
	}
	return ErrValueNotInList
}
