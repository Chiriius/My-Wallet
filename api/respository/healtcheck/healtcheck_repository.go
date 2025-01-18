package infraestructure_repository

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type HealtCheckRepository interface {
	GetHealtcheck(ctx context.Context) (bool, error)
}

type MongoUserRepositoy struct {
	db     *mongo.Client
	logger logrus.FieldLogger
}

func NewMongoUserREpository(db *mongo.Client, logger logrus.FieldLogger) *MongoUserRepositoy {
	return &MongoUserRepositoy{
		db:     db,
		logger: logger,
	}
}

func (repo *MongoUserRepositoy) GetHealtcheck(ctx context.Context) (bool, error) {
	repoCtx, c := context.WithTimeout(ctx, 10*time.Second)
	err := repo.db.Ping(repoCtx, nil)
	if err != nil {
		repo.logger.Errorln("Layer: infraestructure_repository ", "Method:GetHealtcheck ", "Error:", c)
		return false, ErrLoadingDatabase
	}
	return true, nil

}
