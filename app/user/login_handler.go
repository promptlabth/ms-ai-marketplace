package user

import (
	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-ai-marketplace/app"
)

func (u *Handler) LoginHandler(c *gin.Context) {
	ctx := c.Request.Context()

	var req LoginRequestDomain

	if err := c.BindHeader(&req); err != nil {
		c.JSON(200, app.Response[any]{
			Code:    4003,
			Error:   err.Error(),
			Message: "Error For Binding request header",
		})
		return
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
	}

	c.JSON(200, app.Response[LoginResponseDomain]{
		Code: 2000,
		Data: &LoginResponseDomain{
			User: res.User,
			Plan: res.Plan,
		},
	})

}
