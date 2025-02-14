package usecase

import "github/revaldimijaya/lacak-api/app/repository"

func InitUsecase(
	Repository repository.Repository,
) Usecase {
	return Usecase{
		Repository: Repository,
	}
}
