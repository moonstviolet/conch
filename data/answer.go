package data

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Answer struct {
	Aid     int       `json:"aid" bson:"_id"`
	Qid     int       `json:"qid" bson:"qid"`
	Uid     int       `json:"uid" bson:"uid"`
	Detail  string    `json:"detail" bson:"detail"`
	Lastmod time.Time `json:"lastmod" bson:"lastmod"`
}

func (a *Answer) Create() (err error) {
	answerColl := db.Collection("answers")
	_, err = answerColl.InsertOne(context.TODO(), a)
	return
}

func AnswerById(id int) (a Answer, err error) {
	answerColl := db.Collection("answers")
	fliter := bson.M{
		"_id": id,
	}
	res := answerColl.FindOne(context.TODO(), fliter)
	res.Decode(&a)
	if err == nil && a.Aid == 0 {
		err = errors.New("Can find question")
	}
	return
}
