package services

import (
	"context"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
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

func (p *LenderService) Create(Lender *models.Lender) error {
	ctx := p.ctx
	_, err := p.datastoreLender.Create(ctx, Lender)
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
	item, err := p.datastoreLender.Update(ctx, params)
	return item, err
}

func (p *LenderService) Delete(id *string) error {
	ctx := p.ctx
	err := p.datastoreLender.Delete(ctx, id)
	return err
}
