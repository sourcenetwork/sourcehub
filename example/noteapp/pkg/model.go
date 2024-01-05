package pkg

import (
	"gorm.io/gorm"
)

// Note is a text owned by some User
type Note struct {
	gorm.Model
	Title      string
	Body       string
	Creator    string
	LastEditor string
}
