package repository_user

import (
	"context"
	"my_wallet/api/entities"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(user entities.User, ctx context.Context) (entities.User, error)
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

func (repo *MongoUserRepositoy) CreateUser(user entities.User, ctx context.Context) (entities.User, error) {
	coll := repo.db.Database("mywallet").Collection("users")
	result, err := coll.InsertOne(ctx, user)
	repo.logger.Infoln("Layer:user repository ", "Method:user_repository ", "result:", result.InsertedID)
	user.ID = result.InsertedID.(primitive.ObjectID).Hex()
	if err != nil {
		repo.logger.Errorln("Layer:user_repository", "Method:CreateUser", err)
		return user, err
	}
	repo.logger.Infoln("Layer:user_repository", "Method:CreateUser", "User:", user)
	return user, err
}
