package server

import (
	"context"
	"my_wallet/api/endpoints"
	infraestructure_repository "my_wallet/api/respository/healtcheck"
	repository_user "my_wallet/api/respository/user"
	"my_wallet/api/services"
	infraestructure_services "my_wallet/api/services/healtcheck"
	transports "my_wallet/api/transports/http"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	dbMongo  *mongo.Client
	httpMux  *http.ServeMux
	httpAddr string
	logger   logrus.FieldLogger
}

func New(logger logrus.FieldLogger, httpAddr, dburl string, ctx context.Context) (*Server, error) {
	db := GetMongoDB(ctx, dburl)

	healtCheckRepository := infraestructure_repository.NewMongoUserREpository(db, logger)
	healtCheckService := infraestructure_services.NewHealtcheckService(ctx, healtCheckRepository, logger)
	userRepository := repository_user.NewMongoUserREpository(db, logger)
	userService := services.NewUserService(userRepository, logger, ctx)
	userEnpoints := endpoints.MakeServerEndpoints(userService, healtCheckService, logger)
	httpHandler := transports.NewHTTPHandler(userEnpoints, logger)

	httpMux := http.NewServeMux()
	httpMux.Handle("/", httpHandler)

	return &Server{
		dbMongo:  db,
		httpMux:  httpMux,
		httpAddr: httpAddr,
	}, nil
}

func (s *Server) Start() error {

	logrus.Infoln("Layel:Server ", " Method: Start", "Port:", s.httpAddr)
	if err := http.ListenAndServe(s.httpAddr, s.httpMux); err != nil {
		logrus.Fatalf("HTTP server failed: %v", err)
	}

	return nil
}

func GetMongoDB(ctx context.Context, dburl string) *mongo.Client {
	opts := options.Client().ApplyURI(dburl)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		logrus.Panic("Layel:Server Method:GetMongoDB ", err)
		os.Exit(2)
	}

	if err := client.Ping(ctx, nil); err != nil {
		logrus.Panic("Layel:Server Method:GetMongoDB Ping failed: ", err)
		os.Exit(2)
	}

	return client
}
