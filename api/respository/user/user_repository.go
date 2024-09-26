package repository_user

import (
	"my_wallet/api/entities"

	"github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
)

type UserRepository interface {
	CreateUser(user entities.User) (entities.User, error)
}

type MongoUserRepositoy struct {
	db     *mgo.Database
	logger logrus.FieldLogger
}

func NewMongoUserREpository(db *mgo.Database, logger logrus.FieldLogger) *MongoUserRepositoy {
	return &MongoUserRepositoy{
		db:     db,
		logger: logger,
	}
}

func (repo *MongoUserRepositoy) CreateUser(user entities.User) (entities.User, error) {
	err := repo.db.C("users").Insert(user)
	if err != nil {
		repo.logger.Errorln("Layer:user_repository", "Method:CreateUser", err)
		return user, err
	}
	repo.logger.Infoln("Layer:user_repository", "Method:CreateUser", "User:", user)
	return user, err
}
