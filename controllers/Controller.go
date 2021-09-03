package controllers

type IController interface {
	Index() interface{}
	Show() interface{}
	Store() interface{}
	Update() interface{}
}
