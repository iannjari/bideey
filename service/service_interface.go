package service

import "bideey/model"

type Service interface {
	// bids
	CreateBid(bid *model.Bid) (entity *model.Bid, err error)
	UpdateBid(bid *model.Bid) (entity *model.Bid, err error)
	DeleteBid(id string) error
	QueryBids() (*[]model.Biddable, error)
	GetBidById(id string) (*model.Bid, error)

	// biddables
	CreateBiddable(b *model.Biddable) (*model.Biddable, error)
	UpdateBiddable(b *model.Biddable) (*model.Biddable, error)
	QueryBiddables() (*[]model.Biddable, error)
	DeleteBiddable(id string) error
	GetBiddableById(id string) (*model.Biddable, error)
}
