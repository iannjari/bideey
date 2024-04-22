package service

import (
	"bideey/model"
	"bideey/repo"
	"bideey/util"
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	CreateBid(bid *model.Bid) (entity *model.Bid, err error)
	UpdateBid(bid *model.Bid) (entity *model.Bid, err error)
}

type BidsService struct {
}

var repository *repo.Repository[model.Bid]

func NewBidsService() *BidsService {
	repository = repo.NewRepository[model.Bid]()
	return &BidsService{}
}

func (*BidsService) CreateBid(bid *model.Bid) (entity *model.Bid, err error) {

	if err := validate(bid); err != nil {
		return nil, err
	}

	if biddable, err := repository.GetById(&bid.BiddableID); err != nil {
		return nil, err
	} else if biddable == nil {
		return nil, errors.New("did not find biddable with id provided")
	}

	bid.Code = util.GetRandomStrCode()

	//@TODO: save current auth user as owner

	bid.Id = uuid.New()

	result, err := repository.Save(bid)

	if err != nil {
		return
	}
	return result, nil
}

func (*BidsService) UpdateBid(bid *model.Bid) (entity *model.Bid, err error) {
	if err := validate(bid); err != nil {
		return nil, err
	}

	if biddable, err := repository.GetById(&bid.BiddableID); err != nil {
		return nil, err
	} else if biddable == nil {
		return nil, errors.New("did not find bididdable with id provided")
	}

	// fetch bid to update
	result, err := repository.GetById(&bid.Id)

	if err != nil {
		return nil, err
	} else if result == nil {
		return nil, errors.New("did not find bid with id provided")
	}

	result.Amount = bid.Amount

	result, err = repository.Save(result)

	if err != nil {
		return
	}

	return result, nil
}

func validate(bid *model.Bid) error {
	if bid == nil {

		return errors.New("bid cannot be nil")

	}

	if bid.Amount <= 0 {
		return errors.New("bid amount cannot be 0 or below")
	}

	if bid.BiddableID == uuid.Nil {
		return errors.New("bidabble id cannot be nil")
	}

	return nil
}
