package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/controllers"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/datastore"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/services"
	"github.com/gin-contrib/cors"
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

	marketItems  *mongo.Collection
	dsMarketItem *datastore.DatastoreMarketItemMG
	marketItemc  controllers.MarketItemController

	marketCollections  *mongo.Collection
	dsMarketCollection *datastore.DatastoreMarketCollectionMG
	marketCollectionc  controllers.MarketCollectionController

	checkins  *mongo.Collection
	dsCheckin *datastore.DatastoreCheckinMG
	checkinc  controllers.CheckinController

	sellers   *mongo.Collection
	dsSellers *datastore.DatastoreSellerMG
	sellerc   controllers.SellerController

	boxes   *mongo.Collection
	dsBoxes *datastore.DatastoreBoxMG
	boxc    controllers.BoxController

	projects   *mongo.Collection
	dsProjects *datastore.DatastoreProjectMG
	projectc   controllers.ProjectController

	packages   *mongo.Collection
	dsPackages *datastore.DatastorePackageMG
	packagec   controllers.PackageController

	participants   *mongo.Collection
	dsParticipants *datastore.DatastoreParticipantMG
	participantc   controllers.ParticipantController

	mongoclient *mongo.Client
	err         error
)

func init() {
	err := godotenv.Load()

	ctx = context.TODO()

	uri := os.Getenv("DB_URI")

	mongoconn := options.Client().ApplyURI(uri)
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
	loans = mongoclient.Database("hackathon").Collection("loans")
	lenders = mongoclient.Database("hackathon").Collection("lenders")
	borrowers = mongoclient.Database("hackathon").Collection("borrowers")
	marketItems = mongoclient.Database("hackathon").Collection("marketItems")
	marketCollections = mongoclient.Database("hackathon").Collection("marketCollections")
	checkins = mongoclient.Database("hackathon").Collection("checkins")
	sellers = mongoclient.Database("hackathon").Collection("sellers")
	boxes = mongoclient.Database("hackathon").Collection("boxes")
	projects = mongoclient.Database("hackathon").Collection("projects")
	packages = mongoclient.Database("hackathon").Collection("packages")
	participants = mongoclient.Database("hackathon").Collection("participants")

	dspools = datastore.NewDatastorePoolMG(pools, loans)
	dsloans = datastore.NewDatastoreLoanMG(loans)
	dslenders = datastore.NewDatastoreLenderMG(lenders)
	dsborrowers = datastore.NewDatastoreBorrowerMG(borrowers)
	dsMarketItem = datastore.NewDatastoreMarketItemMG(marketItems)
	dsMarketCollection = datastore.NewDatastoreMarketCollectionMG(marketCollections)
	dsCheckin = datastore.NewDatastoreCheckinMG(checkins)
	dsSellers = datastore.NewDatastoreSellerMG(sellers)
	dsBoxes = datastore.NewDatastoreBoxMG(boxes)
	dsProjects = datastore.NewDatastoreProjectMG(projects)
	dsPackages = datastore.NewDatastorePackageMG(packages)
	dsParticipants = datastore.NewDatastoreParticipantMG(participants)

	ps := services.NewPoolService(ctx, dspools)
	pc = controllers.NewPool(*ps)

	ls := services.NewLoanService(ctx, dsloans, dspools)
	lc = controllers.NewLoanController(*ls)

	lens := services.NewLenderService(ctx, dslenders)
	lenc = controllers.NewLenderController(*lens)

	bors := services.NewBorrowerService(ctx, dsborrowers)
	borc = controllers.NewBorrowerController(*bors)

	marketItemService := services.NewMarketItemService(ctx, dsMarketItem, dsMarketCollection)
	marketItemc = controllers.NewMarketItemController(*marketItemService)

	marketCollectionService := services.NewMarketCollectionService(ctx, dsMarketCollection)
	marketCollectionc = controllers.NewMarketCollectionController(*marketCollectionService)

	checkinService := services.NewCheckinService(ctx, dsCheckin)
	checkinc = controllers.NewCheckinController(*checkinService)

	sellersService := services.NewSellerService(ctx, dsSellers)
	sellerc = controllers.NewSellerController(*sellersService)

	boxService := services.NewBoxService(ctx, dsBoxes)
	boxc = controllers.NewBoxController(*boxService)

	projectService := services.NewProjectService(ctx, dsProjects)
	projectc = controllers.NewProjectController(*projectService)

	packageService := services.NewPackageService(ctx, dsPackages)
	packagec = controllers.NewPackageController(*packageService)

	participantService := services.NewParticipantService(ctx, dsParticipants, dsProjects)
	participantc = controllers.NewParticipantController(*participantService)
}

func main() {
	defer mongoclient.Disconnect(ctx)

	server = gin.Default()
	server.Use(cors.Default())
	basepath := server.Group("/v1")

	basepath.GET("/hello", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, nil) })
	uc.RegisterUserRoutes(basepath)
	pc.RegisterRoutes(basepath)
	lc.RegisterRoutes(basepath)
	lenc.RegisterRoutes(basepath)
	borc.RegisterRoutes(basepath)
	marketItemc.RegisterRoutes(basepath)
	marketCollectionc.RegisterRoutes(basepath)
	checkinc.RegisterRoutes(basepath)
	sellerc.RegisterRoutes(basepath)
	boxc.RegisterRoutes(basepath)
	projectc.RegisterRoutes(basepath)
	packagec.RegisterRoutes(basepath)
	participantc.RegisterRoutes(basepath)

	log.Fatal(server.Run())
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("ngrok-skip-browser-warning", "true")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")

		c.Next()
	}
}
