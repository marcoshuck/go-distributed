package sum

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Sum struct {
	gorm.Model
	UUID string `json:"uuid"`
	IP string `json:"ip""`
	User string `json:"user" gorm:"unique"`
	Total int64 `json:"total" gorm:"default:0"`
}

func (Sum) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("uuid", uuid.Must(uuid.NewUUID()).String())
	return nil
}
