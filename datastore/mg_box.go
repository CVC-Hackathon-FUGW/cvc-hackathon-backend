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

type DatastoreBoxMG struct {
	boxCollection *mongo.Collection
}

func NewDatastoreBoxMG(boxCollection *mongo.Collection) *DatastoreBoxMG {
	return &DatastoreBoxMG{boxCollection}
}

var _ models.DatastoreBox = (*DatastoreBoxMG)(nil)

func (ds DatastoreBoxMG) Create(ctx context.Context, params *models.Box) (*models.Box, error) {
	count, err := ds.boxCollection.CountDocuments(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	params.IsActive = true
	params.BoxId = int(count) + 1

	_, err = ds.boxCollection.InsertOne(ctx, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (ds DatastoreBoxMG) FindByID(ctx context.Context, id *string) (*models.Box, error) {
	var box *models.Box
	idInt, err := strconv.Atoi(*id)
	if err != nil {
		return nil, err
	}

	query := bson.D{bson.E{Key: "box_id", Value: idInt}}
	err = ds.boxCollection.FindOne(ctx, query).Decode(&box)
	return box, err
}

func (ds DatastoreBoxMG) List(ctx context.Context) ([]*models.Box, error) {
	var boxes []*models.Box
	cursor, err := ds.boxCollection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var box models.Box
		err := cursor.Decode(&box)
		if err != nil {
			return nil, err
		}
		if box.IsActive == true {
			boxes = append(boxes, &box)
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return boxes, nil
}

func (ds DatastoreBoxMG) Update(ctx context.Context, params *models.Box) (*models.Box, error) {
	var box *models.Box
	query := bson.D{bson.E{Key: "box_id", Value: params.BoxId}}
	err := ds.boxCollection.FindOne(ctx, query).Decode(&box)
	if err != nil {
		return nil, err
	}

	if params.BoxAddress != nil {
		box.BoxAddress = params.BoxAddress
	}

	if params.IsOpened != nil {
		box.IsOpened = params.IsOpened
	}

	if params.BoxPrice != nil {
		box.BoxPrice = params.BoxPrice
	}

	if params.Owner != nil {
		box.Owner = params.Owner
	}

	if params.TokenId != nil {
		box.TokenId = params.TokenId
	}

	filter := bson.D{primitive.E{Key: "box_id", Value: params.BoxId}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "box_address", Value: box.BoxAddress},
		primitive.E{Key: "box_price", Value: box.BoxPrice},
		primitive.E{Key: "is_opened", Value: box.IsOpened},
		primitive.E{Key: "owner", Value: box.Owner},
		primitive.E{Key: "token_id", Value: box.TokenId},
	}}}

	var updatedBox *models.Box
	err = ds.boxCollection.FindOneAndUpdate(ctx, filter, update).Decode(&updatedBox)
	if err != nil {
		return nil, errors.New("no matched document found for update")
	}

	return params, nil
}

func (ds DatastoreBoxMG) Delete(ctx context.Context, id *string) error {
	idInt, err := strconv.Atoi(*id)
	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "box_id", Value: idInt}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "box_id", Value: idInt},
			primitive.E{Key: "is_active", Value: false},
		}}}

	result, _ := ds.boxCollection.UpdateOne(ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}

func (ds DatastoreBoxMG) FindByAddress(ctx context.Context, boxAddress *string) ([]*models.Box, error) {
	filter := bson.D{{Key: "box_address", Value: boxAddress}}

	var boxes []*models.Box
	cursor, err := ds.boxCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var box models.Box
		err := cursor.Decode(&box)
		if err != nil {
			return nil, err
		}
		if box.IsActive == true {
			boxes = append(boxes, &box)
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return boxes, nil
}
