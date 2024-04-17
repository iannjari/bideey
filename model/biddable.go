package model

import (
	"bideey/auth"
)

type Biddable struct {
	Code        string
	GuidePrice  int
	Name        string
	Description string
	HighestBid  int
	Bids        []Bid
	Owner       auth.User
	Category    BiddableCategory
}

type BiddableCategory int

func (b BiddableCategory) GetIndex() int {
	return int(b)
}

func (b BiddableCategory) Equals(other BiddableCategory) bool {
	return int(b) == int(other)
}

const (
	OTHER BiddableCategory = iota
	ART
	SACCO
	HOUSEHOLLD
	PET
)
