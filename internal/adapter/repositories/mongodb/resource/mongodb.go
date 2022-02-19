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

const mongoServerURL = "mongodb://localhost:27017/gasplus"

//Create a Mongo Client
func NewMongoClient() (*mgo.Database, error) {
	dbname := "gasplus"
	DBSession, err := mgo.Dial(mongoServerURL)
	if err != nil {
		panic(errors.Wrap(err, "Unable to connect to Mongo database"))
	}
	log.Println("connected to database successfully")

	Db := DBSession.DB(dbname)
	return Db, nil
}

//NewMongoRepository ...
func (r *MongoRepository) Init() {
	mongoClient, err := NewMongoClient()
	r.Client = mongoClient
	r.DbName = "gasplus"
	if err != nil {
		log.Println(err.Error())
	}
}
