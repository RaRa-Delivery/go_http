package skynet_config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
)

func Load(opts SkynetConfigAgentOpts) map[string]interface{} {
	fmt.Println("Loading skynet agent")
	var conf []RemoteConfig
	env := opts.env
	localOnly := opts.localOnly
	log("Loading config for env " + env + ". Local only: " + strconv.FormatBool(localOnly))
	keys := getMapKeys(opts.mappings.env)
	log(keys)
	auth := Auth{}
	config := Config{}
	if localOnly == false {
		// error handled
		token, _, _, baseUrl, err := auth.fetchAuthTokenAndAuthUserInfo()
		if err == nil {
			// make remote prop map
			finalKeys := make([]string, 0, len(keys))
			for _, eKey := range keys {
				finalKeys = append(finalKeys, opts.mappings.env[eKey].Prop)
			}
			conf, err = config.fetchConfigs(token, finalKeys, []string{env}, baseUrl)
		} else {
			fmt.Println(err)
		}
	}
	return buildEnvMap(opts.mappings, conf, keys)
}

func buildEnvMap(mappings MappingsMap, conf []RemoteConfig, keys []string) map[string]interface{} {
	var cMap map[string]interface{} = make(map[string]interface{})
	for _, key := range keys {
		var remoteConfig RemoteConfig
		for _, c := range conf {
			if c.Name == mappings.env[key].Prop {
				remoteConfig = c
				break
			}
		}
		if remoteConfig == (RemoteConfig{}) {
			remoteConfig = RemoteConfig{key, mappings.env[key].Fallback}
		}
		cMap[remoteConfig.Name] = remoteConfig.Value
		os.Setenv(key, fmt.Sprintf("%v", remoteConfig.Value))
	}
	return cMap
}

func getMapKeys(m map[string]EnvSpec) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func log(message interface{}) {
	bold := color.New(color.Bold, color.FgCyan).SprintFunc()
	fmt.Printf(bold("[CONFIG AGENT] %+v\n"), message)
}
