package main

import (
	"database/sql"
	"os"
)

type Stuff struct {
	ID       int
	Name     string
	Hostname string
}

var stuffs []Stuff

func init() {
	h, _ := os.Hostname()
	stuffs = append(stuffs, Stuff{1, "Stuff1", h})
	stuffs = append(stuffs, Stuff{2, "Stuff2", h})
	stuffs = append(stuffs, Stuff{3, "Stuff3", h})
}

type StuffAPI struct {
	db *sql.DB
}

func NewStuffAPI(db *sql.DB) *StuffAPI {
	return &StuffAPI{db}
}

func (a *StuffAPI) GetAll() interface{} {
	return &stuffs
}

func (a *StuffAPI) Get(id interface{}) interface{} {
	return stuffs[0]
}

func (a *StuffAPI) Put(id interface{}, t interface{}) {

}

func (a *StuffAPI) Post(t interface{}) {

}

func (a *StuffAPI) Delete(id interface{}) {

}
