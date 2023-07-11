package datastore

import (
	"context"
	"errors"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatastorePoolMG struct {
	poolCollection *mongo.Collection
}

func NewDatastorePoolMG(poolCollection *mongo.Collection) *DatastorePoolMG {
	return &DatastorePoolMG{poolCollection}
}

var _ models.DatastorePool = (*DatastorePoolMG)(nil)

func (ds DatastorePoolMG) Create(ctx context.Context, params *models.Pool) (*models.Pool, error) {
	_, err := ds.poolCollection.InsertOne(ctx, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (ds DatastorePoolMG) FindByID(ctx context.Context, id *string) (*models.Pool, error) {
	var pool *models.Pool
	query := bson.D{bson.E{Key: "pool_id", Value: id}}
	err := ds.poolCollection.FindOne(ctx, query).Decode(&pool)
	return pool, err
}

func (ds DatastorePoolMG) List(ctx context.Context) ([]*models.Pool, error) {
	var pools []*models.Pool
	cursor, err := ds.poolCollection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var pool models.Pool
		err := cursor.Decode(&pool)
		if err != nil {
			return nil, err
		}
		pools = append(pools, &pool)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	if len(pools) == 0 {
		return nil, errors.New("documents not found")
	}
	return pools, nil
}

func (ds DatastorePoolMG) Update(ctx context.Context, params *models.Pool) (*models.Pool, error) {
	filter := bson.D{primitive.E{Key: "pool_id", Value: params.PoolId}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "pool_id", Value: params.PoolId},
			primitive.E{Key: "total_pool_amount", Value: params.TotalPoolAmount},
			primitive.E{Key: "token_address", Value: params.TokenAddress},
			primitive.E{Key: "collection_name", Value: params.CollectionName},
			primitive.E{Key: "apy", Value: params.APY},
			primitive.E{Key: "duration", Value: params.Duration},
			primitive.E{Key: "state", Value: params.State},
			primitive.E{Key: "image", Value: params.Image},
		}}}

	var poolUpdated *models.Pool
	err := ds.poolCollection.FindOneAndUpdate(ctx, filter, update).Decode(&poolUpdated)
	if err != nil {
		return nil, errors.New("no matched document found for update")
	}

	return poolUpdated, nil
}

func (ds DatastorePoolMG) Delete(ctx context.Context, id *string) error {
	filter := bson.D{primitive.E{Key: "pool_id", Value: id}}
	result, _ := ds.poolCollection.DeleteOne(ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
