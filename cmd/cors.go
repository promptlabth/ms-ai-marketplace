package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/blendle/zapdriver"
	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-ai-marketplace/logger"
	"go.uber.org/zap"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Request-Id")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func LoggingWithDumbBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request

		httpPayload := &zapdriver.HTTPPayload{
			RequestMethod: req.Method,
			UserAgent:     req.UserAgent(),
			RemoteIP:      req.RemoteAddr,
			Referer:       req.Referer(),
			Protocol:      req.Proto,
		}

		if req.URL != nil {
			httpPayload.RequestURL = req.URL.String()
		}

		// reqBody

		var reqBody []byte
		var mapReqBody = make(map[string]interface{})

		contentType := req.Header.Get("Content-Type")
		if strings.Contains(contentType, "application/json") {
			if c.Request.Body != nil { // Read

				reqBody, _ = io.ReadAll(c.Request.Body)

				if len(reqBody) != 0 {
					err := json.Unmarshal(reqBody, &mapReqBody)
					if err != nil {
						logger.Error(req.Context(), err.Error())
					}
				}
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody)) // Reset
		}

		// resWriter
		respWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = respWriter

		logMsg := fmt.Sprintf("Received API request: method=%s, path=%s", req.Method, c.Request.URL.Path)
		logger.Info(req.Context(), logMsg, zap.Any("request_body", mapReqBody), zap.Any("request_header", req.Header), zapdriver.HTTP(httpPayload))

		start := time.Now()

		c.Next()

		stop := time.Now()
		httpPayload.Status = respWriter.Status()
		httpPayload.ResponseSize = strconv.FormatInt(int64(respWriter.Size()), 10)

		l := stop.Sub(start)
		httpPayload.Latency = l.String()

		var mapResBody = make(map[string]interface{})
		contentType = respWriter.Header().Get("Content-Type")
		if strings.Contains(contentType, "application/json") {
			if respWriter.Size() != 0 {
				err := json.Unmarshal(respWriter.body.Bytes(), &mapResBody)
				if err != nil {
					logger.Error(req.Context(), err.Error())
				}
			}
		}

		logMsg = fmt.Sprintf("API response: method=%s, path=%s", req.Method, c.Request.URL.Path)
		logger.Info(req.Context(), logMsg, zap.Any("request_body", mapReqBody), zap.Any("response_header", respWriter.Header()), zapdriver.HTTP(httpPayload))
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
