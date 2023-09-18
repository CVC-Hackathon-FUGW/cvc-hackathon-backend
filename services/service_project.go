package services

import (
	"context"
	"errors"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	utils "github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/pkg"
)

type ProjectService struct {
	ctx              context.Context
	datastoreProject models.DatastoreProject
}

func NewProjectService(ctx context.Context, datastoreProject models.DatastoreProject) *ProjectService {
	return &ProjectService{
		ctx:              ctx,
		datastoreProject: datastoreProject,
	}
}

func (p *ProjectService) Create(params *models.Project) error {
	ctx := p.ctx
	if ok := utils.ValidateAddress(*params.ProjectAddress); !ok {
		return errors.New("invalid project address")
	}

	if ok := utils.ValidateAddress(*params.ProjectOwner); !ok {
		return errors.New("invalid owner address")
	}

	_, err := p.datastoreProject.Create(ctx, params)
	return err
}

func (p *ProjectService) Show(id *string) (*models.Project, error) {
	ctx := p.ctx
	item, err := p.datastoreProject.FindByID(ctx, id)
	return item, err
}

func (p *ProjectService) List() ([]*models.Project, error) {
	ctx := p.ctx
	items, err := p.datastoreProject.List(ctx)
	return items, err
}

func (p *ProjectService) Update(params *models.Project) (*models.Project, error) {
	ctx := p.ctx
	if params.ProjectAddress != nil {
		if ok := utils.ValidateAddress(*params.ProjectAddress); !ok {
			return nil, errors.New("invalid project address")
		}
	}

	if params.ProjectOwner != nil {
		if ok := utils.ValidateAddress(*params.ProjectOwner); !ok {
			return nil, errors.New("invalid owner address")
		}
	}

	item, err := p.datastoreProject.Update(ctx, params)
	return item, err
}

func (p *ProjectService) Delete(id *string) error {
	ctx := p.ctx
	err := p.datastoreProject.Delete(ctx, id)
	return err
}
