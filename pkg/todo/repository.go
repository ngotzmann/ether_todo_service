package todo

type Repository interface {
	FindListByName(name string) (*List, error)
	SaveList(l *List) (*List, error)
	DeleteListByName(l *List) error
}
