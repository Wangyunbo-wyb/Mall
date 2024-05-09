package middleware

import (
	"Mall/consts"
	"Mall/pkg/utils/track"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

func Jaeger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//函数从请求头中获取 "uber-trace-id"。这个头部通常包含了父 span 的 trace ID 和 span ID，用于在分布式系统中追踪请求
		traceId := c.GetHeader("uber-trace-id")
		var span opentracing.Span
		if traceId != "" {
			var err error
			span, err = track.GetParentSpan(c.FullPath(), traceId, c.Request.Header)
			if err != nil {
				return
			}
		} else {
			span = track.StartSpan(opentracing.GlobalTracer(), c.FullPath())
		}
		defer span.Finish()
		//span 添加到 Gin 的上下文中，这样在后续的处理函数中就可以获取到这个 span
		c.Set(consts.SpanCTX, opentracing.ContextWithSpan(c, span))
		//调用后续的处理函数，处理请求
		c.Next()
	}
}
