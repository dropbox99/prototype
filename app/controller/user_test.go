package controller

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	domain "prototype/domain/user"
	"prototype/domain/user/mocks"
	"prototype/domain/user/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup(userUsecase domain.IUserUsecase) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	handler := &UserController{
		userUsecase: userUsecase,
	}

	g.GET("/user", handler.Fetch)
	g.GET("/user/:user_id", handler.GetByID)
	g.POST("/user", handler.Create)
	g.PUT("/user/:user_id", handler.Update)
	g.DELETE("/user/:user_id", handler.Delete)

	return g
}

func TestUserController_Fetch(t *testing.T) {
	user := []models.User{
		{
			ID:        1,
			Email:     "test@gmail.com",
			Username:  "test",
			FirstName: "test",
			LastName:  "test",
		},
	}

	userUsecaseSuccess := new(mocks.UserUsecase)
	userUsecaseSuccess.On("Fetch", mock.Anything).Return(user, nil)

	type fields struct {
		userUsecase domain.IUserUsecase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				userUsecase: userUsecaseSuccess,
			},
			args: args{
				c: &gin.Context{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup router
			g := setup(tt.fields.userUsecase)

			switch tt.name {
			case "success":
				w := httptest.NewRecorder()

				req, _ := http.NewRequest("GET", "/user", nil)
				g.ServeHTTP(w, req)

				assert.Equal(t, 200, w.Code)
				userUsecaseSuccess.AssertExpectations(t)
			}
		})
	}
}

func TestUserController_Create(t *testing.T) {
	ctx := context.Background()

	user := models.User{
		ID:        1,
		Email:     "test@gmail.com",
		Username:  "test",
		FirstName: "test",
		LastName:  "test",
	}

	userUsecaseSuccess := new(mocks.UserUsecase)
	userUsecaseSuccess.On("Create", ctx, user).Return(user, nil)

	type fields struct {
		userUsecase domain.IUserMysqlRepository
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				userUsecase: userUsecaseSuccess,
			},
			args: args{
				c: &gin.Context{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup router
			g := setup(tt.fields.userUsecase)

			switch tt.name {
			case "success":
				w := httptest.NewRecorder()

				jsonBody := []byte(`{
					"id": 1,
					"name" : "test",
					"email" : "test@gmail.com",
					"username" : "test",
					"firstname": "test",
					"lastname": "test"
				}`)

				bodyReader := bytes.NewReader(jsonBody)

				req, _ := http.NewRequest("POST", "/user", bodyReader)
				g.ServeHTTP(w, req)

				assert.Equal(t, 200, w.Code)
				userUsecaseSuccess.AssertExpectations(t)
			}
		})
	}
}

func TestUserController_GetByID(t *testing.T) {
	user := models.User{
		ID:        1,
		Email:     "test@gmail.com",
		Username:  "test",
		FirstName: "test",
		LastName:  "test",
	}

	userUsecaseSuccess := new(mocks.UserUsecase)
	userUsecaseSuccess.On("GetByID", mock.Anything, uint(1)).Return(user, nil)

	type fields struct {
		userUsecase domain.IUserMysqlRepository
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				userUsecase: userUsecaseSuccess,
			},
			args: args{
				c: &gin.Context{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup router
			g := setup(tt.fields.userUsecase)

			switch tt.name {
			case "success":
				w := httptest.NewRecorder()

				req, _ := http.NewRequest("GET", "/user/1", nil)
				g.ServeHTTP(w, req)

				assert.Equal(t, 200, w.Code)
				userUsecaseSuccess.AssertExpectations(t)
			}
		})
	}
}

func TestUserController_Update(t *testing.T) {
	user := models.User{
		ID:        1,
		Email:     "test@gmail.com",
		Username:  "test",
		FirstName: "test",
		LastName:  "test",
	}

	userUsecaseSuccess := new(mocks.UserUsecase)
	userUsecaseSuccess.On("Update", mock.Anything, user).Return(user, nil)

	type fields struct {
		userUsecase domain.IUserMysqlRepository
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				userUsecase: userUsecaseSuccess,
			},
			args: args{
				c: &gin.Context{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup router
			g := setup(tt.fields.userUsecase)

			switch tt.name {
			case "success":
				w := httptest.NewRecorder()

				jsonBody := []byte(`{
					"id": 1,
					"name" : "test",
					"email" : "test@gmail.com",
					"username" : "test",
					"firstname": "test",
					"lastname": "test"
				}`)

				bodyReader := bytes.NewReader(jsonBody)

				req, _ := http.NewRequest("PUT", "/user/1", bodyReader)
				g.ServeHTTP(w, req)

				assert.Equal(t, 200, w.Code)
				userUsecaseSuccess.AssertExpectations(t)
			}
		})
	}
}

func TestUserController_Delete(t *testing.T) {
	userUsecaseSuccess := new(mocks.UserUsecase)
	userUsecaseSuccess.On("Delete", mock.Anything, uint(1)).Return(nil)

	type fields struct {
		userUsecase domain.IUserMysqlRepository
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				userUsecase: userUsecaseSuccess,
			},
			args: args{
				c: &gin.Context{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup router
			g := setup(tt.fields.userUsecase)

			switch tt.name {
			case "success":
				w := httptest.NewRecorder()

				req, _ := http.NewRequest("DELETE", "/user/1", nil)
				g.ServeHTTP(w, req)

				assert.Equal(t, 200, w.Code)
				userUsecaseSuccess.AssertExpectations(t)
			}
		})
	}
}
