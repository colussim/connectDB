# connectDB
This is a simple go package for a MongoDB database that I needed for a project.

This package allows :
* Connect to MongoDB
* Reading All Documents from a Collection
* Return number of document in Collection
* Return Documents from a Collection with a Filter
* Insert Documents in a Collection
* Update Documents in a Collection
* Delete Documents in a Collection

## Prerequisites

Before you get started, you’ll need to have these things:
* The official [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver/) 
* The Package bson is an implementation of the BSON specification for GO : [mgo.v2/bson](https://gopkg.in/mgo.v2/bson)

## Initial setup

Install the MongoDB Go Driver :
```
go install go.mongodb.org/mongo-driver
```
Install the bson Go Driver :
```
go install gopkg.in/mgo.v2/bson
```
Install the connectDB package :
```
go install github.com/colussim/connectDB
```

## Usage

To use this module you must initialize 3 variables:
* var CONNECTIONSTRING = Connection String URI :
  
*mongodb://[username:password@]host1[:port1][/[defaultauthdb][?options]]*
* var DB = *Name of the database*
* var ISSUES = *Name of the collection*

### Functions

**GetMongoClient()** : *Return mongodb connection (mongo.Client),use variable CONNECTIONSTRING*
```go
var CONNECTIONSTRING = "mongodb://user:password@localhost:27017/?authMechanism=SCRAM-SHA-256&authSource=repmonitor"
databaseInstance, err := GetMongoClient()
```
  
**GetCollectionAll(Coll string) ([]bson.M, error)** : *Reading All Documents from a Collection*
```go
Collection := "log"
EventLogAll, err := connectDB.GetCollectionAll(Collection)
if err != nil {
		log.Fatal(err)
}
```

**GetCountDoc(Coll string, req bson.M) (int64, error)** : *Return number of document in Collection*
```go
var Request1 bson.M
Collection := "log"
CollectionLogID := "62671f794c9950ae8189db65"

ExistDoc, err := connectiondb.GetCountDoc(Collection, Request1)
if err != nil {
	log.Fatal(err)
}
```

**GetReqCollectionAll(Coll string, req bson.M) ([]bson.M, error)** : *Return Documents from a Collection with a Filter*
```go
Collection := "log"
CollectionLogID := "62671f794c9950ae8189db65"
Request0 := bson.M{"_id": CollectionLogID}

DistillRegion, err := connectiondb.GetReqCollectionAll(Collection, Request0)
if err != nil {
	log.Fatal(err)
}
```

**InsertCollection(Coll string, InsertD interface{}) (mongo.InsertOneResult, error)** : *Insert Documents in a Collection*
```go
type Logmessage1 struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Org        string             `json:"org" bson:"org"`
	PusherName string             `json:"pushername" bson:"pushername"`
	DateEvt    time.Time          `json:"dateevt" bson:"dateevt"`
	Messages   string             `json:"messages" bson:"messages"`
}

Collection := "log"
MessageLog := Logmessage1{
		ID:         primitive.NewObjectID(),
		Org:        org,
		PusherName: sender,
		DateEvt:    time.Now(),
		Messages:   "Alert : " ,
	}

	_, insertErr := connectDB.InsertCollection("Collection", MessageLog)
	if insertErr != nil {
		log.Println("⇨ Problem Event not insert in database")

	} else {

		log.Println("⇨ New event: insert in database")
	}
```

**RemoveCollection(Coll string, IDDist string) (mongo.DeleteResult, error)** : *Remove Documents in a Collection*
```go
Collection := "log"
CollectionLogID := "62671f794c9950ae8189db65"

result, err := connectiondb.RemoveCollection(Collection, CollectionLogID)
		
if err != nil {
    log.Println("⇨ Error deleted")
	} else {
			log.Println("⇨ Record deleted")
		}
```

**UpdateCollection(Coll string, IDDist int, request bson.M) (mongo.UpdateResult, error)** : *Update Documents in a Collection*
```go
Collection := "log"
CollectionLogID := "62671f794c9950ae8189db65"
org :="test"
name :="test"

update := bson.M{"$set": bson.M{"org": org, "name": name}}

_, updateErr := connectiondb.UpdateCollection(Collection, CollectionLogID, update)
if updateErr != nil {
	log.Println("⇨ Record not updated")
} else {
	log.Println("⇨ Record updated")		
		}
```

## Resources :

[MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver/)
[Documentation MongoDB Go Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo)
[BSON specification for GO](https://gopkg.in/mgo.v2/bson)


