package tenant

import (
	"github.com/emicklei/go-restful"
	"io.github/devopssphere/pkg/models/tenant"
)

func AddToContainer(container *restful.Container, operator tenant.ManagementInterface) error {
	handler := NewTenantHandler(operator)

	ws := new(restful.WebService)
	ws.Path("/tenant").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	ws.Route(ws.POST("/workspace").To(handler.CreateWorkspace)).
		Doc("Create workspace restapi")

	container.Add(ws)
	return nil
}
