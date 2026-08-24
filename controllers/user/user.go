package controllers

import (
	"net/http"
	errWrap "user-service/common/error"
	"user-service/common/response"
	"user-service/domain/dto"
	services "user-service/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	service services.IServiceRegistry
}

type IUserController interface {
	Login(*gin.Context)
	Register(*gin.Context)
	Update(*gin.Context)
	GetUserLogin(*gin.Context)
	GetUserByUUID(*gin.Context)
}

func NewUserController(service services.IServiceRegistry) IUserController {
	return &UserController{
		service: service,
	}
}

func (c *UserController) Login(ctx *gin.Context) {
	request := &dto.LoginRequest{}

	err := ctx.ShouldBindJSON(request)
	if err != nil {
		response.HTTPResp(response.ParamsHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errWrap.ErrValidationResponse(err)

		response.HTTPResp(response.ParamsHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Err:     err,
			Data:    errResponse,
			Gin:     ctx,
		})
		return
	}

	user, err := c.service.GetUser().Login(ctx.Request.Context(), request)
	if err != nil {
		response.HTTPResp(response.ParamsHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HTTPResp(response.ParamsHTTPResp{
		Code:  http.StatusOK,
		Data:  user.User,
		Token: &user.Token,
		Gin:   ctx,
	})
	return
}
func (c *UserController) Register(ctx *gin.Context) {
	request := &dto.RegisterRequest{}

	err := ctx.ShouldBindJSON(request)
	if err != nil {
		response.HTTPResp(response.ParamsHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errWrap.ErrValidationResponse(err)

		response.HTTPResp(response.ParamsHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Err:     err,
			Data:    errResponse,
			Gin:     ctx,
		})
		return
	}

	user, err := c.service.GetUser().Register(ctx.Request.Context(), request)
	if err != nil {
		response.HTTPResp(response.ParamsHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HTTPResp(response.ParamsHTTPResp{
		Code: http.StatusOK,
		Data: user.User,
		Gin:  ctx,
	})
	return
}

func (c *UserController) Update(ctx *gin.Context) {
	request := &dto.UpdateRequest{}
	uuid := ctx.Param("uuid")

	err := ctx.ShouldBindJSON(request)
	if err != nil {
		response.HTTPResp(response.ParamsHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errWrap.ErrValidationResponse(err)

		response.HTTPResp(response.ParamsHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Err:     err,
			Data:    errResponse,
			Gin:     ctx,
		})
		return
	}

	user, err := c.service.GetUser().Update(ctx, request, uuid)
	if err != nil {
		response.HTTPResp(response.ParamsHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HTTPResp(response.ParamsHTTPResp{
		Code: http.StatusOK,
		Data: user,
		Gin:  ctx,
	})
	return
}

func (c *UserController) GetUserLogin(ctx *gin.Context) {
	user, err := c.service.GetUser().GetUserLogin(ctx.Request.Context())
	if err != nil {
		response.HTTPResp(response.ParamsHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HTTPResp(response.ParamsHTTPResp{
		Code: http.StatusOK,
		Data: user,
		Gin:  ctx,
	})
	return
}
func (c *UserController) GetUserByUUID(ctx *gin.Context) {
	user, err := c.service.GetUser().GetUserByUUID(ctx.Request.Context(), ctx.Param("uuid"))
	if err != nil {
		response.HTTPResp(response.ParamsHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HTTPResp(response.ParamsHTTPResp{
		Code: http.StatusOK,
		Data: user,
		Gin:  ctx,
	})
	return
}
