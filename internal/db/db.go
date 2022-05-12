package db

type MyDb interface {
	User() UserDB
}
