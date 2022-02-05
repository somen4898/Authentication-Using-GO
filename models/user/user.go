package user

import (
	"context"
	"go-auth/database/mongoDB"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserForCredential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Name     string `json:"name"     bson:"name"`
	Gender   string `json:"gender"   bson:"gender"`
	Age      int    `json:"age"      bson:"age"`
	Email    string `json:"email"    bson:"email"`
	Password string `json:"password" bson:"password"`
}

func (u *User) Create() (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id, err := mongoDB.UserCollection.InsertOne(ctx, u)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func GetOneByEmail(email string) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result := mongoDB.UserCollection.FindOne(ctx, bson.M{"email": email})
	user := User{}
	err := result.Decode(&user)
	return user, err
}
