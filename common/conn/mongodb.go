package conn

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/conf"
)

var _mongoClient *mongo.Client
var _mongoOnce sync.Once

func MongoConn() *mongo.Client {
	if conf.Arg.Mongo == nil {
		panic("没有设置 mongodb的配置信息")
	}
	_mongoOnce.Do(func() {
		var err error
		url := "mongodb://"+conf.Arg.Mongo.User+":"+conf.Arg.Mongo.Password+"@"+conf.Arg.Mongo.Host
		logger.Info("mongodb url = ", url)
		clientOptions := options.Client().ApplyURI(url).SetConnectTimeout(20*time.Second)
		_mongoClient, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			panic(err)
		}
		err = _mongoClient.Ping(context.TODO(), nil)
		if err != nil {
			panic(err)
		}
	})
	return _mongoClient
}

func GetMongoDB(dbName string) *mongo.Database {
	return MongoConn().Database(dbName)
}

func GetMongoCollection(dbName, collection string) *mongo.Collection {
	return GetMongoDB(dbName).Collection(collection)
}

type MongoCollection struct {
	DataBase string
	Collection string
	Conn *mongo.Collection
}

func NewMongoCollection(dbName, collection string) *MongoCollection {
	return &MongoCollection{
		DataBase: dbName,
		Collection: collection,
		Conn: GetMongoDB(dbName).Collection(collection),
	}
}

func (mo *MongoCollection) Get() *mongo.Collection {
	return mo.Conn
}

func (mo *MongoCollection) InsertOne(value interface{}) (err error) {
	_, err = mo.Conn.InsertOne(context.TODO(), value)
	return
}

func (mo *MongoCollection) InsertMany(value []interface{}) (err error) {
	_, err = mo.Conn.InsertMany(context.TODO(), value)
	return
}

func (mo *MongoCollection) FindOne(filter, result interface{}) (err error) {
	err = mo.Conn.FindOne(context.TODO(), filter).Decode(&result)
	return
}

func (mo *MongoCollection) UpdateOne(filter, update interface{}) (err error) {
	_, err = mo.Conn.UpdateOne(context.TODO(), filter, update)
	return
}

func (mo *MongoCollection) DeleteMany(filter interface{}) (err error) {
	_, err = mo.Conn.DeleteMany(context.TODO(), filter)
	return
}

