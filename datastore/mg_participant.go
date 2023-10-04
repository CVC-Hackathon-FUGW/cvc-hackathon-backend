package datastore

import (
	"context"
	"errors"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

type DatastoreParticipantMG struct {
	participantCollection *mongo.Collection
}

func NewDatastoreParticipantMG(participantCollection *mongo.Collection) *DatastoreParticipantMG {
	return &DatastoreParticipantMG{participantCollection}
}

var _ models.DatastoreParticipant = (*DatastoreParticipantMG)(nil)

func (ds DatastoreParticipantMG) Create(ctx context.Context, params *models.Participant) (*models.Participant, error) {
	count, err := ds.participantCollection.CountDocuments(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	params.IsActive = true
	params.ParticipantId = int(count) + 1

	_, err = ds.participantCollection.InsertOne(ctx, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (ds DatastoreParticipantMG) FindByID(ctx context.Context, id *string) (*models.Participant, error) {
	var participant *models.Participant
	idInt, err := strconv.Atoi(*id)
	if err != nil {
		return nil, err
	}

	query := bson.D{bson.E{Key: "participant_id", Value: idInt}}
	err = ds.participantCollection.FindOne(ctx, query).Decode(&participant)
	return participant, err
}

func (ds DatastoreParticipantMG) List(ctx context.Context) ([]*models.Participant, error) {
	var participants []*models.Participant
	cursor, err := ds.participantCollection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var participant models.Participant
		err := cursor.Decode(&participant)
		if err != nil {
			return nil, err
		}
		if participant.IsActive == true {
			participants = append(participants, &participant)
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return participants, nil
}

func (ds DatastoreParticipantMG) Update(ctx context.Context, params *models.Participant) (*models.Participant, error) {
	var participant *models.Participant
	query := bson.D{bson.E{Key: "participant_id", Value: params.ParticipantId}}
	err := ds.participantCollection.FindOne(ctx, query).Decode(&participant)
	if err != nil {
		return nil, err
	}

	if params.ProjectName != nil {
		participant.ProjectName = params.ProjectName
	}

	if params.ProjectId != nil {
		participant.ProjectId = params.ProjectId
	}

	if params.FundAttended != nil {
		participant.FundAttended = params.FundAttended
	}

	if params.ParticipantAddress != nil {
		participant.ParticipantAddress = params.ParticipantAddress
	}

	filter := bson.D{primitive.E{Key: "participant_id", Value: params.ParticipantId}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{
			Key: "project_id", Value: participant.ProjectId,
		},
		primitive.E{Key: "project_name", Value: participant.ProjectName},
		primitive.E{Key: "fund_attended", Value: participant.FundAttended},
		primitive.E{Key: "participant_address", Value: participant.ParticipantAddress},
	}}}

	var updatedParticipant *models.Participant
	err = ds.participantCollection.FindOneAndUpdate(ctx, filter, update).Decode(&updatedParticipant)
	if err != nil {
		return nil, errors.New("no matched document found for update")
	}

	return params, nil
}

func (ds DatastoreParticipantMG) Delete(ctx context.Context, id *string) error {
	idInt, err := strconv.Atoi(*id)
	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "participant_id", Value: idInt}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "participant_id", Value: idInt},
			primitive.E{Key: "is_active", Value: false},
		}}}

	result, _ := ds.participantCollection.UpdateOne(ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}

func (ds DatastoreParticipantMG) FindByAddress(ctx context.Context, participantAddress *string) ([]*models.Participant, error) {
	filter := bson.D{{Key: "participant_address", Value: participantAddress}}

	var participants []*models.Participant
	cursor, err := ds.participantCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var participant models.Participant
		err := cursor.Decode(&participant)
		if err != nil {
			return nil, err
		}
		if participant.IsActive == true {
			participants = append(participants, &participant)
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return participants, nil
}
