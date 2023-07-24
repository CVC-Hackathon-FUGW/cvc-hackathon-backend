package services

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/enum"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	utils "github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/pkg"
)

type LoanService struct {
	ctx           context.Context
	datastoreLoan models.DatastoreLoan
	datastorePool models.DatastorePool
}

func NewLoanService(ctx context.Context, datastoreLoan models.DatastoreLoan, datastorePool models.DatastorePool) *LoanService {
	return &LoanService{
		ctx:           ctx,
		datastoreLoan: datastoreLoan,
		datastorePool: datastorePool,
	}
}

func (p *LoanService) Create(loan *models.Loan) error {
	ctx := p.ctx

	if ok := utils.ValidateAddress(*loan.TokenAddress); !ok {
		return errors.New("invalid token address")
	}

	poolIdString := strconv.Itoa(*loan.PoolId)
	pool, err := p.datastorePool.FindByID(ctx, &poolIdString)
	if err != nil {
		return errors.New("invalid poolID")
	}

	// update pool
	updateTotal := *pool.TotalPoolAmount + *loan.Amount
	pool.TotalPoolAmount = &updateTotal

	_, err = p.datastorePool.Update(ctx, pool)
	if err != nil {
		return err
	}

	_, err = p.datastoreLoan.Create(ctx, loan)
	return err
}

func (p *LoanService) Show(id *string) (*models.Loan, error) {
	ctx := p.ctx
	item, err := p.datastoreLoan.FindByID(ctx, id)
	return item, err
}

func (p *LoanService) List() ([]*models.Loan, error) {
	ctx := p.ctx
	items, err := p.datastoreLoan.List(ctx)
	return items, err
}

func (p *LoanService) Update(params *models.Loan) (*models.Loan, error) {
	ctx := p.ctx

	if params.TokenAddress != nil {
		if ok := utils.ValidateAddress(*params.TokenAddress); !ok {
			return nil, errors.New("invalid token address")
		}
	}

	poolIdString := strconv.Itoa(*params.PoolId)
	pool, err := p.datastorePool.FindByID(ctx, &poolIdString)
	if err != nil {
		return nil, errors.New("invalid poolID")
	}

	// update pool

	updateTotal := *pool.TotalPoolAmount - *params.Amount
	if params.PoolId != nil {
		pool.TotalPoolAmount = &(updateTotal)
	}

	_, err = p.datastorePool.Update(ctx, pool)
	if err != nil {
		return nil, err
	}

	item, err := p.datastoreLoan.Update(ctx, params)
	return item, err
}

func (p *LoanService) DeleteWithUpdatePool(id *string) error {
	ctx := p.ctx
	loan, err := p.datastoreLoan.FindByID(ctx, id)
	if err != nil {
		return err
	}

	err = p.datastoreLoan.Delete(ctx, id)
	if err != nil {
		return err
	}

	poolIdString := strconv.Itoa(*loan.PoolId)
	pool, err := p.datastorePool.FindByID(ctx, &poolIdString)
	if err != nil {
		return errors.New("invalid poolID")
	}

	// update pool
	updateTotal := *pool.TotalPoolAmount - *loan.Amount
	fmt.Println("poolID", loan.PoolId)
	if loan.PoolId != nil {
		pool.TotalPoolAmount = &updateTotal
	}
	fmt.Println("pool total", pool.TotalPoolAmount)

	_, err = p.datastorePool.Update(ctx, pool)
	if err != nil {
		return err
	}

	return err
}

func (p *LoanService) Delete(id *string) error {
	ctx := p.ctx
	err := p.datastoreLoan.Delete(ctx, id)
	return err
}

func (p *LoanService) MaxAmount(id *string) ([]*models.Loan, error) {
	ctx := p.ctx
	items, err := p.datastoreLoan.MaxAmount(ctx, id)
	return items, err
}

func (p *LoanService) CountLoans(id *string) (*enum.CountLoans, error) {
	ctx := p.ctx
	items, err := p.datastoreLoan.CountLoans(ctx, id)
	return items, err
}

func (p *LoanService) BorrowserTakeLoan(params *models.Loan) error {
	ctx := p.ctx

	if params.TokenAddress != nil {
		if ok := utils.ValidateAddress(*params.TokenAddress); !ok {
			return errors.New("invalid token address")
		}
	}

	poolIdString := strconv.Itoa(*params.PoolId)
	pool, err := p.datastorePool.FindByID(ctx, &poolIdString)
	if err != nil {
		return errors.New("invalid poolID")
	}

	// update pool
	updateTotal := *pool.TotalPoolAmount - *params.Amount
	if params.PoolId != nil {
		pool.TotalPoolAmount = &(updateTotal)
	}

	_, err = p.datastorePool.Update(ctx, pool)
	if err != nil {
		return err
	}

	return err
}
