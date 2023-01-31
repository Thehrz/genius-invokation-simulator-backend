package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/message"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/middleware"
	"github.com/sunist-c/genius-invokation-simulator-backend/util"
)

var (
	captchaRouter *gin.RouterGroup
)

func initCaptchaService() {
	captchaRouter = http.RegisterServices("/captcha")

	captchaRouter.Use(
		append(
			http.EngineMiddlewares,
			middleware.NewQPSLimiter(middlewareConfig),
		)...,
	)

	captchaRouter.GET("",
		middleware.NewInterdictor(middlewareConfig),
		getCaptchaServiceHandler(),
	)
}

func getCaptchaServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, b64s, err := util.GenerateCaptcha()
		if err != nil {
			ctx.AbortWithStatus(500)
		}

		ctx.JSON(200, message.UploadGetCaptchaResponse{
			B64s: b64s,
			ID:   id,
		})
	}
}
