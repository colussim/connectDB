# connectDB
This is a simple go package for a MongoDB database that I needed for a project.
This package allows :
* Connect to MongoDB
* Reading All Documents from a Collection
* Return number of document in Collection
* Return Documents from a Collection with a Filter
* Insert Documents in a Collection
* Update Documents in a Collection

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
mongodb://[username:password@]host1[:port1][/[defaultauthdb][?options]]
* var DB = Name of the database
* var ISSUES = Name of the collection

### Functions

* GetMongoClient() : Return mongodb connection use variable CONNECTIONSTRING
