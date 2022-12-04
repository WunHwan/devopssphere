package tenant

import (
	"gorm.io/gorm"
	api "io.github/devopssphere/pkg/api/tenant"
	sql "io.github/devopssphere/pkg/storage/tenant"
)

type ManagementInterface interface {
	CreateNamespace(workspace string) error
	FindWorkspace(workspace string) (*api.Workspace, error)
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

func (o *tenantOperator) CreateNamespace(workspace string) error {
	_, err := sql.CreateWorkspace(o.db, workspace)
	if err != nil {
		return err
	}
	return nil
}

func (o *tenantOperator) FindWorkspace(workspace string) (*api.Workspace, error) {
	one, err := sql.QueryOne(o.db, workspace)
	if err != nil {
		return nil, err
	}

	return &api.Workspace{
		Name:      one.Name,
		CreatedAt: one.CreatedAt,
		UpdatedAt: one.UpdatedAt,
	}, nil
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
