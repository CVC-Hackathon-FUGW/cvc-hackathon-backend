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

type PoolService struct {
	ctx           context.Context
	dataStorePool models.DatastorePool
}

func NewPoolService(ctx context.Context, datastorePool models.DatastorePool) *PoolService {
	return &PoolService{
		ctx:           ctx,
		dataStorePool: datastorePool,
	}
}

func (p *PoolService) Create(pool *models.Pool) error {
	ctx := p.ctx

	if ok := utils.ValidateAddress(*pool.TokenAddress); !ok {
		return errors.New("invalid token address")
	}

	_, err := p.dataStorePool.Create(ctx, pool)

	// err = p.UploadFiles(pool.Image)
	// if err != nil {
	// 	return err
	// }

	return err
}

func (p *PoolService) Show(id *string) (*models.Pool, error) {
	ctx := p.ctx
	fmt.Println("id show", id)
	item, err := p.dataStorePool.FindByID(ctx, id)
	return item, err
}

func (p *PoolService) List(params enum.PoolParams) ([]*models.Pool, error) {
	ctx := p.ctx
	items, err := p.dataStorePool.List(ctx, params)
	return items, err
}

func (p *PoolService) Update(params *models.Pool) (*models.Pool, error) {
	ctx := p.ctx
	if params.TokenAddress != nil {
		if ok := utils.ValidateAddress(*params.TokenAddress); !ok {
			return nil, errors.New("invalid token address")
		}
	}

	item, err := p.dataStorePool.Update(ctx, params)
	return item, err
}

func (p *PoolService) Delete(id *string) error {
	ctx := p.ctx
	err := p.dataStorePool.Delete(ctx, id)
	return err
}

func (p *PoolService) MaxAmount(id *string) ([]*models.Loan, error) {
	ctx := p.ctx
	items, err := p.dataStorePool.MaxAmount(ctx, id)
	return items, err
}

func (p *PoolService) CountLoans(id *string) (*enum.CountLoans, error) {
	ctx := p.ctx
	items, err := p.dataStorePool.CountLoans(ctx, id)
	return items, err
}

func (p *PoolService) UploadFiles(uploadFileDir string) error {

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
