package usecases

import (
	"context"
	"errors"
	domain "prototype/domain/user"
	"prototype/domain/user/mocks"
	"prototype/domain/user/models"
	"reflect"
	"testing"
)

func Test_userUsecase_Fetch(t *testing.T) {
	ctx := context.Background()

	user := []models.User{
		{
			ID:        1,
			Email:     "test@gmail.com",
			Username:  "test",
			FirstName: "test",
			LastName:  "test",
		},
	}

	userRepoSuccess := new(mocks.UserRepository)
	userRepoSuccess.On("Fetch", ctx).Return(user, nil)

	userRepoError := new(mocks.UserRepository)
	userRepoError.On("Fetch", ctx).Return([]models.User{}, errors.New("data tidak ditemukan"))

	type fields struct {
		userRepo domain.IUserMysqlRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult []models.User
		wantErr    bool
	}{
		{
			name:   "success",
			fields: fields{userRepo: userRepoSuccess},
			args: args{
				ctx: ctx,
			},
			wantResult: user,
			wantErr:    false,
		},
		{
			name:   "success",
			fields: fields{userRepo: userRepoError},
			args: args{
				ctx: ctx,
			},
			wantResult: []models.User{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := userUsecase{
				userRepo: tt.fields.userRepo,
			}
			gotResult, err := usecase.Fetch(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("userUsecase.Fetch() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_userUsecase_Create(t *testing.T) {
	ctx := context.Background()

	user := models.User{
		ID:        1,
		Email:     "test@gmail.com",
		Username:  "test",
		FirstName: "test",
		LastName:  "test",
	}

	userRepoSuccess := new(mocks.UserRepository)
	userRepoSuccess.On("Create", ctx, user).Return(user, nil)

	userRepoError := new(mocks.UserRepository)
	userRepoError.On("Create", ctx, user).Return(models.User{}, errors.New("data tidak ditemukan"))

	type fields struct {
		userRepo domain.IUserMysqlRepository
	}
	type args struct {
		ctx  context.Context
		user models.User
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult models.User
		wantErr    bool
	}{
		{
			name: "success",
			fields: fields{
				userRepo: userRepoSuccess,
			},
			args: args{
				ctx:  ctx,
				user: user,
			},
			wantResult: user,
			wantErr:    false,
		},
		{
			name: "failed",
			fields: fields{
				userRepo: userRepoError,
			},
			args: args{
				ctx:  ctx,
				user: user,
			},
			wantResult: models.User{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := userUsecase{
				userRepo: tt.fields.userRepo,
			}
			gotResult, err := usecase.Create(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("userUsecase.Create() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_userUsecase_Update(t *testing.T) {
	ctx := context.Background()

	user := models.User{
		ID:        1,
		Email:     "test@gmail.com",
		Username:  "test",
		FirstName: "test",
		LastName:  "test",
	}

	userRepoSuccess := new(mocks.UserRepository)
	userRepoSuccess.On("GetByID", ctx, uint(1)).Return(user, nil)
	userRepoSuccess.On("Update", ctx, user).Return(user, nil)

	userRepoError := new(mocks.UserRepository)
	userRepoError.On("GetByID", ctx, uint(1)).Return(models.User{}, errors.New("data tidak ditemukan"))
	userRepoError.On("Update", ctx, user).Return(models.User{}, errors.New("data tidak ditemukan"))

	type fields struct {
		userRepo domain.IUserMysqlRepository
	}
	type args struct {
		ctx  context.Context
		user models.User
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult models.User
		wantErr    bool
	}{
		{
			name: "success",
			fields: fields{
				userRepo: userRepoSuccess,
			},
			args: args{
				ctx:  ctx,
				user: user,
			},
			wantResult: user,
			wantErr:    false,
		},
		{
			name: "failed",
			fields: fields{
				userRepo: userRepoError,
			},
			args: args{
				ctx:  ctx,
				user: user,
			},
			wantResult: models.User{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := userUsecase{
				userRepo: tt.fields.userRepo,
			}
			gotResult, err := usecase.Update(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("userUsecase.Update() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_userUsecase_GetByID(t *testing.T) {
	ctx := context.Background()

	user := models.User{
		ID:        1,
		Email:     "test@gmail.com",
		Username:  "test",
		FirstName: "test",
		LastName:  "test",
	}

	userRepoSuccess := new(mocks.UserRepository)
	userRepoSuccess.On("GetByID", ctx, uint(1)).Return(user, nil)

	userRepoError := new(mocks.UserRepository)
	userRepoError.On("GetByID", ctx, uint(1)).Return(models.User{}, errors.New("data tidak ditemukan"))

	type fields struct {
		userRepo domain.IUserMysqlRepository
	}
	type args struct {
		ctx context.Context
		id  uint
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult models.User
		wantErr    bool
	}{
		{
			name: "success",
			fields: fields{
				userRepo: userRepoSuccess,
			},
			args: args{
				ctx: ctx,
				id:  user.ID,
			},
			wantResult: user,
			wantErr:    false,
		},
		{
			name: "failed",
			fields: fields{
				userRepo: userRepoError,
			},
			args: args{
				ctx: ctx,
				id:  user.ID,
			},
			wantResult: models.User{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := userUsecase{
				userRepo: tt.fields.userRepo,
			}
			gotResult, err := usecase.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("userUsecase.GetByID() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_userUsecase_Delete(t *testing.T) {
	ctx := context.Background()

	userRepoSuccess := new(mocks.UserRepository)
	userRepoSuccess.On("Delete", ctx, uint(1)).Return(nil)

	userRepoError := new(mocks.UserRepository)
	userRepoError.On("Delete", ctx, uint(1)).Return(errors.New("data tidak ditemukan"))

	type fields struct {
		userRepo domain.IUserMysqlRepository
	}
	type args struct {
		ctx context.Context
		id  uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				userRepo: userRepoSuccess,
			},
			args: args{
				ctx: ctx,
				id:  uint(1),
			},
			wantErr: false,
		},
		{
			name: "failed",
			fields: fields{
				userRepo: userRepoError,
			},
			args: args{
				ctx: ctx,
				id:  uint(1),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := userUsecase{
				userRepo: tt.fields.userRepo,
			}
			if err := usecase.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
