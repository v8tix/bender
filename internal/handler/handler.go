package handler

import (
	. "github.com/v8tix/kit/handler"
)

type BenderHandler struct {
	GenericHandler
}

func NewBenderHandler(genericHandler GenericHandler) BenderHandler {
	handler := BenderHandler{
		genericHandler,
	}
	return handler
}
