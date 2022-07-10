package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/richiMarchi/scratchpay-challenge/internal/core/domain"
	"github.com/richiMarchi/scratchpay-challenge/internal/core/ports"
)

type usersHandler struct {
	usersService ports.UsersService
}

func New(usersSvc ports.UsersService) *usersHandler {
	return &usersHandler{
		usersService: usersSvc,
	}
}

func (hdl *usersHandler) ListUsers(ctx *gin.Context) {
	users, err := hdl.usersService.List()
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": "error during users list retrieval"})
		return
	}

	ctx.JSON(200, users)
}

func (hdl *usersHandler) GetUser(ctx *gin.Context) {
	userId, convErr := strconv.Atoi(ctx.Param("userId"))
	if convErr != nil || userId < 0 {
		ctx.AbortWithStatusJSON(400, gin.H{"message": "the userId must be a not negative integer"})
		return
	}

	users, err := hdl.usersService.Get(uint(userId))
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": "error during users retrieval: " + err.Error()})
		return
	}

	ctx.JSON(200, users)
}

func (hdl *usersHandler) CreateUser(ctx *gin.Context) {
	userDto := domain.User{}
	if err := ctx.BindJSON(&userDto); err != nil {
		ctx.AbortWithStatusJSON(400, "request body is not valid: "+err.Error())
		return
	}

	err := hdl.usersService.Create(userDto.Id, userDto.Name)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": "error during user creation: " + err.Error()})
		return
	}

	ctx.Status(200)
}
