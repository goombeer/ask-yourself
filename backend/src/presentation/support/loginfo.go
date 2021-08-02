package support

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/goombeer/ask-yourself/backend/src/domain/log"
	"net/http"
)

func LogInfoMiddleware(c *gin.Context) {
	c.Set(log.RequestInfoContextName, log.RequestInfo{
		RequestId:     getRequestId(c),
		RequestURL:    c.Request.RequestURI,
		BrowserUserId: getBrowserUserId(c),
		RemoteAddr:    c.Request.RemoteAddr,
		ClientIP:      c.ClientIP(),
		UserAgent:     c.Request.UserAgent(),
		Referer:       c.Request.Referer(),
	})
	c.Next()
}

func getRequestId(c *gin.Context) string {
	requestId := c.GetHeader("X-Amzn-Trace-Id")
	if requestId != "" {
		return requestId
	}
	requestId = c.GetHeader("X-Request-Id")
	if requestId != "" {
		return requestId
	}
	return uuid.Must(uuid.NewV4()).String()
}

func getBrowserUserId(c *gin.Context) string {
	cookie, err := c.Cookie("browser_user_id")
	if err != nil && err != http.ErrNoCookie {
		panic(fmt.Sprintf("gin cookie access return unexpected error: %s", err.Error()))
	}
	return cookie
}