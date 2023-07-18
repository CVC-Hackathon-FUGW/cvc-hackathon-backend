package services

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/enum"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	utils "github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/pkg"
	aws_pkg "github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/pkg/aws"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type CheckinService struct {
	ctx              context.Context
	dataStoreCheckin models.DatastoreCheckin
}

func NewCheckinService(ctx context.Context, datastoreCheckin models.DatastoreCheckin) *CheckinService {
	return &CheckinService{
		ctx:              ctx,
		dataStoreCheckin: datastoreCheckin,
	}
}

func (p *CheckinService) Create(Checkin *models.Checkin) error {
	ctx := p.ctx

	if ok := utils.ValidateAddress(*Checkin.Wallet); !ok {
		return errors.New("invalid token address")
	}

	_, err := p.dataStoreCheckin.Create(ctx, Checkin)

	return err
}

func (p *CheckinService) Show(id *string) (*models.Checkin, error) {
	ctx := p.ctx
	fmt.Println("id show", id)
	item, err := p.dataStoreCheckin.FindByID(ctx, id)
	return item, err
}

func (p *CheckinService) List(params enum.CheckinParams) ([]*models.Checkin, error) {
	ctx := p.ctx
	items, err := p.dataStoreCheckin.List(ctx, params)
	return items, err
}

func (p *CheckinService) Update(params *models.Checkin) (*models.Checkin, error) {
	ctx := p.ctx
	if params.Wallet != nil {
		if ok := utils.ValidateAddress(*params.Wallet); !ok {
			return nil, errors.New("invalid token address")
		}
	}

	item, err := p.dataStoreCheckin.Update(ctx, params)
	return item, err
}

func (p *CheckinService) Delete(id *string) error {
	ctx := p.ctx
	err := p.dataStoreCheckin.Delete(ctx, id)
	return err
}

func (p *CheckinService) UploadFiles(uploadFileDir string) error {

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
