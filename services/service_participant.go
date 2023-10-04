package services

import (
	"context"
	"errors"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	utils "github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/pkg"
	"strconv"
)

type ParticipantService struct {
	ctx                  context.Context
	datastoreParticipant models.DatastoreParticipant
	datastoreProject     models.DatastoreProject
}

func NewParticipantService(ctx context.Context, datastoreParticipant models.DatastoreParticipant, datastoreProject models.DatastoreProject) *ParticipantService {
	return &ParticipantService{
		ctx:                  ctx,
		datastoreParticipant: datastoreParticipant,
		datastoreProject:     datastoreProject,
	}
}

func (p *ParticipantService) Create(participant *models.Participant) error {
	ctx := p.ctx
	if ok := utils.ValidateAddress(*participant.ParticipantAddress); !ok {
		return errors.New("invalid participant address")
	}
	_, err := p.datastoreParticipant.Create(ctx, participant)
	return err
}

func (p *ParticipantService) Show(id *string) (*models.Participant, error) {
	ctx := p.ctx
	item, err := p.datastoreParticipant.FindByID(ctx, id)
	return item, err
}

func (p *ParticipantService) List() ([]*models.Participant, error) {
	ctx := p.ctx
	items, err := p.datastoreParticipant.List(ctx)
	return items, err
}

func (p *ParticipantService) Update(params *models.Participant) (*models.Participant, error) {
	ctx := p.ctx
	if params.ParticipantAddress != nil {
		if ok := utils.ValidateAddress(*params.ParticipantAddress); !ok {
			return nil, errors.New("invalid participant address")
		}
	}

	item, err := p.datastoreParticipant.Update(ctx, params)
	return item, err
}

func (p *ParticipantService) Delete(id *string) error {
	ctx := p.ctx
	err := p.datastoreParticipant.Delete(ctx, id)
	return err
}

func (p *ParticipantService) Invest(params *models.Participant) error {
	ctx := p.ctx
	if ok := utils.ValidateAddress(*params.ParticipantAddress); !ok {
		return errors.New("invalid participant address")
	}

	//update the project raised amount
	amount := params.FundAttended
	projectId := strconv.Itoa(*params.ProjectId)

	project, err := p.datastoreProject.FindByID(ctx, &projectId)
	if err != nil {
		return errors.New("project ID Error")
	}

	updateRaised := *project.TotalFundRaised + *amount
	project.TotalFundRaised = &updateRaised

	_, err = p.datastoreProject.Update(ctx, project)
	if err != nil {
		return errors.New("update Project error")
	}

	_, err = p.datastoreParticipant.Create(ctx, params)
	if err != nil {
		return errors.New("invest Error in create participant phase")
	}

	return nil
}

func (p *ParticipantService) FindByAddress(participantAddress *string) ([]*models.Participant, error) {
	ctx := p.ctx
	items, err := p.datastoreParticipant.FindByAddress(ctx, participantAddress)
	return items, err
}
