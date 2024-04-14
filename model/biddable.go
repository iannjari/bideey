package model

import "bideey/auth"

type Biddable struct {
	GuidePrice  int
	Name        string
	Description string
	HighestBid  int
	Bids        []Bid
	Owner       auth.User
	Category    BiddableCategory
}

type BiddableCategory struct {
}
