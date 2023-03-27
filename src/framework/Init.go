package framework

import (
	"log"

	skynet_config "github.com/RaRa-Delivery/rara-ms-notification/src/framework/skynet-config"
)

func Init(flavour string, envJson string, appBaseDir string) Framework {
	framework := Framework{flavour: flavour, appBaseDir: appBaseDir}
	log.Println("Initializing...")
	framework.LoadEnv()

	// load env
	framework.CoreLog(framework.GetEnv())
	framework.CoreLog(framework.GetFlavour())

	// Load config fetching agent
	mappingsMap := skynet_config.MappingsMap{}
	mappingsMap.SetEnvSpecFromJSONString(envJson)
	confAgentOpts := skynet_config.SkynetConfigAgentOpts{}
	confAgentOpts.SetEnv(framework.GetEnv())
	confAgentOpts.SetLocalOnly(false)
	confAgentOpts.SetMappings(mappingsMap)
	framework.SetConfig(skynet_config.Load(confAgentOpts))

	// load http
	framework.LoadPort()
	httpSystem := HTTPSystem{
		Port: framework.port,
	}
	httpSystem.init()
	framework.setHTTPSystem(httpSystem)

	return framework
}
