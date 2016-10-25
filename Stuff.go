package main

import (
	"database/sql"
	"os"
)

//Stuff is stuff :)
type Stuff struct {
	ID       int
	Name     string
	Hostname string
}

var stuffs []Stuff

func init() {
	//using this to demo load balancing from nginx
	h, _ := os.Hostname()
	stuffs = append(stuffs, Stuff{1, "Stuff1", h})
	stuffs = append(stuffs, Stuff{2, "Stuff2", h})
	stuffs = append(stuffs, Stuff{3, "Stuff3", h})
}

//StuffAPI is an implementation of the RestAPI for thhe Stuff model
type StuffAPI struct {
	db *sql.DB
}

//NewStuffAPI provides a StuffAPI with an unused db today
func NewStuffAPI(db *sql.DB) *StuffAPI {
	return &StuffAPI{db}
}

//GetAll returns all of the Stuff
func (a *StuffAPI) GetAll() interface{} {
	return &stuffs
}

//Get returns a single Stuff
func (a *StuffAPI) Get(id int) interface{} {
	return stuffs[id]
}

//Put puts updates a single item for the given ID
func (a *StuffAPI) Put(id int, t interface{}) {

}

//Post inserts a new Stuff
func (a *StuffAPI) Post(t interface{}) {

}

//Delete removes an item for the given ID
func (a *StuffAPI) Delete(id int) {

}
