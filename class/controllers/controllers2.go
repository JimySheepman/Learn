package controllers

import (
	"fmt"

	"github.com/JimySheepman/class/dto"
)

type asyncAdminServices interface {
	AsyncRead(req *dto.RestRequest2) error
	AsyncWrite(req *dto.RestRequest2) (*dto.RestResponse2, error)
}

type asyncAdminController struct {
	asyncAdminServices asyncAdminServices
}

func NewAsyncAdminControllers(asyncAdminServices asyncAdminServices) *asyncAdminController {
	return &asyncAdminController{
		asyncAdminServices: asyncAdminServices,
	}
}

func (c *asyncAdminController) ReadHandler2(req *dto.RestRequest2) (*dto.RestResponse2, error) {
	err := c.asyncAdminServices.AsyncRead(req)
	if err != nil {
		return nil, err // custom error + log
	}

	resp, err := c.asyncAdminServices.AsyncWrite(req)
	if err != nil {
		return nil, err
	}

	fmt.Println("ReadHandler2 run...")

	return resp, nil
}
