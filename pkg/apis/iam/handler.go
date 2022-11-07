package iam

import (
	"github.com/emicklei/go-restful"
	"io.github/devopssphere/pkg/models/iam/am"
)

type iamHandler struct {
	am am.AccessManagementInterface
}

func NewIAMHandler(am am.AccessManagementInterface) *iamHandler {
	return &iamHandler{
		am: am,
	}
}

func (h *iamHandler) CreateUser(req *restful.Request, resp *restful.Response) {

}
