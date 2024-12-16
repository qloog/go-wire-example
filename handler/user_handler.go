package handler

import (
	"go-wire-example/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler 包含 UserService
type UserHandler struct {
	UserService *service.UserService
}

// NewUserHandler 创建一个新的 UserHandler
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// LoginHandler 是一个具体的 HandlerFunc
func (h *UserHandler) LoginHandler(c *gin.Context) {
	// 从请求中提取参数（例如 JSON 或表单参数）
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 调用 UserService 的逻辑
	user, err := h.UserService.Login(c.Request.Context(), loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"user": user})
}
