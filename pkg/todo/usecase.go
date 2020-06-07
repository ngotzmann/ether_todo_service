package todo

import "github.com/ngotzmann/gorror"

type Usecase interface {
	FindListByName(name string) (*List, error)
	SaveList(l *List) (*List, error)
}

type usecase struct {
	repo    Repository
	service *Service
}

func NewUsecase(repo Repository, service *Service) *usecase {
	return &usecase{
		repo:    repo,
		service: service,
	}
}

func (uc *usecase) FindListByName(name string) (*List, error) {
	if name == "" {
		err := gorror.CreateError(gorror.ValidationError, "name is not set")
		return nil, err
	}
	l, err := uc.repo.FindListByName(name)
	return l, err
}

func (uc *usecase) SaveList(l *List) (*List, error) {
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
