package conf

import (
	"fmt"
	"log"

	"github.com/RaRa-Delivery/rara-ms-notification/src/controllers"
	"github.com/RaRa-Delivery/rara-ms-notification/src/framework"
)

func Bootstrap(appCtx framework.Framework) {
	log.Println("Running Bootstrap...")

	controllers.NotificationApis(appCtx)
	fmt.Println("App is ready!")
}
