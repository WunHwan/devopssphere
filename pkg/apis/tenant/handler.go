package tenant

import (
	"errors"
	"github.com/emicklei/go-restful"
	"go.uber.org/zap"
	"io.github/devopssphere/pkg/models/tenant"
	"net/http"
)

const (
	parameterWorkspaceBadRequest = "workspace cannot be empty"
)

type Handler struct {
	log      *zap.SugaredLogger
	operator tenant.ManagementInterface
}

func NewTenantHandler(log *zap.SugaredLogger, operator tenant.ManagementInterface) *Handler {
	return &Handler{
		log:      log,
		operator: operator,
	}
}

func (h *Handler) CreateWorkspace(req *restful.Request, resp *restful.Response) {
	workspace := req.QueryParameter("workspace")
	if len(workspace) == 0 {
		_ = resp.WriteError(http.StatusBadRequest, errors.New(parameterWorkspaceBadRequest))
		return
	}

	if err := h.operator.CreateNamespace(workspace); err != nil {
		_ = resp.WriteError(http.StatusInternalServerError, err)
		return
	}
}

func (h *Handler) FindWorkspace(req *restful.Request, resp *restful.Response) {
	workspace := req.QueryParameter("workspace")
	if len(workspace) == 0 {
		_ = resp.WriteError(http.StatusBadRequest, errors.New(parameterWorkspaceBadRequest))
		return
	}

	if workspace, err := h.operator.FindWorkspace(workspace); err != nil {
		_ = resp.WriteError(http.StatusInternalServerError, err)
	} else if err = resp.WriteAsJson(workspace); err != nil {
		h.log.Errorf("tenant find-workspace response error: %s", err.Error())
	}
}
