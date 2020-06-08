package data

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Uuid     primitive.ObjectID `json:"uuid" bson:"_id"`
	Uid      int                `json:"uid" bson:"uid"`
	Username string             `json:"username" bson:"useranme"`
	Password string             `json:"password" bson:"password"`
	Nickname string             `json:"nickname" bson:"nickname"`
}

func (user *User) Create() (err error) {
	userList := db.Collection("users")
	_, err = userList.InsertOne(context.TODO(), user)
	return
}

func (user *User) Delete() (err error) {
	userList := db.Collection("users")
	_, err = userList.DeleteOne(context.TODO(), user)
	return
}

func UserDeleteAll() (err error) {
	userList := db.Collection("users")
	_, err = userList.DeleteMany(context.TODO(), bson.M{})
	return
}
