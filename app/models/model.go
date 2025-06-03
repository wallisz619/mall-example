package models

import "time"

type BaseModel struct {
	ID uint64 `gorm:"primary_key;column:id;autoIncrement;" json:"id,omitempty"`
}

type CommonTimestampsFields struct {
	CreatedAt time.Time `gorm:"column:created_at;index" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index" json:"updated_at,omitempty"`
}
