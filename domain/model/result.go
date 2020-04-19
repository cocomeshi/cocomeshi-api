package model

type Result struct {
	Success bool
	Error   *Error
	Data    interface{}
}

type Error struct {
	Code    string
	Message string
}
