package data

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	Qid    primitive.ObjectID `json:"uuid" bson:"_id"`
	Uid    int                `json:"uid" bson:"uid"`
	Title  string             `json:"title" bson:"title"`
	Detail string             `json:"detail" bson:"detail"`
}
