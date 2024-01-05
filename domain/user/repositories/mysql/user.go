package repository_mysql

import (
	"context"
	domain "prototype/domain/user"
	"prototype/domain/user/models"
	"prototype/lib/log"

	"gorm.io/gorm"
)

type userMysqlRepository struct {
	DB  *gorm.DB
	log log.ILogs
}

func NewMysqlUserRepo(DB *gorm.DB, log log.ILogs) domain.IUserMysqlRepository {
	return userMysqlRepository{DB, log}
}

func (repo userMysqlRepository) Fetch(ctx context.Context) (result []models.User, err error) {
	if err = repo.DB.WithContext(ctx).Find(&result).Error; err != nil {
		repo.log.Error(ctx, "repo.DB.WithContext(ctx).Find(&result)", err)
		return
	}

	return
}

func (repo userMysqlRepository) Create(ctx context.Context, user models.User) (result models.User, err error) {
	if err = repo.DB.WithContext(ctx).Create(&user).Error; err != nil {
		repo.log.Error(ctx, "repo.DB.WithContext(ctx).Create(&user)", err)
		return
	}

	result = user
	return
}

func (repo userMysqlRepository) Update(ctx context.Context, user models.User) (result models.User, err error) {
	if err = repo.DB.WithContext(ctx).Save(&user).Error; err != nil {
		repo.log.Error(ctx, "repo.DB.WithContext(ctx).Save(&user)", err)
		return
	}

	result = user
	return
}

func (repo userMysqlRepository) GetByID(ctx context.Context, id uint) (result models.User, err error) {
	if err = repo.DB.WithContext(ctx).Where("id = ?", id).First(&result).Error; err != nil {
		repo.log.Error(ctx, "repo.DB.WithContext(ctx).Where('id = ?', id).First(&result)", err)
		return
	}

	return
}

func (repo userMysqlRepository) Delete(ctx context.Context, id uint) (err error) {
	if err = repo.DB.WithContext(ctx).Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		repo.log.Error(ctx, "repo.DB.WithContext(ctx).Where('id = ?', id).Delete(&models.User{})", err)
		return
	}

	return
}
