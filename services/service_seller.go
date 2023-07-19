package services

import (
	"context"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
)

type SellerService struct {
	ctx             context.Context
	dataStoreSeller models.DatastoreSeller
}

func NewSellerService(ctx context.Context, datastoreSeller models.DatastoreSeller) *SellerService {
	return &SellerService{
		ctx:             ctx,
		dataStoreSeller: datastoreSeller,
	}
}

func (p *SellerService) Create(Seller *models.Seller) error {
	ctx := p.ctx
	_, err := p.dataStoreSeller.Create(ctx, Seller)
	return err
}

func (p *SellerService) Show(address *string) (*models.Seller, error) {
	ctx := p.ctx
	item, err := p.dataStoreSeller.FindByAddress(ctx, address)
	return item, err
}

func (p *SellerService) List() ([]*models.Seller, error) {
	ctx := p.ctx
	items, err := p.dataStoreSeller.List(ctx)
	return items, err
}
