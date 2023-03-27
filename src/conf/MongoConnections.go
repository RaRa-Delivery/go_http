package conf

import "github.com/RaRa-Delivery/rara-ms-notification/src/framework"

var MongoConnections []framework.MongoConnectionDescription = []framework.MongoConnectionDescription{
	{
		Name:        "oms-staging",
		EnvVarName:  "MONGO_URL_OMS",
		Description: "Connects to 'oms' Mongo Database.",
		CanFail:     true,
	},
	{
		Name:        "dms-staging",
		EnvVarName:  "MONGO_URL_DMS",
		Description: "Connects to 'dms' Mongo Database.",
		CanFail:     true,
	},
}
