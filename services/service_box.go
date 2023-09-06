package services

import (
	"context"
	"errors"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	utils "github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/pkg"
)

type BoxService struct {
	ctx          context.Context
	datastoreBox models.DatastoreBox
}

func NewBoxService(ctx context.Context, datastoreBox models.DatastoreBox) *BoxService {
	return &BoxService{
		ctx:          ctx,
		datastoreBox: datastoreBox,
	}
}

func (b *BoxService) Create(box *models.Box) error {
	ctx := b.ctx
	if ok := utils.ValidateAddress(*box.BoxAddress); !ok {
		return errors.New("invalid box address")
	}
	if ok := utils.ValidateAddress(*box.Owner); !ok {
		return errors.New("invalid owner address")
	}
	_, err := b.datastoreBox.Create(ctx, box)
	return err
}

func (b *BoxService) Show(id *string) (*models.Box, error) {
	ctx := b.ctx
	item, err := b.datastoreBox.FindByID(ctx, id)
	return item, err
}

func (b *BoxService) List() ([]*models.Box, error) {
	ctx := b.ctx
	items, err := b.datastoreBox.List(ctx)
	return items, err
}

func (b *BoxService) Update(params *models.Box) (*models.Box, error) {
	ctx := b.ctx
	if params.BoxAddress != nil {
		if ok := utils.ValidateAddress(*params.BoxAddress); !ok {
			return nil, errors.New("invalid box address")
		}
	}

	if params.Owner != nil {
		if ok := utils.ValidateAddress(*params.Owner); !ok {
			return nil, errors.New("invalid owner address")
		}
	}

	item, err := b.datastoreBox.Update(ctx, params)
	return item, err
}

func (b *BoxService) Delete(id *string) error {
	ctx := b.ctx
	err := b.datastoreBox.Delete(ctx, id)
	return err
}