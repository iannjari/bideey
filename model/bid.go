package model

import "bideey/auth"

type Bid struct {
	Code   string
	Amount string
	Owner  auth.User
}
