package tenant

import (
	"github.com/emicklei/go-restful"
	api "io.github/devopssphere/pkg/api/tenant"
	"io.github/devopssphere/pkg/models/tenant"
	"net/http"
)

type Handler struct {
	operator tenant.ManagementInterface
}

func NewTenantHandler(operator tenant.ManagementInterface) *Handler {
	return &Handler{
		operator: operator,
	}
}

func (h *Handler) CreateWorkspace(req *restful.Request, resp *restful.Response) {
	workspace := new(api.Workspace)
	var err error

	err = req.ReadEntity(&workspace)
	if err != nil {
		_ = resp.WriteError(http.StatusBadRequest, err)
		return
	}

	workspace, err = h.operator.CreateNamespace(workspace)
	if err != nil {
		_ = resp.WriteError(http.StatusInternalServerError, err)
		return
	}

	err = resp.WriteAsJson(workspace)
	if err != nil {

	}
}
