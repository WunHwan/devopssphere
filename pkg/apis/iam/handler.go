package iam

import (
	"github.com/emicklei/go-restful"
	"io.github/devopssphere/pkg/models/iam/im"
)

type iamHandler struct {
	im *im.ManagementInterface
}

func NewIAMHandler(am *im.ManagementInterface) *iamHandler {
	return &iamHandler{
		im: am,
	}
}

func (h *iamHandler) CreateUser(req *restful.Request, resp *restful.Response) {

}
