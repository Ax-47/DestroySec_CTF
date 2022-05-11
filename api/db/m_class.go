package db

import (
	"context"
	"fmt"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Db_mongo struct {
	url        string
	collection *mongo.Collection
}

func (db *Db_mongo) Db_start() {
	db.url = "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(db.url))
	if err != nil {
		fmt.Print(err)
	}
	ctx, cte := context.WithTimeout(context.Background(), 10*time.Second)
	defer cte()

	err = client.Connect(ctx)
	collection := client.Database("WEB").Collection("webidk")

	db.collection = collection

}
func (db Db_mongo) Db_InsertOne(Insert map[string]string) {

	_, err := db.collection.InsertOne(context.TODO(), Insert)
	if err != nil {
		fmt.Print(err)
	}
}

type DATA struct {
	Email   string `bson:"Email" json:"Email"`
	Subdata struct {
		Password, Username, Tag, UserId string
	}
}

func (db Db_mongo) Db_InsertOneS(Insert interface{}) {

	_, err := db.collection.InsertOne(context.TODO(), Insert)
	if err != nil {
		fmt.Print(err)
	}
}
func (db Db_mongo) Db_FindtOne(dfkdf string, Username string) primitive.D {
	var result bson.D
	f := bson.D{{dfkdf, Username}}
	coll := db.collection
	err := coll.FindOne(context.TODO(), f).Decode(&result)
	if err != nil {
		fmt.Print(err)
	}

	return result
}
func (db Db_mongo) Db_FindALL(dfkdf string, something string) ([]primitive.M, error) {

	f := bson.D{{dfkdf, something}}
	coll := db.collection
	axc, err := coll.Find(context.TODO(), f)
	var results []bson.M
	if err != nil {
		fmt.Print(err)
	}
	if err = axc.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	var result bson.M
	if err = coll.FindOne(context.TODO(), f).Decode(&result); err != nil {
		return nil, err
	}
	return results, nil
}
