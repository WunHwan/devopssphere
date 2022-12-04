package tenant

import (
	"gorm.io/gorm"
	"time"
)

type Workspace struct {
	Name      string    `gorm:"name"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

func (Workspace) TableName() string {
	return "devopsspshere_workspace"
}

func CreateWorkspace(tx *gorm.DB, name string) (*Workspace, error) {
	var workspace = Workspace{
		Name: name,
	}

	if err := tx.Create(&workspace).Error; err != nil {
		return nil, err
	}
	return &workspace, nil
}

func DeleteWorkspace(tx *gorm.DB, name string) error {
	return tx.Delete(&Workspace{Name: name}).Error
}

func QueryOne(tx *gorm.DB, name string) (*Workspace, error) {
	workspace := new(Workspace)
	if err := tx.Where("name=?", name).Find(workspace).Error; err != nil {
		return nil, err
	}

	return workspace, nil
}
