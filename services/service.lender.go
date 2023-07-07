package services

import (
	"context"
	"errors"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	utils "github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/pkg"
)

type LenderService struct {
	ctx             context.Context
	datastoreLender models.DatastoreLender
}

func NewLenderService(ctx context.Context, datastoreLender models.DatastoreLender) *LenderService {
	return &LenderService{
		ctx:             ctx,
		datastoreLender: datastoreLender,
	}
}

func (p *LenderService) Create(lender *models.Lender) error {
	ctx := p.ctx

	if ok := utils.ValidateAddress(lender.WalletAddress); !ok {
		return errors.New("invalid wallet address")
	}

	_, err := p.datastoreLender.Create(ctx, lender)
	return err
}

func (p *LenderService) Show(id *string) (*models.Lender, error) {
	ctx := p.ctx
	item, err := p.datastoreLender.FindByID(ctx, id)
	return item, err
}

func (p *LenderService) List() ([]*models.Lender, error) {
	ctx := p.ctx
	items, err := p.datastoreLender.List(ctx)
	return items, err
}

func (p *LenderService) Update(params *models.Lender) (*models.Lender, error) {
	ctx := p.ctx

	if params.WalletAddress != "" {
		if ok := utils.ValidateAddress(params.WalletAddress); !ok {
			return nil, errors.New("invalid wallet address")
		}
	}
	item, err := p.datastoreLender.Update(ctx, params)
	return item, err
}

func (p *LenderService) Delete(id *string) error {
	ctx := p.ctx
	err := p.datastoreLender.Delete(ctx, id)
	return err
}
