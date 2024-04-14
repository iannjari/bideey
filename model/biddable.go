package model

import "bideey/auth"

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

const (
	OTHER BiddableCategory = iota
	ART
	SACCO
	HOUSEHOLLD
	PET
)
