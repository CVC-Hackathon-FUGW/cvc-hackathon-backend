package datastore

import (
	"context"
	"errors"
	"strconv"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/enum"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatastoreCheckinMG struct {
	CheckinCollection *mongo.Collection
}

func NewDatastoreCheckinMG(CheckinCollection *mongo.Collection) *DatastoreCheckinMG {
	return &DatastoreCheckinMG{CheckinCollection}
}

var _ models.DatastoreCheckin = (*DatastoreCheckinMG)(nil)

func (ds DatastoreCheckinMG) Create(ctx context.Context, params *models.Checkin) (*models.Checkin, error) {
	count, err := ds.CheckinCollection.CountDocuments(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	params.CheckinId = int(count) + 1
	params.IsActive = true

	_, err = ds.CheckinCollection.InsertOne(ctx, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (ds DatastoreCheckinMG) FindByID(ctx context.Context, id *string) (*models.Checkin, error) {
	var checkin *models.Checkin

	idInt, err := strconv.Atoi(*id)
	if err != nil {
		return nil, err
	}
	query := bson.D{bson.E{Key: "checkin_id", Value: idInt}}
	err = ds.CheckinCollection.FindOne(ctx, query).Decode(&checkin)
	return checkin, err
}

func (ds DatastoreCheckinMG) List(ctx context.Context, params enum.CheckinParams) ([]*models.Checkin, error) {
	var checkins []*models.Checkin
	filter := bson.D{{}}
	opts := options.Find().SetSort(bson.D{{Key: "volume", Value: 1}})
	if params.Sort == "volume" {
		opts = options.Find().SetSort(bson.D{{Key: "volume", Value: -1}})
	}

	cursor, err := ds.CheckinCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var Checkin models.Checkin
		err := cursor.Decode(&Checkin)
		if err != nil {
			return nil, err
		}
		if Checkin.IsActive == true {
			checkins = append(checkins, &Checkin)
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return checkins, nil
}

func (ds DatastoreCheckinMG) Update(ctx context.Context, params *models.Checkin) (*models.Checkin, error) {
	var checkinDB *models.Checkin

	query := bson.D{bson.E{Key: "checkin_id", Value: params.CheckinId}}
	err := ds.CheckinCollection.FindOne(ctx, query).Decode(&checkinDB)
	if err != nil {
		return nil, err
	}

	if params.CurrentTokenAmount != nil {
		checkinDB.CurrentTokenAmount = params.CurrentTokenAmount
	}

	if params.NumberChecked != nil {
		checkinDB.NumberChecked = params.NumberChecked
	}

	if params.Wallet != nil {
		checkinDB.Wallet = params.Wallet
	}

	filter := bson.D{primitive.E{Key: "checkin_id", Value: params.CheckinId}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "current_token_amount", Value: checkinDB.CurrentTokenAmount},
			primitive.E{Key: "number_checked", Value: checkinDB.NumberChecked},
			primitive.E{Key: "wallet", Value: checkinDB.Wallet},
		}}}

	var checkinUpdated *models.Checkin
	err = ds.CheckinCollection.FindOneAndUpdate(ctx, filter, update).Decode(&checkinUpdated)
	if err != nil {
		return nil, errors.New("no matched document found for update")
	}

	return checkinUpdated, nil
}

func (ds DatastoreCheckinMG) Delete(ctx context.Context, id *string) error {
	idInt, err := strconv.Atoi(*id)
	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "checkin_id", Value: idInt}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "checkin_id", Value: idInt},
			primitive.E{Key: "is_active", Value: false},
		}}}

	result, _ := ds.CheckinCollection.UpdateOne(ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
