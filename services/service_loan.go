package services

import (
	"context"
	"errors"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	utils "github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/pkg"
)

type LoanService struct {
	ctx           context.Context
	datastoreLoan models.DatastoreLoan
}

func NewLoanService(ctx context.Context, datastoreLoan models.DatastoreLoan) *LoanService {
	return &LoanService{
		ctx:           ctx,
		datastoreLoan: datastoreLoan,
	}
}

func (p *LoanService) Create(Loan *models.Loan) error {
	ctx := p.ctx

	if ok := utils.ValidateAddress(Loan.TokenAddress); !ok {
		return errors.New("invalid token address")
	}

	_, err := p.datastoreLoan.Create(ctx, Loan)
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

	if params.TokenAddress != "" {
		if ok := utils.ValidateAddress(params.TokenAddress); !ok {
			return nil, errors.New("invalid token address")
		}
	}

	item, err := p.datastoreLoan.Update(ctx, params)
	return item, err
}

func (p *LoanService) Delete(id *string) error {
	ctx := p.ctx
	err := p.datastoreLoan.Delete(ctx, id)
	return err
}
