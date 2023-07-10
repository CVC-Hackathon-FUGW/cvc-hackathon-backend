package main

import (
	"context"
	"fmt"
	"log"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/controllers"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/datastore"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server *gin.Engine

	us    services.UserService
	uc    controllers.UserController
	ctx   context.Context
	userc *mongo.Collection

	pools   *mongo.Collection
	dspools *datastore.DatastorePoolMG
	pc      controllers.PoolController

	loans   *mongo.Collection
	dsloans *datastore.DatastoreLoanMG
	lc      controllers.LoanController

	lenders   *mongo.Collection
	dslenders *datastore.DatastoreLenderMG
	lenc      controllers.LenderController

	borrowers   *mongo.Collection
	dsborrowers *datastore.DatastoreBorrowerMG
	borc        controllers.BorrowerController

	mongoclient *mongo.Client
	err         error
)

func init() {
	err := godotenv.Load()

	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	userc = mongoclient.Database("hackathon").Collection("users")
	us = services.NewUserService(userc, ctx)
	uc = controllers.NewUser(us)

	pools = mongoclient.Database("hackathon").Collection("pools")
	dspools = datastore.NewDatastorePoolMG(pools)
	ps := services.NewPoolService(ctx, dspools)
	pc = controllers.NewPool(*ps)

	loans = mongoclient.Database("hackathon").Collection("loans")
	dsloans = datastore.NewDatastoreLoanMG(loans)
	ls := services.NewLoanService(ctx, dsloans)
	lc = controllers.NewLoanController(*ls)

	lenders = mongoclient.Database("hackathon").Collection("lenders")
	dslenders = datastore.NewDatastoreLenderMG(lenders)
	lens := services.NewLenderService(ctx, dslenders)
	lenc = controllers.NewLenderController(*lens)

	borrowers = mongoclient.Database("hackathon").Collection("borrowers")
	dsborrowers = datastore.NewDatastoreBorrowerMG(borrowers)
	bors := services.NewBorrowerService(ctx, dsborrowers)
	borc = controllers.NewBorrowerController(*bors)

	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	uc.RegisterUserRoutes(basepath)
	pc.RegisterRoutes(basepath)
	lc.RegisterRoutes(basepath)
	lenc.RegisterRoutes(basepath)
	borc.RegisterRoutes(basepath)
	log.Fatal(server.Run(":9090"))
}
