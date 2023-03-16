package servantgo

type Task interface {
	Hash() Hash
	Exec()
}
