package repository_user

import (
	"context"
	"my_wallet/api/entities"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	CreateUser(user entities.User, ctx context.Context) (entities.User, error)
	GetUser(id string, ctx context.Context) (entities.User, error)
	GetUserByEmail(email string, ctx context.Context) (entities.User, error)
	DeleteUser(id string, ctx context.Context) error
	UpdateUser(userUpr entities.User, ctx context.Context) (entities.User, error)
	SoftDeleteUser(id string, ctx context.Context) error
	UpdateUserToken(userUpr entities.User, ctx context.Context) (entities.User, error)
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
		repo.logger.Errorln("Layer:user_repository ", "Method:CreateUser ", "Error:", err)
		return user, err
	}
	repo.logger.Infoln("Layer:user_repository ", "Method:CreateUser ", "User:", user)
	return user, err
}

func (repo *MongoUserRepositoy) GetUser(id string, ctx context.Context) (entities.User, error) {
	var user entities.User
	idd, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, ErrUserNotfound
	}

	filter := bson.D{{"_id", idd}}
	opts := options.FindOne()
	coll := repo.db.Database("mywallet").Collection("users")

	err = coll.FindOne(ctx, filter, opts).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, ErrUserNotfound
		}
		return user, err
	}
	if user.Enabled != true {
		return entities.User{}, ErrDisbledUser
	}
	return user, nil
}

func (repo *MongoUserRepositoy) GetUserByEmail(email string, ctx context.Context) (entities.User, error) {
	var user entities.User
	filter := bson.D{{"email", email}}
	opts := options.FindOne()
	coll := repo.db.Database("mywallet").Collection("users")

	err := coll.FindOne(ctx, filter, opts).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, ErrUserNotfound
		}
		return user, err
	}
	if user.Enabled != true {
		return entities.User{}, ErrDisbledUser
	}
	return user, nil
}

func (repo *MongoUserRepositoy) UpdateUser(userUpr entities.User, ctx context.Context) (entities.User, error) {
	ide := string(userUpr.ID)
	idd, err := primitive.ObjectIDFromHex(ide)
	if err != nil {
		repo.logger.Errorln("Layer:user_repository ", "Method:UpdateUser ", "Error:", err)
		return entities.User{}, err
	}

	filter := bson.D{{"_id", idd}}
	coll := repo.db.Database("mywallet").Collection("users")
	userUpdate := bson.M{
		"$set": bson.M{
			"typedni":  userUpr.TypeDNI,
			"name":     userUpr.Name,
			"email":    userUpr.Email,
			"password": userUpr.Password,
			"address":  userUpr.Address,
			"phone":    userUpr.Phone,
			"enabled":  userUpr.Enabled,
		},
	}

	_, err = coll.UpdateOne(ctx, filter, userUpdate)
	if err != nil {
		return entities.User{}, err
	}
	repo.logger.Infoln("Layer:user_repository ", "Method:UpdateUser ", "User:", userUpr)
	return userUpr, nil
}
func (repo *MongoUserRepositoy) UpdateUserToken(userUpr entities.User, ctx context.Context) (entities.User, error) {
	filter := bson.D{{"dni", userUpr.DNI}}
	coll := repo.db.Database("mywallet").Collection("users")
	userUpdate := bson.M{
		"$set": bson.M{
			"token":      userUpr.Token,
			"updated_at": userUpr.Update_at,
		},
	}

	_, err := coll.UpdateOne(ctx, filter, userUpdate)
	if err != nil {
		return entities.User{}, err
	}
	repo.logger.Infoln("Layer:user_repository ", "Method:UpdateUserToken ", "User:", userUpr)
	return userUpr, nil
}

func (repo *MongoUserRepositoy) SoftDeleteUser(id string, ctx context.Context) error {
	idd, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", idd}}
	coll := repo.db.Database("mywallet").Collection("users")
	userUpdate := bson.M{
		"$set": bson.M{
			"enabled": false,
		},
	}

	_, err = coll.UpdateOne(ctx, filter, userUpdate)
	if err != nil {
		return err
	}
	repo.logger.Infoln("Layer:user_repository ", "Method: SoftDeleteUser ", "User:", idd)
	return nil
}

func (repo *MongoUserRepositoy) DeleteUser(id string, ctx context.Context) error {
	coll := repo.db.Database("mywallet").Collection("users")
	idd, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		repo.logger.Errorln("Layer:user_repository ", "Method: DeleteUser ", "Error:", err)
		return err
	}

	filter := bson.D{{"_id", idd}}
	res, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		repo.logger.Errorln("Layer:user_repository ", "Method: DeleteUser ", "Error:", err)

		return err
	}

	if res.DeletedCount == 0 {
		repo.logger.Errorln("Layer:user_repository ", "Method: DeleteUser ", "Error: No tasks were deleted")
		return ErrNotasks
	}
	repo.logger.Infoln("Layer:user_repository ", "Method: DeleteUser ", "User:", idd)
	return nil

}
