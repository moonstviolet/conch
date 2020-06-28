package data

import "context"

type Qfollow struct {
	Qid int `json:"qid" bson:"qid"`
	Uid int `json:"uid" bson:"uid"`
}

func (qf *Qfollow) Create() (err error) {
	qfollowColl := db.Collection("qfollows")
	_, err = qfollowColl.InsertOne(context.TODO(), qf)
	return
}