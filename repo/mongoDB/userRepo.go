package mongodb

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/nilspolek/Workout-Tracker/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUserRepository struct {
	userCollection  *mongo.Collection
	tokenCollection *mongo.Collection
}

func NewUserRepository(uri string) (repo.UserRepositoryInterface, error) {
	var (
		err    error
		temp   = MongoUserRepository{}
		client *mongo.Client
	)
	clientOptions := options.Client().ApplyURI(uri)
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	temp.userCollection = client.Database("workout-tracker").Collection("users")
	temp.tokenCollection = client.Database("workout-tracker").Collection("tokens")
	return temp, err
}

func (r MongoUserRepository) GetUsers() ([]repo.User, error) {
	var (
		users = []repo.User{}
		ctx   = context.Background()
	)
	cur, err := r.userCollection.Find(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var user repo.User
		err = cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r MongoUserRepository) GetUser(username string) (repo.User, error) {
	var (
		err  error
		ctx  = context.Background()
		user repo.User
	)
	err = r.userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r MongoUserRepository) CreateUser(user repo.User) error {
	var (
		err error
		ctx = context.Background()
	)
	_, err = r.userCollection.InsertOne(ctx, user)
	return err
}

func (r MongoUserRepository) UpdateUser(user repo.User) error {
	var (
		err error
		ctx = context.Background()
	)
	_, err = r.userCollection.UpdateOne(ctx, repo.User{Username: user.Username}, user)
	return err
}

func (r MongoUserRepository) GenerateJWT(userID uuid.UUID) (string, error) {
	token, err := generateJWT(userID)
	if err != nil {
		return "", err
	}
	r.tokenCollection.InsertOne(context.Background(), repo.Token{UserID: userID.String(), Token: token, ExpiresAt: time.Now().Add(1 * time.Hour)})
	return token, nil
}

func (r MongoUserRepository) ValidateJWT(token string) (bool, error) {
	var (
		err error
		ctx = context.Background()
	)
	var tokenStruct repo.Token
	err = r.tokenCollection.FindOne(ctx, bson.M{"token": token}).Decode(&tokenStruct)
	if err == mongo.ErrNoDocuments {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if time.Now().After(tokenStruct.ExpiresAt) {
		return false, nil
	}
	return true, nil
}

func generateJWT(userID uuid.UUID) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   userID.String(),
		ExpiresAt: expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(repo.JwtKey)
}

func (r MongoUserRepository) Close() error {
	return r.userCollection.Database().Client().Disconnect(context.Background())
}
