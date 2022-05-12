package db

type MyDb interface {
	User() UserDB
	Order() OrderDB
}
