package data

import (
	"context"
	"time"

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

type Session struct {
	Sid        string    `json:"sid" bson:"_id"`
	Uid        int       `json:"uid" bson:"uid"`
	CreateTime time.Time `json:"createTime" bson:"createTime"`
}

func (user *User) Create() (err error) {
	userColl := db.Collection("users")
	_, err = userColl.InsertOne(context.TODO(), user)
	return
}

func (user *User) Delete() (err error) {
	userColl := db.Collection("users")
	_, err = userColl.DeleteOne(context.TODO(), user)
	return
}

func UserDeleteAll() (err error) {
	userColl := db.Collection("users")
	_, err = userColl.DeleteMany(context.TODO(), bson.M{})
	return
}

func (s *Session) Check() (valid bool) {
	sessionColl := db.Collection("sessions")
	res := sessionColl.FindOne(context.TODO(), s)
	var t Session
	res.Decode(&t)
	return s.Sid == t.Sid && s.Uid == t.Uid
}

func (s *Session) User() (user User) {
	userColl := db.Collection("users")
	fliter := bson.M{
		"uid": s.Uid,
	}
	res := userColl.FindOne(context.TODO(), fliter)
	res.Decode(&user)
	return
}
