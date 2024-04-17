package model

import (
	"bideey/auth"

	"github.com/google/uuid"
)

type Bid struct {
	Id         uuid.UUID
	Code       string
	Amount     int
	Owner      auth.User
	BiddableId uuid.UUID
}
