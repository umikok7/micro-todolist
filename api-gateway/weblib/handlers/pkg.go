package handlers

import (
	"api-gateway/pkg/logging"
	"errors"
)

// PanicIfUserError 包装错误
func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("userService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}

func PanicIfTaskError(err error) {
	if err != nil {
		err = errors.New("TaskService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}
