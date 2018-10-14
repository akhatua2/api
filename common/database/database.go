package database

import (
	"gopkg.in/mgo.v2"
)

/*
	Database interface exposing the methods necessary to querying, inserting, updating, upserting, and removing records
*/
type Database interface {
	Connect(host string) error
	FindOne(collection_name string, query interface{}, result interface{}) error
	FindAll(collection_name string, query interface{}, result interface{}) error
	RemoveOne(collection_name string, query interface{}) error
	RemoveAll(collection_name string, query interface{}) (*mgo.ChangeInfo, error)
	Insert(collection_name string, item interface{}) error
	Upsert(collection_name string, selector interface{}, update interface{}) (*mgo.ChangeInfo, error)
	Update(collection_name string, selector interface{}, update interface{}) error
	UpdateAll(collection_name string, selector interface{}, update interface{}) (*mgo.ChangeInfo, error)
}

/*
	An alias of a string -> interface{} map used for database queries and selectors
*/
type QuerySelector map[string]interface{}
