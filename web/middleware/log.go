package middleware

import (
	"bytes"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

func Logger(log *zap.Logger) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		method := ctx.Request().Method
		start := time.Now()
		fields := make([]zap.Field, 0, 13)
		fields = append(fields, zap.String("ip", ctx.Request().RemoteAddr))
		fields = append(fields, zap.String("url", ctx.Request().URL.String()))
		fields = append(fields, zap.String("method", method))
		fields = append(fields, zap.String("proto", ctx.Request().Proto))
		fields = append(fields, zap.String("user_agent", ctx.Request().UserAgent()))
		fields = append(fields, zap.String("x_request_id", ctx.GetHeader("X-Request-Id")))

		if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
			body, err := ioutil.ReadAll(ctx.Request().Body)
			if err == nil {
				defer ctx.Request().Body.Close()
				buf := bytes.NewBuffer(body)
				ctx.Request().Body = ioutil.NopCloser(buf)
				fields = append(fields, zap.Int64("request_content_length", ctx.GetContentLength()))
				fields = append(fields, zap.String("request_body", string(body)))
			}
		}

		ctx.Next()
		useTime := time.Since(start).Nanoseconds() / 1e6
		fields = append(fields, zap.Int64("use_time", useTime))
		fields = append(fields, zap.Int("response_status", ctx.ResponseWriter().StatusCode()))
		if ctx.Values().GetString("out_err") != "" {
			fields = append(fields, zap.String("response_err", ctx.Values().GetString("out_err")))
		}
		fields = append(fields, zap.String("response_length", ctx.ResponseWriter().Header().Get("size")))
		if v := ctx.Values().Get("res_body"); v != nil {
			if b, ok := v.([]byte); ok {
				fields = append(fields, zap.String("response_body", string(b)))
			}
		}
		log.Info("access log", fields...)
	}
}
