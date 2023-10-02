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

type DatastoreBoxCollectionMG struct {
	BoxCollection *mongo.Collection
}

func NewDatastoreBoxCollectionMG(boxCollection *mongo.Collection) *DatastoreBoxCollectionMG {
	return &DatastoreBoxCollectionMG{boxCollection}
}

func (ds DatastoreBoxCollectionMG) Create(ctx context.Context, params *models.BoxCollection) (*models.BoxCollection, error) {
	count, err := ds.BoxCollection.CountDocuments(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	params.IsActive = true
	params.BoxCollectionId = int(count) + 1

	_, err = ds.BoxCollection.InsertOne(ctx, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (ds DatastoreBoxCollectionMG) FindByID(ctx context.Context, id *string) (*models.BoxCollection, error) {
	var boxCollection *models.BoxCollection
	idInt, err := strconv.Atoi(*id)
	if err != nil {
		return nil, err
	}

	query := bson.D{bson.E{Key: "box_collection_id", Value: idInt}}
	err = ds.BoxCollection.FindOne(ctx, query).Decode(&boxCollection)
	return boxCollection, err
}

func (ds DatastoreBoxCollectionMG) List(ctx context.Context) ([]*models.BoxCollection, error) {
	var boxCollections []*models.BoxCollection
	cursor, err := ds.BoxCollection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var boxCollection models.BoxCollection
		err := cursor.Decode(&boxCollection)
		if err != nil {
			return nil, err
		}
		if boxCollection.IsActive == true {
			boxCollections = append(boxCollections, &boxCollection)
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return boxCollections, nil
}

func (ds DatastoreBoxCollectionMG) Update(ctx context.Context, params *models.BoxCollection) (*models.BoxCollection, error) {
	var boxCollection *models.BoxCollection
	query := bson.D{bson.E{Key: "box_collection_id", Value: params.BoxCollectionId}}
	err := ds.BoxCollection.FindOne(ctx, query).Decode(&boxCollection)
	if err != nil {
		return nil, err
	}

	if params.BoxCollectionAddress != nil {
		boxCollection.BoxCollectionAddress = params.BoxCollectionAddress
	}

	if params.Image != nil {
		boxCollection.Image = params.Image
	}

	filter := bson.D{primitive.E{Key: "box_collection_id", Value: params.BoxCollectionId}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "box_collection_address", Value: boxCollection.BoxCollectionAddress},
		primitive.E{Key: "image", Value: boxCollection.Image},
	}}}

	var updatedBoxCollection *models.BoxCollection
	err = ds.BoxCollection.FindOneAndUpdate(ctx, filter, update).Decode(&updatedBoxCollection)
	if err != nil {
		return nil, errors.New("no matched document found for update")
	}

	return params, nil
}

func (ds DatastoreBoxCollectionMG) Delete(ctx context.Context, id *string) error {
	idInt, err := strconv.Atoi(*id)
	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "box_collection_id", Value: idInt}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "box_collection_id", Value: idInt},
			primitive.E{Key: "is_active", Value: false},
		}}}

	result, _ := ds.BoxCollection.UpdateOne(ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
