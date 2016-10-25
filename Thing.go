package main

import (
	"database/sql"
	"os"
)

//Thing is a thing :)
type Thing struct {
	ID       int
	Name     string
	Hostname string
}

var things []Thing
var thingMap map[int]Thing

func init() {
	h, _ := os.Hostname()
	things = append(things, Thing{1, "Thing1", h})
	things = append(things, Thing{2, "Thing2", h})
	things = append(things, Thing{3, "Thing3", h})
	thingMap = make(map[int]Thing)
	for i, e := range things {
		thingMap[i] = e
	}
}

//ThingAPI is an implementation of the RestAPI for thhe Thing model
type ThingAPI struct {
	db *sql.DB
}

//NewThingAPI provides a ThingAPI with an unused db today
func NewThingAPI(db *sql.DB) *ThingAPI {
	return &ThingAPI{db}
}

//GetAll returns all of the Things
func (a *ThingAPI) GetAll() interface{} {
	return &things
}

//Get returns a single Thing
func (a *ThingAPI) Get(id int) interface{} {
	return thingMap[id]
}

//Put puts updates a single item for the given ID
func (a *ThingAPI) Put(id int, t interface{}) {

}

//Post inserts a new Thing
func (a *ThingAPI) Post(t interface{}) {

}

//Delete removes an item for the given ID
func (a *ThingAPI) Delete(id int) {

}
