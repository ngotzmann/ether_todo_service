package todo

import "strings"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
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
