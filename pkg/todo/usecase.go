package todo

type Usecase interface {
	Function() error
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

func (uc *usecase) Function() error {
	return uc.service.Function()
}
