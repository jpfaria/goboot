package echo

import (
	"net/http"

	"github.com/b2wdigital/goignite/pkg/transport/server/http/rest"
	"github.com/go-playground/validator/v10"
	"github.com/juju/errors"
	"github.com/labstack/echo/v4"
)

func JSONResponse(c echo.Context, code int, i interface{}, err error) error {

	if err != nil {

		return JSONErrorResponse(c, err)

	}

	return c.JSONPretty(code, i, "  ")
}

func JSONErrorResponse(c echo.Context, err error) error {

	if errors.IsNotFound(err) {
		return c.JSONPretty(
			http.StatusNotFound,
			rest.ErrorResponse{HttpStatusCode: http.StatusNotFound, Message: err.Error()},
			"  ")
	} else if errors.IsNotValid(err) {

		return c.JSONPretty(
			http.StatusBadRequest,
			rest.ErrorResponse{HttpStatusCode: http.StatusBadRequest, Message: err.Error()},
			"  ")
	} else {

		switch t := err.(type) {
		default:
			return c.JSONPretty(
				http.StatusInternalServerError,
				rest.ErrorResponse{HttpStatusCode: http.StatusInternalServerError, Message: t.Error()},
				"  ")
		case validator.ValidationErrors:
			return c.JSONPretty(
				http.StatusUnprocessableEntity,
				rest.NewUnprocessableEntityErrorResponse(t),
				"  ")
		}
	}

}
