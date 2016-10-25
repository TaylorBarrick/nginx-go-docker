package main

type RestAPI interface {
	GetAll() interface{}
	Get(id interface{}) interface{}
	Put(id interface{}, t interface{})
	Post(t interface{})
	Delete(id interface{})
}
