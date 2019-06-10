package db

type IRepository interface {
	Init(host string) error

	Close()
}
