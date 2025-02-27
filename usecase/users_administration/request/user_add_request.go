package request

import (
	"github.com/labstack/echo/v4"
	"github.com/private-project-pp/pos-general-lib/stacktrace"
)

type UserAddRequest struct {
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
}

func NewUserAdd(in echo.Context) (out UserAddRequest, err error) {
	if err = in.Bind(&out); err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error())
	}

	return out, nil
}
