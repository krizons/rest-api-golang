package db

import "context"

type MyDb interface {
	User() UserDB
	Order() OrderDB
	Close(ctx context.Context) error
}
