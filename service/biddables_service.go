package service

import (
	"bideey/model"
	"bideey/repo"
	"bideey/util"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type BService interface {
	CreateBiddable(b *model.Biddable) (*model.Biddable, error)
	UpdateBiddable(b *model.Biddable) (*model.Biddable, error)
	QueryBiddables() (*[]model.Biddable, error)
	DeleteBiddable(id string) error
}

type BiddablesService struct {
}

var this_repository *repo.Repository[model.Biddable]

func NewBidabblesService() *BiddablesService {
	this_repository = repo.NewRepository[model.Biddable]()
	return &BiddablesService{}
}

func (*BiddablesService) CreateBiddable(b *model.Biddable) (entity *model.Biddable, err error) {
	if err := validateBiddable(b); err != nil {
		return nil, err
	}

	fmt.Println(*b)

	b.Code = util.GetRandomStrCode()

	//@TODO: save current auth user as owner

	b.ID = uuid.New()

	result, err := this_repository.Save(b)

	if err != nil {
		return
	}
	return result, nil
}

func (*BiddablesService) UpdateBiddable(b *model.Biddable) (entity *model.Biddable, err error) {
	if err := validateBiddable(b); err != nil {
		return nil, err
	}

	// fetch biddable to update
	result, err := this_repository.GetById(&b.ID)

	if err != nil {
		return nil, err
	} else if result == nil {
		return nil, errors.New("did not find biddable with id provided")
	}

	result.GuidePrice = b.GuidePrice

	result, err = this_repository.Update(result)

	if err != nil {
		return
	}

	return result, nil
}

func (*BiddablesService) DeleteBiddable(id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return errors.New("error parsing uuid from str: " + err.Error())
	}
	return this_repository.Delete(&uuid)
}

func (*BiddablesService) QueryBiddables() (*[]model.Biddable, error) {
	return this_repository.Query()
}

func validateBiddable(biddable *model.Biddable) error {
	if biddable == nil {

		return errors.New("bid cannot be nil")

	}

	if biddable.GuidePrice <= 0 {
		return errors.New("bid guide price amount cannot be 0 or below")
	}

	if biddable.Description == "" {
		return errors.New("bidabble must have a description")
	}

	if biddable.Name == "" {
		return errors.New("bidabble must have a name")
	}

	return nil
}
