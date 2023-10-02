package services

import (
	"context"
	"errors"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	utils "github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/pkg"
)

type BoxCollectionService struct {
	ctx                    context.Context
	datastoreBoxCollection models.DatastoreBoxCollection
}

func NewBoxCollectionService(ctx context.Context, datastoreBoxCollection models.DatastoreBoxCollection) *BoxCollectionService {
	return &BoxCollectionService{
		ctx:                    ctx,
		datastoreBoxCollection: datastoreBoxCollection,
	}
}

func (p *BoxCollectionService) Create(boxCollection *models.BoxCollection) error {
	ctx := p.ctx
	if ok := utils.ValidateAddress(*boxCollection.BoxCollectionAddress); !ok {
		return errors.New("invalid participant address")
	}
	_, err := p.datastoreBoxCollection.Create(ctx, boxCollection)
	return err
}

func (p *BoxCollectionService) Show(id *string) (*models.BoxCollection, error) {
	ctx := p.ctx
	item, err := p.datastoreBoxCollection.FindByID(ctx, id)
	return item, err
}

func (p *BoxCollectionService) List() ([]*models.BoxCollection, error) {
	ctx := p.ctx
	items, err := p.datastoreBoxCollection.List(ctx)
	return items, err
}

func (p *BoxCollectionService) Update(params *models.BoxCollection) (*models.BoxCollection, error) {
	ctx := p.ctx
	if params.BoxCollectionAddress != nil {
		if ok := utils.ValidateAddress(*params.BoxCollectionAddress); !ok {
			return nil, errors.New("invalid participant address")
		}
	}

	item, err := p.datastoreBoxCollection.Update(ctx, params)
	return item, err
}

func (p *BoxCollectionService) Delete(id *string) error {
	ctx := p.ctx
	err := p.datastoreBoxCollection.Delete(ctx, id)
	return err
}
