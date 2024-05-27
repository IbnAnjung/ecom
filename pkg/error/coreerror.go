package coreerror

type CoreErrorType int

type CoreError struct {
	Type    CoreErrorType
	Message string
}

func NewCoreError(errorType CoreErrorType, message string) CoreError {
	if message == "" {
		message = "Internal Server Error"

		switch errorType {
		case CoreErrorTypeForbidden:
			message = "You don't hace any access"
		case CoreErrorTypeAuthorization:
			message = "Unauthorized"
		case CoreErrorTypeNotFound:
			message = "Data not found"
		case CoreErrorTypeUnprocessable:
			message = "Unproccessable entity"
		}
	}

	return CoreError{
		Type:    errorType,
		Message: message,
	}
}

const (
	CoreErrorTypeForbidden           CoreErrorType = 403
	CoreErrorTypeAuthorization       CoreErrorType = 401
	CoreErrorTypeNotFound            CoreErrorType = 404
	CoreErrorTypeUnprocessable       CoreErrorType = 400
	CoreErrorTypeInternalServerError CoreErrorType = 500
)

func (e CoreError) Error() string {
	return e.Message
}

// validation error
type ValidationError struct {
	Message string
	Errors  map[string]string
}

func NewValidationError() ValidationError {
	return ValidationError{
		Message: "Bad Request",
	}
}

func (e ValidationError) Error() string {
	return e.Message
}

func (e ValidationError) GetMessage() map[string]string {
	return e.Errors
}
