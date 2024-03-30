package app

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-orch-user-service/logger"
)

type context struct {
	ctx        *gin.Context
	logHandler slog.Handler
}

func NewContext(c *gin.Context, logHandler slog.Handler) Context {
	return &context{ctx: c, logHandler: logHandler}
}

func (c *context) Bind(v any) error {
	return c.ctx.ShouldBindJSON(v)
}

func (c *context) WriteHeader(key, value string) {
	c.ctx.Writer.Header().Set(key, value)
}

func (c *context) Request() *http.Request {
	return c.ctx.Request
}

func (c *context) AbortWithStatus(code int) {
	c.ctx.AbortWithStatus(code)
}

func (c *context) Next() {
	c.ctx.Next()
}

func (c *context) Param(key string) string {
	return c.ctx.Param(key)
}

func (c *context) OK(v any) {
	if v == nil {
		c.ctx.Status(http.StatusOK)
		return
	}
	c.ctx.JSON(http.StatusOK, Response{
		Status: Success,
		Data:   v,
	})
}

func (c *context) BadRequest(err error) {
	logger.AppErrorf(c.logHandler, "%s", err)
	c.ctx.JSON(http.StatusBadRequest, Response{
		Status:  Fail,
		Message: err.Error(),
	})
}

func (c *context) StoreError(err error) {
	logger.AppErrorf(c.logHandler, "%s", err)
	c.ctx.JSON(storeErrorStatus, Response{
		Status:  Fail,
		Message: err.Error(),
	})
}

func NewGinHandler(handler HandlerFunc, logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewContext(c, logger.Handler().WithAttrs([]slog.Attr{slog.String("transaction-id", c.Request.Header.Get("transaction-id"))})))
	}
}

type RouterGin struct {
	*gin.Engine
	logger *slog.Logger
}

func NewRouterGin(logger *slog.Logger) *RouterGin {
	r := gin.Default()

	return &RouterGin{Engine: r, logger: logger}
}

func (r *RouterGin) Use(h ...gin.HandlerFunc) {
	r.Engine.Use(h...)
}

func (r *RouterGin) GET(path string, handler HandlerFunc) {
	r.Engine.GET(path, NewGinHandler(handler, r.logger))
}

func (r *RouterGin) POST(path string, handler HandlerFunc) {
	r.Engine.POST(path, NewGinHandler(handler, r.logger))
}
