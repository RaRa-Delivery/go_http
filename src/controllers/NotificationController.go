package controllers

import (
	"github.com/RaRa-Delivery/rara-ms-notification/src/constants"
	"github.com/RaRa-Delivery/rara-ms-notification/src/framework"
	"github.com/RaRa-Delivery/rara-ms-notification/src/middleware"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func NotificationApis(appCtx framework.Framework) {

	appCtx.Router.Post(constants.DS+constants.API+constants.DS+constants.V1+constants.DS+constants.NOTIFICATION+constants.DS+constants.SEND, func(c *routing.Context) error {
		resp := middleware.ResultStatus{}
		return middleware.BuildJsonResponse("", resp, fasthttp.StatusOK, c)
	})

	appCtx.Router.Post(constants.DS+constants.API+constants.DS+constants.V1+constants.DS+constants.NOTIFICATION+constants.DS+constants.CALLBACK, func(c *routing.Context) error {

		return nil
	})

}
