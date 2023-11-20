package viewmodels

// ResponseBase -
type ResponseBase struct {
	Success         bool        `json:"success"`
	RequestID       string      `json:"request_id"`
	Data            interface{} `json:"data"`
	Error           interface{} `json:"error"`
	ValidationError interface{} `json:"validation_error"`
}
