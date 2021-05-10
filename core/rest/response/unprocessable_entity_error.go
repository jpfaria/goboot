package response

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type UnprocessableEntityError struct {
	Error
	ValidationErrors []ValidationError `json:"validationErrors,omitempty"`
}

func NewUnprocessableEntity(err validator.ValidationErrors) UnprocessableEntityError {

	var fe validator.FieldError
	var verrs []ValidationError

	for i := 0; i < len(err); i++ {

		fe = err[i].(validator.FieldError)

		verr := ValidationError{
			FieldName: fe.Field(),
			Message:   "invalid value",
		}

		verrs = append(verrs, verr)
	}

	return UnprocessableEntityError{
		Error: Error{
			HttpStatusCode: http.StatusUnprocessableEntity,
			Message:        "The server understands the content type of the request entity but was unable to process the contained instructions.",
		},
		ValidationErrors: verrs,
	}
}
