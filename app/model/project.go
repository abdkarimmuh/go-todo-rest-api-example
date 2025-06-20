package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Project struct {
	gorm.Model
	Title     string `gorm:"unique" json:"title"`
	Archived  bool   `json:"archived"`
	Tasks     []Task `gorm:"ForeignKey:ProjectID" json:"tasks"`
	CreatedBy uint   `json:"created_by"`
	UpdatedBy uint   `json:"updated_by"`
	DeletedBy uint   `json:"deleted_by"`
}

func (p *Project) Archive() {
	p.Archived = true
}

func (p *Project) Restore() {
	p.Archived = false
}
