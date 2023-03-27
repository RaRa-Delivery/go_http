package controllers

import (
	"fmt"

	"github.com/RaRa-Delivery/rara-ms-notification/src/framework"
	routing "github.com/qiangxue/fasthttp-routing"
)

type IndexController struct {
	AppCtx framework.Framework
}

func (c *IndexController) Index(ctx *routing.Context) error {
	fmt.Fprintf(ctx, "Welcome to Boilerplate System! Env: "+c.AppCtx.GetEnv())
	return nil
}
