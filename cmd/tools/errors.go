package tools

import (
	"fangkm.demo/cmd/model"
	"github.com/pkg/errors"
	"net/http"
)

func CatchError(w http.ResponseWriter, r *http.Request) {
	if err := recover(); err != nil {
		resp := model.Resp{
			Code:    "9001",
			Message: err.(error).Error(),
		}

		MarshalJson(w, &resp)
	}
}

func ThrowError(err error) {
	if err != nil {
		panic(errors.WithStack(err))
	}
}
