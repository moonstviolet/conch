package data

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Question struct {
	Qid      int       `json:"qid" bson:"_id"`
	Uid      int       `json:"uid" bson:"uid"`
	Title    string    `json:"title" bson:"title"`
	Detail   string    `json:"detail" bson:"detail"`
	Follow   int       `json:"follow" bson:"follow"`
	Pageview int       `json:"pageview" bson:"pageview"`
	Lastmod  time.Time `json:"lastmod" bson:"lastmod"`
}

func (q *Question) Create() (err error) {
	questionColl := db.Collection("questions")
	_, err = questionColl.InsertOne(context.TODO(), q)
	return
}

func QuestionById(id int) (q Question, err error) {
	questionColl := db.Collection("questions")
	fliter := bson.M{
		"_id": id,
	}
	res := questionColl.FindOne(context.TODO(), fliter)
	res.Decode(&q)
	if err == nil && q.Qid == 0 {
		err = errors.New("Can find question")
	}
	return
}
