/*******************************************************************/
/*  															   */
/*  @project     : WebHook webserver							   */
/*  @package     :   										       */
/*  @subpackage  : connectDB									   */
/*  @access      :												   */
/*  @paramtype   : 												   */
/*  @argument    :												   */
/*  @description : Functions to connect MongoDB database		   */
/*                                                                 */
/*				                                                   */
/*																   */
/*  @author Emmanuel COLUSSI									   */
/*  @version 1.00												   */
/******************************************************************/

package connectDB

import (
	"context"
	"log"
	"strconv"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var clientInstance *mongo.Client
var databaseInstance *mongo.Database

var clientInstanceError error

var mongoOnce sync.Once

// Declare a Const for mongodb connexion
const (
	CONNECTIONSTRING = "mongodb://repmonitor:Demos2022@localhost:27017/?authMechanism=SCRAM-SHA-256&authSource=repmonitor"
	DB               = "repmonitor"
	ISSUES           = "loggithub"
)

type Logmessage struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Org        string             `json:"org" bson:"org"`
	PusherName string             `json:"pushername" bson:"pushername"`
	PusherLink string             `json:"pusherlink" bson:"pusherlink"`
	ActionHook string             `json:"actionhook" bson:"actionhook"`
	Repos      string             `json:"repos" bson:"repos"`
	DateEvt    time.Time          `json:"dateevt" bson:"dateevt"`
}

//GetMongoClient - Return mongodb connection to work with
func GetMongoClient() (*mongo.Database, error) {
	//Perform connection creation operation only once.
	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		//client, err := mongo.NewClient(options.Client().ApplyURI(CONNECTIONSTRING))
		if err != nil {
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
		databaseInstance = client.Database(DB)
	})
	return databaseInstance, clientInstanceError
}

func GetCollectionAll(Coll string) ([]bson.M, error) {

	databaseInstance, err := GetMongoClient()
	if err != nil {
		clientInstanceError = err
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	//defer client.Disconnect(ctx)

	regionCollection := databaseInstance.Collection(Coll)

	cursor, err := regionCollection.Find(ctx, bson.M{})
	if err != nil {
		clientInstanceError = err
	}

	// Return ALL
	var collection []bson.M

	if err = cursor.All(ctx, &collection); err != nil {
		clientInstanceError = err
	}

	// Close the cursor once finished
	cursor.Close(ctx)

	//defer client.Disconnect(ctx)
	return collection, clientInstanceError

}

func GetCountDoc(Coll string, req bson.M) (int64, error) {
	databaseInstance, err := GetMongoClient()
	if err != nil {
		clientInstanceError = err
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	regionCollection := databaseInstance.Collection(Coll)

	nbrdoc, err := regionCollection.CountDocuments(ctx, req)
	if err != nil {
		clientInstanceError = err
	}

	return nbrdoc, clientInstanceError

}

func GetReqCollectionAll(Coll string, req bson.M) ([]bson.M, error) {

	databaseInstance, err := GetMongoClient()
	if err != nil {
		clientInstanceError = err
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	regionCollection := databaseInstance.Collection(Coll)

	//cursor, err := regionCollection.Find(ctx, Request)
	cursor, err := regionCollection.Find(ctx, req)
	if err != nil {
		clientInstanceError = err
	}

	// Return ALL
	var collection1 []bson.M

	if err = cursor.All(ctx, &collection1); err != nil {
		clientInstanceError = err
	}

	// Close the cursor once finished
	cursor.Close(ctx)

	//defer client.Disconnect(ctx)
	return collection1, clientInstanceError

}
func SearchDist(Coll string, search string) ([]bson.M, error) {

	var collection1 []bson.M

	databaseInstance, err := GetMongoClient()
	if err != nil {
		clientInstanceError = err
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	regionCollection := databaseInstance.Collection(Coll)

	//patternName := `.*` + search + `*.`
	patternName := search + `*.`
	cursor, err := regionCollection.Find(ctx, bson.M{"name": primitive.Regex{Pattern: patternName, Options: "i"}})
	if err != nil {
		clientInstanceError = err
	}

	if err = cursor.All(ctx, &collection1); err != nil {
		clientInstanceError = err
	}
	// Close the cursor once finished
	cursor.Close(ctx)
	return collection1, clientInstanceError

}

func InsertCollection(Coll string, InsertD interface{}) (*mongo.InsertOneResult, error) {

	databaseInstance, err := GetMongoClient()
	if err != nil {
		clientInstanceError = err
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	regionCollection := databaseInstance.Collection(Coll)

	result, insertErr := regionCollection.InsertOne(ctx, InsertD)

	return result, insertErr

}

func RemoveCollection(Coll string, IDDist string) (*mongo.DeleteResult, error) {

	databaseInstance, err := GetMongoClient()
	if err != nil {
		clientInstanceError = err
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	regionCollection := databaseInstance.Collection(Coll)
	IDDist1, _ := strconv.Atoi(IDDist)

	result, err := regionCollection.DeleteOne(ctx, bson.M{"distillerID": IDDist1})

	if err != nil {
		log.Fatal(err)
	}

	return result, nil

}

func UpdateCollection(Coll string, IDDist int, request bson.M) (*mongo.UpdateResult, error) {

	databaseInstance, err := GetMongoClient()
	if err != nil {
		clientInstanceError = err
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	regionCollection := databaseInstance.Collection(Coll)

	update, err := regionCollection.UpdateOne(ctx, bson.M{"distillerID": IDDist}, request)

	if err != nil {
		log.Fatal(err)
	}
	return update, nil

}
