package main

import (
	"database/sql"
	"os"
)

type Model interface{
    
}

type Thing struct {
	ID       int
	Name     string
	Hostname string
}

var things []Thing

func init() {
	h, _ := os.Hostname()
	things = append(things, Thing{1, "Thing1", h})
	things = append(things, Thing{2, "Thing2", h})
	things = append(things, Thing{3, "Thing3", h})
}

type ThingAPI struct {
	db *sql.DB
}

func NewThingAPI(db *sql.DB) *ThingAPI {
	return &ThingAPI{db}
}

func (a *ThingAPI) GetAll() interface{} {
	return &things
}

func (a *ThingAPI) Get(id interface{}) interface{} {
	return things[0]
}

func (a *ThingAPI) Put(id interface{}, t interface{}) {

}

func (a *ThingAPI) Post(t interface{}) {

}

func (a *ThingAPI) Delete(id interface{}) {

}
