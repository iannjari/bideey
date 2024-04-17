package service

import (
	"bideey/model"
	"bideey/repo"
	"bideey/util"
	"errors"

	"github.com/google/uuid"
)

type BidsService struct {
}

var repository *repo.Repository

func NewBidsService() *BidsService {
	repository = repo.NewRepository()
	return &BidsService{}
}

func CreateBid(bid *model.Bid) (entity *model.Bid, err error) {

	if err := validate(bid); err != nil {
		return nil, err
	}

	if biddable, err := repository.GetById(&bid.BiddableId); err != nil {
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

	savedBid := *result

	bidCast, ok := savedBid.(model.Bid)

	if ok {
		return &bidCast, nil
	} else {
		return nil, errors.New("casting error")
	}
}

func UpdateBid(bid *model.Bid) (entity *model.Bid, err error) {
	if err := validate(bid); err != nil {
		return nil, err
	}

	if biddable, err := repository.GetById(&bid.BiddableId); err != nil {
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

	dbBid := *result

	// only update amount
	bidCast, ok := dbBid.(model.Bid)

	if !ok {
		return nil, errors.New("casting error")
	}

	bidCast.Amount = bid.Amount

	result, err = repository.Save(bidCast)

	if err != nil {
		return
	}

	savedBid := *result

	castedBid, ok := savedBid.(model.Bid)

	if ok {
		return &castedBid, nil
	} else {
		return nil, errors.New("casting error")
	}
}

func validate(bid *model.Bid) error {
	if bid == nil {

		return errors.New("bid cannot be nil")

	}

	if bid.Amount <= 0 {
		return errors.New("bid amount cannot be 0 or below")
	}

	if bid.BiddableId == uuid.Nil {
		return errors.New("bidabble id cannot be nil")
	}

	return nil
}
