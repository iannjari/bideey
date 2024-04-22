package model

import (
	// "bideey/auth"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bid struct {
	Id     uuid.UUID
	Code   string
	Amount int
	// Owner      auth.User
	BiddableID uuid.UUID
	gorm.Model
}
