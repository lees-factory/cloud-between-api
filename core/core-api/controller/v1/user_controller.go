package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"io.lees.cloud-between/core/core-api/controller/v1/request"
	"io.lees.cloud-between/core/core-api/controller/v1/response"
	"io.lees.cloud-between/core/core-domain/user"
)

type UserController struct {
	userService *user.UserService
}

func NewUserController(userService *user.UserService) *UserController {
	return &UserController{userService: userService}
}

func (ctrl *UserController) Signup(c *gin.Context) {
	var req request.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.userService.Signup(c.Request.Context(), req.Email, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "signup successful"})
}

func (ctrl *UserController) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := ctrl.userService.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toUserResponse(u))
}

func (ctrl *UserController) SocialLogin(c *gin.Context) {
	var req request.SocialLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	provider := user.SocialProvider(req.Provider)

	u, isNew, err := ctrl.userService.LoginBySocial(c.Request.Context(), req.SocialID, provider, req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if isNew {
		c.JSON(http.StatusCreated, toUserResponse(u))
		return
	}
	c.JSON(http.StatusOK, toUserResponse(u))
}

func toUserResponse(u *user.User) response.UserResponse {
	return response.UserResponse{
		ID:          u.ID,
		Email:       u.Email,
		CreatedAt:   u.CreatedAt,
		LastLoginAt: u.LastLoginAt,
	}
}
