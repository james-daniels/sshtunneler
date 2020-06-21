package env

import (
    "context"
		"fmt"
		"time"
		"go.mongodb.org/mongo-driver/bson"
		"go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
		"go.mongodb.org/mongo-driver/mongo/options"
)

//Document is a struct
type Document struct {
	ID primitive.ObjectID `bson:"_id, omitempty"`
	Name string
	IP string
	LocalPort string
	RemotePort string
	JumpServer string
	Path string
	Pause time.Duration
}

//DB is exported
func DB(env, host, db, coll string) Document {
	
	//setup client
	clientOptions := options.Client().ApplyURI("mongodb://" + host + ":27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	getError(err)

	//test connection
	err = client.Ping(context.TODO(), nil)
	getError(err)

	collection := client.Database(db).Collection(coll)

	var result Document

	err = collection.FindOne(context.TODO(), bson.M{"name": env}).Decode(&result)
	getError(err)

	return result
}

func getError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}