package data

import (
	"context"
	"errors"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Uuid     primitive.ObjectID `json:"uuid" bson:"_id"`
	Uid      int                `json:"uid" bson:"uid"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Nickname string             `json:"nickname" bson:"nickname"`
	Motto    string             `json:"motto" bson:"motto"`
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

func UserByUsername(name string) (user User, err error) {
	userColl := db.Collection("users")
	fliter := bson.M{
		"username": name,
	}
	res := userColl.FindOne(context.TODO(), fliter)
	res.Decode(&user)
	if err == nil && user.Uid == 0 {
		err = errors.New("Can find user")
	}
	return
}

func (user *User) CreateSession() (session Session, err error) {
	sessionColl := db.Collection("sessions")
	session = Session{
		Sid:        createUUID(),
		Uid:        user.Uid,
		CreateTime: time.Now(),
	}
	_, err = sessionColl.InsertOne(context.TODO(), session)
	return
}

func CheckSession(w http.ResponseWriter, r *http.Request) (s Session, err error) {
	cookie, err := r.Cookie("session")
	if err == nil {
		sessionColl := db.Collection("sessions")
		fliter := bson.M{
			"_id": cookie.Value,
		}
		res := sessionColl.FindOne(context.TODO(), fliter)
		err = res.Decode(&s)
		if err == nil && s.Sid == "" {
			err = errors.New("Invalid session")
		}
	}
	return
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
