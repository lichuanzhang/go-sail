package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/keepchen/go-sail/v3/lib/logger"
	"go.uber.org/zap"
)

// RequestEntry 请求入口
//
// 注入内容：
//
// 1.请求id
//
// 2.请求到来那一刻的纳秒时间戳
//
// 3.包装了请求id的日志组件实例
//
// 4.客户端语言代码
//
// 作用是在请求入口注入必要的内容到上下文，供后续的请求调用链使用，一般用于日志追踪、链路追踪
func RequestEntry() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			requestId string
			spanId    string
		)
		requestIdInHeader := c.Request.Header.Get("requestId")
		if len(requestIdInHeader) == 0 {
			requestIdInHeader = c.Request.Header.Get("X-Request-Id")
		}
		if len(requestIdInHeader) > 0 {
			requestId = requestIdInHeader
			spanId = uuid.New().String()
		} else {
			requestId = uuid.New().String()
			spanId = requestId
		}
		c.Set("requestId", requestId)
		c.Set("spanId", spanId)
		c.Set("entryAt", time.Now().UnixNano())
		c.Set("logger", logger.GetLogger().With(zap.String("requestId", requestId),
			zap.String("spanId", spanId)))

		//解析客户端语言并注入上下文
		var language = "en"
		al := c.Request.Header.Get("accept-language")
		if len(al) > 0 {
			//example: zh-CN,zh;q=0.9,en;q=0.8,ja;q=0.7
			als := strings.Split(al, ",")
			if len(als) > 0 {
				language = als[0]
			}
		}
		c.Set("language", language)

		c.Next()
	}
}
