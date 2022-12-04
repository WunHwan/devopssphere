package tenant

import (
	"github.com/emicklei/go-restful"
	"go.uber.org/zap"
	"io.github/devopssphere/pkg/models/tenant"
)

func AddToContainer(container *restful.Container, log *zap.SugaredLogger, operator tenant.ManagementInterface) error {
	handler := NewTenantHandler(log, operator)

	ws := new(restful.WebService)
	ws.Path("/tenant")

	ws.Route(ws.POST("/workspace").To(handler.CreateWorkspace)).
		Doc("Create workspace restapi")
	ws.Route(ws.GET("/workspace").To(handler.FindWorkspace)).Produces(restful.MIME_JSON).
		Doc("Create workspace restapi")

	container.Add(ws)
	return nil
}
