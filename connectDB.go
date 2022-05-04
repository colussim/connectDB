/*******************************************************************/
/*  															   */
/*  @project     : WebHook webserver							   */
/*  @package     :   										       */
/*  @subpackage  : connectDB									   */
/*  @access      :												   */
/*  @paramtype   : 	CONNECTIONSTRING,DB,ISSUES				       */
/*  @argument    :												   */
/*  @description : Functions to connect MongoDB database		   */
/*                 and List, ADD, Delete                                                 */
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

type Logmessage struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Org        string             `json:"org" bson:"org"`
	PusherName string             `json:"pushername" bson:"pushername"`
	PusherLink string             `json:"pusherlink" bson:"pusherlink"`
	ActionHook string             `json:"actionhook" bson:"actionhook"`
	Repos      string             `json:"repos" bson:"repos"`
	DateEvt    time.Time          `json:"dateevt" bson:"dateevt"`
}

//GetMongoClient : Return mongodb connection to work with
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

//GetCollectionAll : Reading All Documents from a Collection
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

// GetCountDoc : Return number of document in Collection

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

// GetReqCollectionAll : Return Documents from a Collection with a Filter
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

// InsertCollection : Insert Documents in a Collection
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

// RemoveCollection : Remove Documents in a Collection
func RemoveCollection(Coll string, IDDist string) (*mongo.DeleteResult, error) {

	databaseInstance, err := GetMongoClient()
	if err != nil {
		clientInstanceError = err
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	regionCollection := databaseInstance.Collection(Coll)
	IDDist1, _ := strconv.Atoi(IDDist)

	result, err := regionCollection.DeleteOne(ctx, bson.M{"_id": IDDist1})

	if err != nil {
		log.Fatal(err)
	}

	return result, nil

}

//UpdateCollection : Update Documents in a Collection
func UpdateCollection(Coll string, IDDist int, request bson.M) (*mongo.UpdateResult, error) {

	databaseInstance, err := GetMongoClient()
	if err != nil {
		clientInstanceError = err
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	regionCollection := databaseInstance.Collection(Coll)

	update, err := regionCollection.UpdateOne(ctx, bson.M{"_id": IDDist}, request)

	if err != nil {
		log.Fatal(err)
	}
	return update, nil

}

//RemoveAllCollection : Remove All Documents in a Collection
func RemoveAllCollection(Coll string) (*mongo.DeleteResult, error) {

	databaseInstance, err := GetMongoClient()
	if err != nil {
		clientInstanceError = err
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	regionCollection := databaseInstance.Collection(Coll)

	result, err := regionCollection.DeleteMany(ctx, bson.D{{}})

	regionCollection.DeleteMany(ctx, bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	return result, nil

}
