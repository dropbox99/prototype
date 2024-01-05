package mocks

import (
	"context"
	"prototype/domain/user/models"

	"github.com/stretchr/testify/mock"
)

type UserUsecase struct {
	mock.Mock
}

func (m *UserUsecase) Fetch(ctx context.Context) ([]models.User, error) {
	ret := m.Called(ctx)

	var (
		r0 []models.User
		r1 error
	)

	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]models.User)
	}

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *UserUsecase) Create(ctx context.Context, user models.User) (models.User, error) {
	ret := m.Called(ctx, user)

	var (
		r0 models.User
		r1 error
	)

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(models.User)
	}

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *UserUsecase) Update(ctx context.Context, user models.User) (models.User, error) {
	ret := m.Called(ctx, user)

	var (
		r0 models.User
		r1 error
	)

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(models.User)
	}

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *UserUsecase) GetByID(ctx context.Context, id uint) (models.User, error) {
	ret := m.Called(ctx, id)

	var (
		r0 models.User
		r1 error
	)

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(models.User)
	}

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *UserUsecase) Delete(ctx context.Context, id uint) error {
	ret := m.Called(ctx, id)

	var (
		r0 error
	)

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}
