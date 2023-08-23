package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marioheryanto/erajaya/go-app/model"
)

func NewServiceError(code int, message string) model.ServiceError {
	return model.ServiceError{
		Code:    code,
		Message: message,
	}
}

func GenerateResponse(c *gin.Context, r *model.Response, err error) (int, model.Response) {
	code := http.StatusInternalServerError

	//safety assert
	serviceErr, ok := err.(model.ServiceError)
	if !ok {
		r.Message = err.Error()
		return code, *r
	}

	if serviceErr.Code != 0 {
		code = serviceErr.Code
	}

	r.Message = serviceErr.Error()

	return code, *r
}
