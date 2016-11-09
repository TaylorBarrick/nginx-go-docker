package main

import "os"

//Thing is a thing :)
type Thing struct {
	ID       int
	Name     string
	Hostname string
}

var thingMap map[int]Thing

func init() {
	h, _ := os.Hostname()
	thingMap = make(map[int]Thing)
	thingMap[1] = Thing{1, "Thing1", h}
	thingMap[2] = Thing{2, "Thing2", h}
	thingMap[3] = Thing{3, "Thing3", h}
}

//ThingAPI is an implementation of the RestAPI for thhe Thing model
type ThingAPI struct{}

//NewThingAPI provides a ThingAPI
func NewThingAPI() *ThingAPI {
	return &ThingAPI{}
}

//GetAll returns all of the Things
func (a *ThingAPI) GetAll() interface{} {
	var t []Thing
	for _, e := range thingMap {
		t = append(t, e)
	}
	return t
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
	delete(thingMap, id)
}
