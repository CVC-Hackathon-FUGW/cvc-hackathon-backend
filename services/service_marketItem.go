package services

import (
	"context"
	"errors"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	utils "github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/pkg"
)

type MarketItemService struct {
	ctx                       context.Context
	datastoreMarketItem       models.DatastoreMarketItem
	datastoreMarketCollection models.DatastoreMarketCollection
}

func NewMarketItemService(ctx context.Context, datastoreMarketItem models.DatastoreMarketItem, datastoreMarketCollection models.DatastoreMarketCollection) *MarketItemService {
	return &MarketItemService{
		ctx:                       ctx,
		datastoreMarketItem:       datastoreMarketItem,
		datastoreMarketCollection: datastoreMarketCollection,
	}
}

func (p *MarketItemService) Create(MarketItem *models.MarketItem) error {
	ctx := p.ctx

	if ok := utils.ValidateAddress(*MarketItem.TokenAddress); !ok {
		return errors.New("invalid token address")
	}

	_, err := p.datastoreMarketItem.Create(ctx, MarketItem)

	// err = p.UploadFiles(MarketItem.Image)
	// if err != nil {
	// 	return err
	// }

	return err
}

func (p *MarketItemService) Show(id *string) (*models.MarketItem, error) {
	ctx := p.ctx

	item, err := p.datastoreMarketItem.FindByID(ctx, id)
	return item, err
}

func (p *MarketItemService) List() ([]*models.MarketItem, error) {
	ctx := p.ctx
	items, err := p.datastoreMarketItem.List(ctx)
	return items, err
}

func (p *MarketItemService) Update(params *models.MarketItem) (*models.MarketItem, error) {
	ctx := p.ctx

	if params.TokenAddress != nil {
		if ok := utils.ValidateAddress(*params.TokenAddress); !ok {
			return nil, errors.New("invalid token address")
		}
	}

	item, err := p.datastoreMarketItem.Update(ctx, params)
	return item, err
}

func (p *MarketItemService) Delete(id *string) error {
	ctx := p.ctx
	err := p.datastoreMarketItem.Delete(ctx, id)
	return err
}

func (p *MarketItemService) FindByAddress(tokenAddress *string) ([]*models.MarketItem, error) {
	ctx := p.ctx
	items, err := p.datastoreMarketItem.FindByAddress(ctx, tokenAddress)
	return items, err
}

func (p *MarketItemService) BuyMarketItem(id *string) error {
	ctx := p.ctx
	marketItem, err := p.datastoreMarketItem.FindByID(ctx, id)
	if err != nil {
		return err
	}

	err = p.datastoreMarketItem.Delete(ctx, id)
	if err != nil {
		return err
	}

	marketCollection, err := p.datastoreMarketCollection.FindByID(ctx, id)
	if err != nil {
		return err
	}

	vollumUpdate := *marketCollection.Volume + *marketItem.Price
	marketCollection.Volume = &vollumUpdate
	_, err = p.datastoreMarketCollection.Update(ctx, marketCollection)

	return err
}

func (p *MarketItemService) OfferMarketItem(id *string) error {
	ctx := p.ctx
	marketItem, err := p.datastoreMarketItem.FindByID(ctx, id)
	if err != nil {
		return err
	}

	err = p.datastoreMarketItem.Delete(ctx, id)
	if err != nil {
		return err
	}

	marketCollection, err := p.datastoreMarketCollection.FindByID(ctx, id)
	if err != nil {
		return err
	}

	vollumUpdate := *marketCollection.Volume + *marketItem.CurrentOfferValue
	marketCollection.Volume = &vollumUpdate
	_, err = p.datastoreMarketCollection.Update(ctx, marketCollection)

	return err
}
