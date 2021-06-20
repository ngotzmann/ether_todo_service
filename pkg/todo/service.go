package todo

import (
	"errors"
	"github.com/kataras/i18n"
	"strings"
)

type IService interface {
	IsListDuplicated(name string) (bool, error)
	OverwriteExistsList(l *List) (*List, error)
	FindListByName(name string) (*List, error)
	SaveList(l *List) (*List, error)
	DeleteListByName(name string) error
	CleanOutatedLists()
	Migration() error
}

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) IService {
	return &Service{
		repo: repo,
	}
}

func (s *Service) IsListDuplicated(name string) (bool, error) {
	fl, err := s.repo.FindListByName(name)
	if err != nil {
		return false, err
	}
	if fl == nil {
		return false, nil
	} else {
		return true, nil
	}
}

//If a list with the given name is already exists it will be deleted with all appended tasks and a new list will be created
//If a list with the given name is not exist it will just created
func (s *Service) OverwriteExistsList(l *List) (*List, error) {
	l.Name = strings.ToLower(l.Name)

	err := s.repo.DeleteListByName(l)
	if err != nil {
		return nil, err
	}

	l, err = s.repo.SaveList(l)
	if err != nil {
		return nil, err
	}

	return l, nil
}

func (s *Service) FindListByName(name string) (*List, error) {
	if name == "" {
		err := errors.New(i18n.Tr("en-US", "ValidationError") + " name is missing")
		return nil, err
	}
	l, err := s.repo.FindListByName(name)
	return l, err
}

func (s *Service) SaveList(l *List) (*List, error) {
	err := l.Validation()
	if err != nil {
		return nil, err
	}
	l, err = s.OverwriteExistsList(l)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (s *Service) DeleteListByName(name string) error {
	if name == "" {
		return errors.New(i18n.Tr("en-US","ValidationError") + " name is missing")
	}
	return s.repo.DeleteListByName(&List{Name: name})
}

func (s *Service) Migration() error {
	err := s.repo.Migration()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CleanOutatedLists() {
	s.repo.DeleteOutdatedLists()
}