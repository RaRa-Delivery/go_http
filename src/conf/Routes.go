package conf

import (
	"log"

	"github.com/RaRa-Delivery/rara-ms-notification/src/controllers"
	"github.com/RaRa-Delivery/rara-ms-notification/src/framework"
	"github.com/newrelic/go-agent/v3/newrelic"
	routing "github.com/qiangxue/fasthttp-routing"
)

func Route(appCtx framework.Framework) {

	// Init controllers
	var indexController = controllers.IndexController{AppCtx: appCtx}

	//// init new relic
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("BMS"),
		newrelic.ConfigLicense("eu01xx0d986115c7cc37a2e6490d28d1621cNRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err == nil {
		log.Println("--------------NEW RELIC ERROR-------------------")
		log.Println(err)
		log.Println("-------------END NEW RELIC ERROR----------------")
	}

	appCtx.Router.Get("/", func(ctx *routing.Context) error {
		txn := app.StartTransaction("index_route")
		defer txn.End()
		return indexController.Index(ctx)
	})

}
