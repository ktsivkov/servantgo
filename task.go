package servantgo

type Hash string

type Task interface {
	Hash() Hash
	Exec()
}
