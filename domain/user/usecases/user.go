package usecases

import (
	"context"
	domain "prototype/domain/user"
	"prototype/domain/user/models"
	"prototype/lib/log"
)

type userUsecase struct {
	userRepo domain.IUserMysqlRepository
	log      log.ILogs
}

func NewUserUsecase(userRepo domain.IUserMysqlRepository, log log.ILogs) domain.IUserMysqlRepository {
	return &userUsecase{userRepo, log}
}

func (usecase userUsecase) Fetch(ctx context.Context) (result []models.User, err error) {
	result, err = usecase.userRepo.Fetch(ctx)

	if err != nil {
		usecase.log.Error(ctx, "usecase.userRepo.Fetch Error", err)
		return
	}

	return
}

func (usecase userUsecase) Create(ctx context.Context, user models.User) (result models.User, err error) {
	result, err = usecase.userRepo.Create(ctx, user)
	if err != nil {
		return
	}

	return
}

func (usecase userUsecase) Update(ctx context.Context, user models.User) (result models.User, err error) {
	userData, err := usecase.userRepo.GetByID(ctx, user.ID)
	if err != nil {
		usecase.log.Error(ctx, "usecase.userRepo.GetByID Error", err)
		return
	}

	userData.FirstName = user.FirstName
	userData.LastName = user.LastName
	userData.Email = user.Email

	result, err = usecase.userRepo.Update(ctx, userData)
	if err != nil {
		usecase.log.Error(ctx, "usecase.userRepo.Update Error", err)
		return
	}

	return
}

func (usecase userUsecase) GetByID(ctx context.Context, id uint) (result models.User, err error) {
	result, err = usecase.userRepo.GetByID(ctx, id)
	if err != nil {
		usecase.log.Error(ctx, "usecase.userRepo.GetByID Error", err)
		return
	}

	return
}

func (usecase userUsecase) Delete(ctx context.Context, id uint) (err error) {
	if err = usecase.userRepo.Delete(ctx, id); err != nil {
		usecase.log.Error(ctx, "usecase.userRepo.Delete Error", err)
		return
	}

	return
}
