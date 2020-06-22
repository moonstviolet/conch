package data

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Question struct {
	Qid      int       `json:"qid" bson:"_id"`
	Uid      int       `json:"uid" bson:"uid"`
	Title    string    `json:"title" bson:"title"`
	Detail   string    `json:"detail" bson:"detail"`
	ModUser  string    `json:"moduser" bson:"moduser"`
	Modify   string    `json:"modify" bson:"modify"`
	Follow   int       `json:"follow" bson:"follow"`
	Pageview int       `json:"pageview" bson:"pageview"`
	Lastmod  time.Time `json:"lastmod" bson:"lastmod"`
}

func (q *Question) Create() (err error) {
	questionColl := db.Collection("questions")
	_, err = questionColl.InsertOne(context.TODO(), q)
	return
}

func QuestionById(qid int) (q Question, err error) {
	questionColl := db.Collection("questions")
	filter := bson.M{
		"_id": qid,
	}
	res := questionColl.FindOne(context.TODO(), filter)
	res.Decode(&q)
	if err == nil && q.Qid == 0 {
		err = errors.New("Can find question")
	}
	return
}

func (q *Question) Update() (err error) {
	questionColl := db.Collection("questions")
	filter := bson.M{
		"_id": q.Qid,
	}
	update := bson.M{
		"$set": q,
	}
	_, err = questionColl.UpdateOne(context.TODO(), filter, update)
	return
}

func QuestionList() (qlist []Question, err error) {
	questionColl := db.Collection("questions")
	cursor, err := questionColl.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err)
		return
	}
	err = cursor.All(context.TODO(), &qlist)
	return
}

func QuestionsByUser(uid int) (qlist []Question, err error) {
	return
}
