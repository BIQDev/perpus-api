package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type dbMongo struct{
	Ctx context.Context
	Connection *mongo.Client
	Database *mongo.Database
	Error error
}

type IMongo interface {
	Init()
	Conn() *mongo.Client
	DB() *mongo.Database
	Err() error
	GetCtx() context.Context
}

func NewMongo() IMongo{
	return &dbMongo{}
}

func (_d *dbMongo) GetCtx() context.Context {
	return _d.Ctx
}

func (_d *dbMongo) Init() {
	_d.Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	_d.Connection, _d.Error = mongo.Connect(_d.Ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if _d.Error == nil {
		_d.Database = _d.Connection.Database(os.Getenv("MONGO_DBNAME"))
	}
	log.Println("Connected to MongoDB")
}

func (_d *dbMongo) Conn() *mongo.Client {
	return _d.Connection
}

func (_d *dbMongo) Err() error {
	return _d.Error
}

func (_d *dbMongo) DB() *mongo.Database {
	return _d.Database
}

var Mongo = NewMongo()

