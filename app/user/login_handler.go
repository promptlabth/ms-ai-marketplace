package user

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-ai-marketplace/app"
	"github.com/promptlabth/ms-ai-marketplace/logger"
	"go.uber.org/zap"
)

func (u *Handler) LoginHandler(c *gin.Context) {
	ctx := c.Request.Context()

	var req LoginRequestDomain

	authorizationToken := c.Request.Header.Get("authorization")
	logger.Info(ctx, authorizationToken)
	tokenSplite := strings.Fields(authorizationToken)
	if len(tokenSplite) > 1 && tokenSplite[0] == "Bearer" {
		req.Authorization = tokenSplite[1]
	} else {
		req.Authorization = tokenSplite[0]
	}

	if err := c.Bind(&req); err != nil {
		c.JSON(200, app.Response[any]{
			Code:    4004,
			Error:   err.Error(),
			Message: "Error For Binding request body",
		})
		return
	}

	res, err := u.userUsecase.LoginService(ctx, req)
	if err != nil {
		c.JSON(200, app.Response[any]{
			Code:    4000,
			Error:   err.Error(),
			Message: "Login Service Error",
		})
		return
	}

	logger.Info(ctx, "log response value", zap.Any("data", res))
	c.JSON(200, app.Response[LoginResponseDomain]{
		Code: 2000,
		Data: &LoginResponseDomain{
			User: res.User,
			Plan: res.Plan,
		},
	})
}
