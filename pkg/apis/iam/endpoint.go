package iam

import (
	"github.com/emicklei/go-restful"
	"io.github/devopssphere/pkg/apiserver"
)

func AddToContainer(s *apiserver.APIServer) {
	handler := NewIAMHandler(nil)

	ws := new(restful.WebService)
	ws.Path("/iam").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	ws.Route(ws.POST("/add").To(handler.CreateUser).
		Doc("Add iam restapi"))
}
