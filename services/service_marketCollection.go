package services

import (
	"context"
	"errors"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	utils "github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/pkg"
	aws_pkg "github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/pkg/aws"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
)

type MarketCollectionService struct {
	ctx                       context.Context
	dataStoreMarketCollection models.DatastoreMarketCollection
}

func NewMarketCollectionService(ctx context.Context, datastoreMarketCollection models.DatastoreMarketCollection) *MarketCollectionService {
	return &MarketCollectionService{
		ctx:                       ctx,
		dataStoreMarketCollection: datastoreMarketCollection,
	}
}

func (p *MarketCollectionService) Create(MarketCollection *models.MarketCollection) error {
	ctx := p.ctx

	if ok := utils.ValidateAddress(MarketCollection.TokenAddress); !ok {
		return errors.New("invalid token address")
	}

	_, err := p.dataStoreMarketCollection.Create(ctx, MarketCollection)

	// err = p.UploadFiles(MarketCollection.Image)
	// if err != nil {
	// 	return err
	// }

	return err
}

func (p *MarketCollectionService) Show(id *string) (*models.MarketCollection, error) {
	ctx := p.ctx

	item, err := p.dataStoreMarketCollection.FindByID(ctx, id)
	return item, err
}

func (p *MarketCollectionService) List() ([]*models.MarketCollection, error) {
	ctx := p.ctx
	items, err := p.dataStoreMarketCollection.List(ctx)
	return items, err
}

func (p *MarketCollectionService) Update(params *models.MarketCollection) (*models.MarketCollection, error) {
	ctx := p.ctx

	if params.TokenAddress != "" {
		if ok := utils.ValidateAddress(params.TokenAddress); !ok {
			return nil, errors.New("invalid token address")
		}
	}

	item, err := p.dataStoreMarketCollection.Update(ctx, params)
	return item, err
}

func (p *MarketCollectionService) Delete(id *string) error {
	ctx := p.ctx
	err := p.dataStoreMarketCollection.Delete(ctx, id)
	return err
}

func (p *MarketCollectionService) UploadFiles(uploadFileDir string) error {

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
