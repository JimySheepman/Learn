package controllers

import (
	"fmt"

	"github.com/JimySheepman/class/dto"
)

type adminServices interface {
	Read(req *dto.RestRequest) error
	Write(req *dto.RestRequest) (*dto.RestResponse, error)
}

type adminController struct {
	adminServices adminServices
}

func NewAdminControllers(adminServices adminServices) *adminController {
	return &adminController{
		adminServices: adminServices,
	}
}

func (c *adminController) ReadHandler(req *dto.RestRequest) (*dto.RestResponse, error) {
	err := c.adminServices.Read(req)
	if err != nil {
		return nil, err // custom error + log
	}

	resp, err := c.adminServices.Write(req)
	if err != nil {
		return nil, err
	}

	fmt.Println("ReadHandler run...")

	return resp, nil
}
