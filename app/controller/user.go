package controller

import (
	"net/http"
	domain "prototype/domain/user"
	model "prototype/domain/user/models"
	"prototype/lib/log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase domain.IUserUsecase
	log         log.ILogs
}

func NewUserController(userUsecase domain.IUserUsecase, log log.ILogs) *UserController {
	return &UserController{
		userUsecase,
		log,
	}
}

func (handler *UserController) Fetch(c *gin.Context) {
	var (
		statusCode int
		res        Response

		ctx     = c.Request.Context()
		traceID = c.GetHeader("Trace-ID")
	)

	// setup
	res.TraceID = traceID

	defer func() {
		c.JSON(statusCode, res)
	}()

	user, err := handler.userUsecase.Fetch(ctx)

	if err != nil {
		statusCode = http.StatusInternalServerError
		res.Set(statusCode, nil, err)
		handler.log.Error(ctx, "usecase.userRepo.Fetch Error Error", err)
		return
	}

	statusCode = http.StatusOK
	res.Set(statusCode, user, nil)
}

func (handler *UserController) GetByID(c *gin.Context) {
	var (
		statusCode int
		res        Response

		ctx = c.Request.Context()
	)

	defer func() {
		c.JSON(statusCode, res)
	}()

	userIdP, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {

		statusCode = http.StatusInternalServerError
		res.Set(statusCode, nil, err)
		handler.log.Error(ctx, "strconv.Atoi(c.Param('user_id')) Error", err)
		return
	}

	user_id := uint(userIdP)

	user, err := handler.userUsecase.GetByID(ctx, user_id)

	if err != nil {
		statusCode = http.StatusInternalServerError
		res.Set(statusCode, nil, err)
		handler.log.Error(ctx, "handler.userUsecase.GetByID Error", err)
		return
	}

	statusCode = http.StatusOK
	res.Set(http.StatusInternalServerError, user, nil)
}

func (handler *UserController) Create(c *gin.Context) {
	var (
		statusCode int
		request    model.User
		res        Response

		ctx = c.Request.Context()
	)

	defer func() {
		c.JSON(statusCode, res)
	}()

	if err := c.ShouldBindJSON(&request); err != nil {

		statusCode = http.StatusBadRequest
		res.Set(statusCode, nil, err)
		handler.log.Error(ctx, "c.ShouldBindJSON Error", err)
		return
	}

	user, err := handler.userUsecase.Create(ctx, request)

	if err != nil {

		statusCode = http.StatusInternalServerError
		res.Set(http.StatusInternalServerError, nil, err)
		handler.log.Error(ctx, "handler.userUsecase.Create Error", err)
		return
	}

	statusCode = http.StatusOK
	res.Set(http.StatusOK, user, nil)
}

func (handler *UserController) Update(c *gin.Context) {
	var (
		statusCode int
		request    model.User
		res        Response

		ctx = c.Request.Context()
	)

	defer func() {
		c.JSON(statusCode, res)
	}()

	userIdP, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {

		statusCode = http.StatusBadRequest
		res.Set(statusCode, nil, err)
		handler.log.Error(ctx, "strconv.Atoi(c.Param('user_id')) Error", err)

		return
	}

	user_id := uint(userIdP)

	if err := c.ShouldBindJSON(&request); err != nil {

		statusCode = http.StatusBadRequest
		res.Set(statusCode, nil, err)
		handler.log.Error(ctx, "c.ShouldBindJSON Error", err)

		return
	}

	request.ID = user_id

	user, err := handler.userUsecase.Update(ctx, request)

	if err != nil {

		statusCode = http.StatusInternalServerError
		res.Set(http.StatusInternalServerError, nil, err)
		handler.log.Error(ctx, "handler.userUsecase.Update Error", err)

		return
	}

	statusCode = http.StatusOK
	res.Set(http.StatusOK, user, nil)
}

func (handler *UserController) Delete(c *gin.Context) {
	var (
		statusCode int
		res        Response

		ctx = c.Request.Context()
	)

	defer func() {
		c.JSON(statusCode, res)
	}()

	userIdP, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {

		statusCode = http.StatusInternalServerError
		res.Set(statusCode, nil, err)
		handler.log.Error(ctx, "strconv.Atoi(c.Param('user_id')) Error", err)

		return
	}

	user_id := uint(userIdP)

	err = handler.userUsecase.Delete(ctx, user_id)

	if err != nil {

		statusCode = http.StatusInternalServerError
		res.Set(statusCode, nil, err)
		handler.log.Error(ctx, "handler.userUsecase.Delete Error", err)

		return
	}

	statusCode = http.StatusOK
	res.Set(http.StatusOK, nil, nil)
}
