package services

import (
	"context"
	"errors"
	"os"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	utils "github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/pkg"
	aws_pkg "github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/pkg/aws"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type MarketItemService struct {
	ctx                 context.Context
	dataStoreMarketItem models.DatastoreMarketItem
}

func NewMarketItemService(ctx context.Context, datastoreMarketItem models.DatastoreMarketItem) *MarketItemService {
	return &MarketItemService{
		ctx:                 ctx,
		dataStoreMarketItem: datastoreMarketItem,
	}
}

func (p *MarketItemService) Create(MarketItem *models.MarketItem) error {
	ctx := p.ctx

	if ok := utils.ValidateAddress(*MarketItem.TokenAddress); !ok {
		return errors.New("invalid token address")
	}

	_, err := p.dataStoreMarketItem.Create(ctx, MarketItem)

	// err = p.UploadFiles(MarketItem.Image)
	// if err != nil {
	// 	return err
	// }

	return err
}

func (p *MarketItemService) Show(id *string) (*models.MarketItem, error) {
	ctx := p.ctx

	item, err := p.dataStoreMarketItem.FindByID(ctx, id)
	return item, err
}

func (p *MarketItemService) List() ([]*models.MarketItem, error) {
	ctx := p.ctx
	items, err := p.dataStoreMarketItem.List(ctx)
	return items, err
}

func (p *MarketItemService) Update(params *models.MarketItem) (*models.MarketItem, error) {
	ctx := p.ctx

	if params.TokenAddress != nil {
		if ok := utils.ValidateAddress(*params.TokenAddress); !ok {
			return nil, errors.New("invalid token address")
		}
	}

	item, err := p.dataStoreMarketItem.Update(ctx, params)
	return item, err
}

func (p *MarketItemService) Delete(id *string) error {
	ctx := p.ctx
	err := p.dataStoreMarketItem.Delete(ctx, id)
	return err
}

func (p *MarketItemService) UploadFiles(uploadFileDir string) error {

	session, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("AWS_S3_REGION"))})
	if err != nil {
		return err
	}

	// Upload Files
	err = aws_pkg.UploadFile(session, uploadFileDir)
	if err != nil {
		return err
	}
	return nil
}
