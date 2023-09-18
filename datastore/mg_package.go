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

type DatastorePackageMG struct {
	packageCollection *mongo.Collection
}

func NewDatastorePackageMG(packageCollection *mongo.Collection) *DatastorePackageMG {
	return &DatastorePackageMG{packageCollection}
}

var _ models.DatastorePackage = (*DatastorePackageMG)(nil)

func (ds DatastorePackageMG) Create(ctx context.Context, params *models.Package) (*models.Package, error) {
	count, err := ds.packageCollection.CountDocuments(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	params.IsActive = true
	params.PackageId = int(count) + 1

	_, err = ds.packageCollection.InsertOne(ctx, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (ds DatastorePackageMG) FindByID(ctx context.Context, id *string) (*models.Package, error) {
	var pkg *models.Package
	idInt, err := strconv.Atoi(*id)
	if err != nil {
		return nil, err
	}

	query := bson.D{bson.E{Key: "package_id", Value: idInt}}
	err = ds.packageCollection.FindOne(ctx, query).Decode(&pkg)
	return pkg, err
}

func (ds DatastorePackageMG) List(ctx context.Context) ([]*models.Package, error) {
	var pkgs []*models.Package
	cursor, err := ds.packageCollection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var pkg models.Package
		err := cursor.Decode(&pkg)
		if err != nil {
			return nil, err
		}
		if pkg.IsActive == true {
			pkgs = append(pkgs, &pkg)
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return pkgs, nil
}

func (ds DatastorePackageMG) Update(ctx context.Context, params *models.Package) (*models.Package, error) {
	var pkg *models.Package
	query := bson.D{bson.E{Key: "package_id", Value: params.PackageId}}
	err := ds.packageCollection.FindOne(ctx, query).Decode(&pkg)
	if err != nil {
		return nil, err
	}

	if params.ProjectId != nil {
		pkg.ProjectId = params.ProjectId
	}

	if params.PackageName != nil {
		pkg.PackageName = params.PackageName
	}

	if params.PackageDescription != nil {
		pkg.PackageDescription = params.PackageDescription
	}

	if params.PackagePrice != nil {
		pkg.PackagePrice = params.PackagePrice
	}

	if params.ProjectAddress != nil {
		pkg.ProjectAddress = params.ProjectAddress
	}

	filter := bson.D{primitive.E{Key: "package_id", Value: params.PackageId}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "project_id", Value: pkg.ProjectId},
		primitive.E{Key: "package_price", Value: pkg.PackageName},
		primitive.E{Key: "package_description", Value: pkg.PackageDescription},
		primitive.E{Key: "package_price", Value: pkg.PackagePrice},
	}}}

	var updatedPkg *models.Package
	err = ds.packageCollection.FindOneAndUpdate(ctx, filter, update).Decode(&updatedPkg)
	if err != nil {
		return nil, errors.New("no matched document found for update")
	}

	return params, nil
}

func (ds DatastorePackageMG) Delete(ctx context.Context, id *string) error {
	idInt, err := strconv.Atoi(*id)
	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "package_id", Value: idInt}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "package_id", Value: idInt},
			primitive.E{Key: "is_active", Value: false},
		}}}

	result, _ := ds.packageCollection.UpdateOne(ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
