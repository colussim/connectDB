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
* The official [MongoDB Go Driver] (https://github.com/mongodb/mongo-go-driver/) 
* The Package bson is an implementation of the BSON specification for GO : [mgo.v2/bson](https://gopkg.in/mgo.v2/bson)

# Initial setup

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

* **GetMongoClient()** : *Return mongodb connection (mongo.Client),use variable CONNECTIONSTRING*
* **GetCollectionAll(Coll string) ([]bson.M, error)** : *Reading All Documents from a Collection, use return an object of type []bson.M*
* **GetCountDoc(Coll string, req bson.M) (int64, error)** : *Return number of document in Collection, takes as parameters the name of the* *collection (type string) and filter (type bson.M) ex : **req:=bson.M{"_id": id}***
* **GetReqCollectionAll(Coll string, req bson.M) ([]bson.M, error)** : *Return Documents from a Collection with a Filter,takes as parameters the name of the* *collection (type string) and filter (type bson.M)*
*  **InsertCollection(Coll string, InsertD interface{}) (mongo.InsertOneResult, error)** : *Insert Documents in a Collection, takes as parameters the name of the collection (type string) and the values to insert (type interface)*
* **RemoveCollection(Coll string, IDDist string) (mongo.DeleteResult, error)** : *Remove Documents in a Collection,takes as parameters the name of the collection (type string) and the _id of document*
* **UpdateCollection(Coll string, IDDist int, request bson.M) (mongo.UpdateResult, error)** : *Update Documents in a Collection,akes as parameters the name of the collection (type string) and the _id of document*
