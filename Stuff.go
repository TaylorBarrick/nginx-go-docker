package main

import "os"

//Stuff is stuff :)
type Stuff struct {
	ID       int
	Name     string
	Hostname string
}

var stuffMap map[int]Stuff

func init() {
	//using this to demo load balancing from nginx
	h, _ := os.Hostname()
	stuffMap = make(map[int]Stuff)
	stuffMap[1] = Stuff{1, "Stuff1", h}
	stuffMap[2] = Stuff{2, "Stuff2", h}
	stuffMap[3] = Stuff{3, "Stuff3", h}
}

//StuffAPI is an implementation of the RestAPI for thhe Stuff model
type StuffAPI struct{}

//NewStuffAPI provides a StuffAPI
func NewStuffAPI() *StuffAPI {
	return &StuffAPI{}
}

//GetAll returns all of the Stuff
func (a *StuffAPI) GetAll() interface{} {
	var s []Stuff
	for _, e := range stuffMap {
		s = append(s, e)
	}
	return s
}

//Get returns a single Stuff
func (a *StuffAPI) Get(id int) interface{} {
	return stuffMap[id]
}

//Put puts updates a single item for the given ID
func (a *StuffAPI) Put(id int, t interface{}) {

}

//Post inserts a new Stuff
func (a *StuffAPI) Post(t interface{}) {

}

//Delete removes an item for the given ID
func (a *StuffAPI) Delete(id int) {
	delete(stuffMap, id)
}
