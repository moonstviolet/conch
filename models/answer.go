package models

import (
	"context"
	"errors"
	"log"
	"sort"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Answer struct {
	Aid      int       `json:"aid" bson:"_id"`
	Qid      int       `json:"qid" bson:"qid"`
	Uid      int       `json:"uid" bson:"uid"`
	Username string    `json:"username" bson:"username"`
	Detail   string    `json:"detail" bson:"detail"`
	Lastmod  time.Time `json:"lastmod" bson:"lastmod"`
}

type aList []Answer

func (al aList) Len() int {
	return len(al)
}

func (al aList) Less(i, j int) bool {
	return al[i].Lastmod.After(al[j].Lastmod)
}

func (al aList) Swap(i, j int) {
	al[i], al[j] = al[j], al[i]
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
	if err = res.Decode(&a); err == nil && a.Aid == 0 {
		err = errors.New("Can find question")
	}
	return
}

func AnswersByQid(qid int) (alist []Answer, err error) {
	answerColl := db.Collection("answers")
	cursor, err := answerColl.Find(context.TODO(), bson.M{
		"qid": qid,
	})
	if err != nil {
		log.Println(err)
		return
	}
	err = cursor.All(context.TODO(), &alist)
	sort.Sort(aList(alist))
	return
}

func AnswersByUid(uid int) (alist []Answer, err error) {
	answerColl := db.Collection("answers")
	cursor, err := answerColl.Find(context.TODO(), bson.M{
		"uid": uid,
	})
	if err != nil {
		log.Println(err)
		return
	}
	err = cursor.All(context.TODO(), &alist)
	sort.Sort(aList(alist))
	return
}
