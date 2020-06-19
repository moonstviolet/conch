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
	Pageview int       `json:"pageview" bson:"lastmod"`
	Lastmod  time.Time `json:"lastmod" bson:"follow"`
}

func (question *Question) Create() (err error) {
	questionColl := db.Collection("questions")
	_, err = questionColl.InsertOne(context.TODO(), question)
	return
}

func QuestionById(id int) (question Question, err error) {
	questionColl := db.Collection("questions")
	fliter := bson.M{
		"_id": id,
	}
	res := questionColl.FindOne(context.TODO(), fliter)
	res.Decode(&question)
	if err == nil && question.Qid == 0 {
		err = errors.New("Can find question")
	}
	return
}
