package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authService *AuthService
}

func NewAuthHandler(conn *gorm.DB) *AuthHandler {
	return &AuthHandler{authService: NewAuthService(conn)}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req AuthRequest

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorAuthResponse{Error: err.Error()})

		return
	}

	authToken, err := h.authService.AuthUser(c.Request.Context(), req.Email, req.Password, RegisterAction)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorAuthResponse{Error: err.Error()})

		return
	}

	c.JSON(http.StatusOK, SuccessAuthResponse{AuthToken: authToken})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req AuthRequest

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorAuthResponse{Error: err.Error()})

		return
	}

	authToken, err := h.authService.AuthUser(c.Request.Context(), req.Email, req.Password, LoginAction)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorAuthResponse{Error: err.Error()})

		return
	}

	c.JSON(http.StatusOK, SuccessAuthResponse{AuthToken: authToken})
}
