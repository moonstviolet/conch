package models

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Uid      int    `json:"uid" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
	Avatar   string `json:"avatar" bson:"avatar"`
	Nickname string `json:"nickname" bson:"nickname"`
	Motto    string `json:"motto" bson:"motto"`
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
	if err = res.Decode(&user); err == nil && user.Uid == 0 {
		err = errors.New("Can find user")
	}
	return
}

func (user *User) EmailHash() {
	s := fmt.Sprintf("%x", md5.Sum([]byte(user.Email)))
	user.Avatar = "https://www.gravatar.com/avatar/" + s
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

func CheckSession(sid string) (s Session, err error) {
	sessionColl := db.Collection("sessions")
	fliter := bson.M{
		"_id": sid,
	}
	res := sessionColl.FindOne(context.TODO(), fliter)
	err = res.Decode(&s)
	if err == nil && s.Sid == "" {
		err = errors.New("Invalid session")
	}
	return
}

func (s *Session) User() (user User) {
	userColl := db.Collection("users")
	fliter := bson.M{
		"_id": s.Uid,
	}
	res := userColl.FindOne(context.TODO(), fliter)
	_ = res.Decode(&user)
	return
}

func (s *Session) DeleteBySid() (err error) {
	sessionColl := db.Collection("sessions")
	fliter := bson.M{
		"_id": s.Sid,
	}
	_, err = sessionColl.DeleteOne(context.TODO(), fliter)
	return
}
