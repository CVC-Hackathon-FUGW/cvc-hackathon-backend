package services

import (
	"context"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
)

type BorrowerService struct {
	ctx               context.Context
	datastoreBorrower models.DatastoreBorrower
}

func NewBorrowerService(ctx context.Context, datastoreBorrower models.DatastoreBorrower) *BorrowerService {
	return &BorrowerService{
		ctx:               ctx,
		datastoreBorrower: datastoreBorrower,
	}
}

func (p *BorrowerService) Create(borrower *models.Borrower) error {
	ctx := p.ctx
	_, err := p.datastoreBorrower.Create(ctx, borrower)
	return err
}

func (p *BorrowerService) Show(id *string) (*models.Borrower, error) {
	ctx := p.ctx
	item, err := p.datastoreBorrower.FindByID(ctx, id)
	return item, err
}

func (p *BorrowerService) List() ([]*models.Borrower, error) {
	ctx := p.ctx
	items, err := p.datastoreBorrower.List(ctx)
	return items, err
}

func (p *BorrowerService) Update(params *models.Borrower) (*models.Borrower, error) {
	ctx := p.ctx
	item, err := p.datastoreBorrower.Update(ctx, params)
	return item, err
}

func (p *BorrowerService) Delete(id *string) error {
	ctx := p.ctx
	err := p.datastoreBorrower.Delete(ctx, id)
	return err
}
