package model

import "time"

// Memo represents memo
type Memo struct {
	ID        int        `json:"id" db:"id"`
	Date      string     `json:"date" db:"date" validate:"date_validation,required"`
	Title     string     `json:"title" db:"title" validate:"required"`
	Content   string     `json:"content" db:"content" validate:"required"`
	Check     bool       `json:"check" db:"check"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

// TableName return table name
func (m *Memo) TableName() string {
	return "memo"
}
