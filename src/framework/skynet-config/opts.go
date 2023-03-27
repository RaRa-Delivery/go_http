package skynet_config

import "encoding/json"

type SkynetConfigAgentOpts struct {
	localOnly bool
	env       string
	mappings  MappingsMap
}

func (o *SkynetConfigAgentOpts) SetLocalOnly(localOnly bool) {
	o.localOnly = localOnly
}

func (o *SkynetConfigAgentOpts) SetEnv(env string) {
	o.env = env
}

func (o *SkynetConfigAgentOpts) SetMappings(mappings MappingsMap) {
	o.mappings = mappings
}

type EnvSpec struct {
	Prop     string `json:"prop"`
	Fallback string `json:"fallback"`
}

type MappingsMap struct {
	env map[string]EnvSpec
}

func (m *MappingsMap) SetEnvSpec(spec map[string]EnvSpec) {
	m.env = spec
}

func (m *MappingsMap) SetEnvSpecFromJSONString(jsonString string) {
	var result map[string]EnvSpec
	_ = json.Unmarshal([]byte(jsonString), &result)
	m.env = result
}
