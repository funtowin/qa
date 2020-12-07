package model

import (
	"github.com/jinzhu/gorm"
	util "qa/util"
	"time"
)

// GORMBase struct
type GORMBase struct {
	ID        uint64
	CreatedAt string
	UpdatedAt string
	DeletedAt *string `sql:"index"`
}

// BeforeCreate func
func (m *GORMBase) BeforeCreate(scope *gorm.Scope) error {

	id, _ := util.GetID()
	m.ID = id

	if m.UpdatedAt == "" {
		scope.SetColumn("UpdatedAt", time.Now().Format("2006-01-02 15:04:05"))
	}

	scope.SetColumn("CreatedAt", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

// BeforeUpdate func
func (m *GORMBase) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}
