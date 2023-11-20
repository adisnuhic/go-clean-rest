package apperror

// ValidationError is custom defined validation error
type ValidationError struct {
	Tag     string `json:"tag"`
	Message string `json:"message"`
	Field   string `json:"field"`
	Param   string `json:"param"`
}

// NewValidationError returns new validation error object
func NewValidationError(tag string, msg string, field string, param string) *ValidationError {
	return &ValidationError{
		Tag:     tag,
		Message: msg,
		Field:   field,
		Param:   param,
	}
}
