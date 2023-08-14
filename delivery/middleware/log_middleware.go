package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/interview-bootcamp-apps/utils/logger"
	"github.com/jutionck/interview-bootcamp-apps/utils/model"
	"time"
)

type LogMiddleware interface {
	LogRequest() gin.HandlerFunc
}

type logMiddleware struct {
	loggerService logger.MyLogger
}

func (l *logMiddleware) LogRequest() gin.HandlerFunc {
	err := l.loggerService.InitializeLogger()
	if err != nil {
		return nil
	}

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Since(startTime)
		entryLog := model.RequestLog{
			EndTime:      endTime,
			StatusCode:   c.Writer.Status(),
			ClientIP:     c.ClientIP(),
			Method:       c.Request.Method,
			RelativePath: c.Request.URL.Path,
			UserAgent:    c.Request.UserAgent(),
		}

		switch {
		case c.Writer.Status() >= 500:
			l.loggerService.LogError(entryLog)
		case c.Writer.Status() >= 400:
			l.loggerService.LogWarning(entryLog)
		default:
			l.loggerService.LogInfo(entryLog)
		}
	}
}

func NewLogMiddleware(loggerService logger.MyLogger) LogMiddleware {
	return &logMiddleware{loggerService: loggerService}
}
