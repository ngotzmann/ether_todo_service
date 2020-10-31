package todo

type IRepository interface {
	FindListByName(name string) (*List, error)
	SaveList(l *List) (*List, error)
	DeleteListByName(l *List) error
	DeleteOutdatedLists()
	Migration() error
}
