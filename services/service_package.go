package services

import (
	"context"
	"errors"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	utils "github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/pkg"
)

type PackageService struct {
	ctx              context.Context
	datastorePackage models.DatastorePackage
}

func NewPackageService(ctx context.Context, datastorePackage models.DatastorePackage) *PackageService {
	return &PackageService{
		ctx:              ctx,
		datastorePackage: datastorePackage,
	}
}

func (p *PackageService) Create(pkg *models.Package) error {
	ctx := p.ctx
	if ok := utils.ValidateAddress(*pkg.ProjectAddress); !ok {
		return errors.New("invalid box address")
	}
	_, err := p.datastorePackage.Create(ctx, pkg)
	return err
}

func (p *PackageService) Show(id *string) (*models.Package, error) {
	ctx := p.ctx
	item, err := p.datastorePackage.FindByID(ctx, id)
	return item, err
}

func (p *PackageService) List() ([]*models.Package, error) {
	ctx := p.ctx
	items, err := p.datastorePackage.List(ctx)
	return items, err
}

func (p *PackageService) Update(params *models.Package) (*models.Package, error) {
	ctx := p.ctx
	if params.ProjectAddress != nil {
		if ok := utils.ValidateAddress(*params.ProjectAddress); !ok {
			return nil, errors.New("invalid box address")
		}
	}

	item, err := p.datastorePackage.Update(ctx, params)
	return item, err
}

func (p *PackageService) Delete(id *string) error {
	ctx := p.ctx
	err := p.datastorePackage.Delete(ctx, id)
	return err
}
