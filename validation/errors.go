package validation

import "fmt"

type ValidationError struct {
	ValidationRule string
	Error          error
}

type ValidationErrors []ValidationError

type ParsingError struct {
	RuleText string
	Error    error
}

func NewValidationError(rule string, err error) *ValidationError {
	return &ValidationError{
		ValidationRule: rule,
		Error:          err,
	}
}

func NewParsingError(rule string, err error) *ParsingError {
	return &ParsingError{
		RuleText: rule,
		Error:    err,
	}
}

func (err *ParsingError) String() string {
	return fmt.Sprintf("error parsing rule '%s' with error '%s'", err.RuleText, err.Error.Error())
}
