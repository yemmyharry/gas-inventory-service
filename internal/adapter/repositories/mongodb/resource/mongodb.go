package resource

import (
	"github.com/globalsign/mgo"
	"github.com/pkg/errors"
	"log"
)

type MongoRepository struct {
	//*mongo.Client
	Client *mgo.Database
	DbName string
}

//Create a Mongo Client
func NewMongoClient(dbName, dbHost string) (*mgo.Database, error) {
	dbname := dbName
	DBSession, err := mgo.Dial(dbHost)
	if err != nil {
		panic(errors.Wrap(err, "Unable to connect to Mongo database"))
	}
	log.Println("connected to database successfully")

	Db := DBSession.DB(dbname)
	return Db, nil
}

//NewMongoRepository ...
func (r *MongoRepository) Init(dbHost, dbName string) {
	mongoClient, err := NewMongoClient(dbName, dbHost)
	r.Client = mongoClient
	r.DbName = dbName
	if err != nil {
		log.Println(err.Error())
	}
}
