package services

import (
	"context"
	"errors"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	utils "github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/pkg"
)

type ParticipantService struct {
	ctx                  context.Context
	datastoreParticipant models.DatastoreParticipant
}

func NewParticipantService(ctx context.Context, datastoreParticipant models.DatastoreParticipant) *ParticipantService {
	return &ParticipantService{
		ctx:                  ctx,
		datastoreParticipant: datastoreParticipant,
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
