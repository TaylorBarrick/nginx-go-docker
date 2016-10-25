package main

//RestAPI is a generic Restful API interface with some of the basic HTTP methods
type RestAPI interface {
	GetAll() interface{}
	Get(id int) interface{}
	Put(id int, t interface{})
	Post(t interface{})
	Delete(id int)
}
