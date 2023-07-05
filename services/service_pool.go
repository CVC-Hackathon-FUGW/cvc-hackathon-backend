package services

import (
	"context"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
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
	_, err := p.dataStorePool.Create(ctx, pool)
	return err
}

func (p *PoolService) Show(id *string) (*models.Pool, error) {
	ctx := p.ctx
	item, err := p.dataStorePool.FindByID(ctx, id)
	return item, err
}

func (p *PoolService) List() ([]*models.Pool, error) {
	ctx := p.ctx
	items, err := p.dataStorePool.List(ctx)
	return items, err
}

func (p *PoolService) Update(params *models.Pool) (*models.Pool, error) {
	ctx := p.ctx
	item, err := p.dataStorePool.Update(ctx, params)
	return item, err
}

func (p *PoolService) Delete(id *string) error {
	ctx := p.ctx
	err := p.dataStorePool.Delete(ctx, id)
	return err
}
