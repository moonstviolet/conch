package data

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Counter struct {
	Name  string `bson:"_id"`
	Count int    `bson:"count"`
}

var db *mongo.Database

func init() {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	db = client.Database("conch")
}

func AutoIncrement(name string) int {
	counters := db.Collection("counters")
	fliter := bson.M{
		"_id": name,
	}
	res := Counter{}
	cursor := counters.FindOne(context.TODO(), fliter)
	cursor.Decode(&res)
	if res.Count == 0 {
		res.Name = name
		res.Count++
		_, err := counters.InsertOne(context.TODO(), res)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		res.Count++
		update := bson.M{
			"$set": bson.M{
				"count": res.Count,
			},
		}
		_, err := counters.UpdateOne(context.TODO(), fliter, update)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return res.Count
}

func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}
