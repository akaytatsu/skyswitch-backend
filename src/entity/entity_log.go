package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type EntityLog struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Code      string    `json:"code" gorm:"varchar(50)"`
	Type      string    `json:"type" gorm:"varchar(50)"`
	Instance  string    `json:"instance" gorm:"varchar(90)"`
	Content   string    `json:"content" gorm:"varchar(1200)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	Error     string    `json:"error" gorm:"varchar(1200)"`
}

func NewEntityLog(entityLogParam EntityLog) (*EntityLog, error) {
	u := &EntityLog{
		ID:        entityLogParam.ID,
		Code:      entityLogParam.Code,
		Type:      entityLogParam.Type,
		Instance:  entityLogParam.Instance,
		Content:   entityLogParam.Content,
		CreatedAt: entityLogParam.CreatedAt,
		Error:     entityLogParam.Error,
	}

	return u, nil
}

func (u *EntityLog) Validate() error {
	return validator.New().Struct(u)
}

type SearchEntityLogParams struct {
	OrderBy   string `json:"order_by"`
	SortOrder string `json:"sort_order"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
	Q         string `json:"q"`
	CreatedAt string `json:"created_at"`
}
