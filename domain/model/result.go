package model

type Result struct {
	success bool
	error   Error
	data    interface{}
}

type Error struct {
	Code    string
	Message string
}
