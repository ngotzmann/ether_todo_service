package todo

import (
	"github.com/ngotzmann/gorror"
)

type IUsecase interface {
	FindListByName(name string) (*List, error)
	SaveList(l *List) (*List, error)
	DeleteListByName(name string) error
	CleanOutatedLists()
}

type Usecase struct {
	repo    IRepository
	service IService
}

func NewUsecase(repo IRepository, service IService) IUsecase {
	return &Usecase{
		repo:    repo,
		service: service,
	}
}

func (uc *Usecase) FindListByName(name string) (*List, error) {
	if name == "" {
		err := gorror.CreateError(gorror.ValidationError, "name is not set")
		return nil, err
	}
	l, err := uc.repo.FindListByName(name)
	return l, err
}

func (uc *Usecase) SaveList(l *List) (*List, error) {
	err := l.Validation()
	if err != nil {
		return nil, err
	}
	l, err = uc.service.OverwriteExistsList(l)
	if err != nil {
			return nil, err
		}
	return l, nil
}

func (uc *Usecase) DeleteListByName(name string) error {
	if name == "" {
		return gorror.CreateError(gorror.ValidationError, "name is missing")
	}
	return uc.repo.DeleteListByName(&List{Name: name})
}

func (uc *Usecase) CleanOutatedLists() {
	uc.repo.DeleteOutdatedLists()
}