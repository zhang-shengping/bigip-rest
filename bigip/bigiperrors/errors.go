package bigiperrors

import (
	"fmt"
	"net/http"
)

type ResponseError struct {
	Resp *http.Response
}

func (e ResponseError) Error() string {
	return fmt.Sprintf(
		"Http Error, Method is: %s,\n Status is: %s,\n Body is %s\n",
		e.Resp.Request.Method, e.Resp.Status, e.Resp.Body,
	)
}

type ServiceError struct {
	ResourceError string
	HttpError     error
}

func (e ServiceError) Error() string {
	return fmt.Sprintf(
		"Resource Error! Information: %s\n, It causes by %s",
		e.ResourceError, e.HttpError,
	)
}
