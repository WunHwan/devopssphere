package tenant

import (
	"gorm.io/gorm"
	api "io.github/devopssphere/pkg/api/tenant"
	sql "io.github/devopssphere/pkg/storage/tenant"
)

type ManagementInterface interface {
	CreateNamespace(*api.Workspace) (*api.Workspace, error)
	GetWorkspace(workspace string) (*api.Workspace, error)
	DeleteWorkspace(workspace string) error
	ListWorkspace()
}

type tenantOperator struct {
	db *gorm.DB
}

func NewOperator(db *gorm.DB) ManagementInterface {
	return &tenantOperator{
		db: db,
	}
}

func (o *tenantOperator) CreateNamespace(workspace *api.Workspace) (*api.Workspace, error) {
	//TODO implement me
	panic("implement me")
}

func (o *tenantOperator) GetWorkspace(workspace string) (*api.Workspace, error) {
	//TODO implement me
	panic("implement me")
}

func (o *tenantOperator) ListWorkspace() {
	//TODO implement me
	panic("implement me")
}

func (o *tenantOperator) DeleteWorkspace(workspace string) error {
	err := sql.DeleteWorkspace(o.db, workspace)
	if err != nil {
		return err
	}
	return nil
}
