package skynet_config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const CONFIG_URL = "/api/v1/config/fetch"

type Config struct{}

type RemoteConfig struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type fetchConfigBody struct {
	Env  []string `json:"env"`
	Keys []string `json:"keys"`
}

type configResponseBody struct {
	Config []RemoteConfig `json:"config"`
}

func (c *Config) fetchConfigs(token string, keys []string, envs []string, baseUrl string) ([]RemoteConfig, error) {
	body := fetchConfigBody{}
	body.Keys = keys
	body.Env = envs
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return make([]RemoteConfig, 0, 0), err
	}
	var postBody string = string(bodyBytes)
	client := &http.Client{}
	req, err := http.NewRequest("POST", baseUrl+CONFIG_URL, bytes.NewBufferString(postBody))
	if err != nil {
		return make([]RemoteConfig, 0, 0), err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return make([]RemoteConfig, 0, 0), err
	}
	defer resp.Body.Close()
	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	respBody, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(respBody))
	var result configResponseBody
	slimJson := strings.TrimSpace(string(respBody))
	_ = json.Unmarshal([]byte(slimJson), &result)
	return result.Config, nil
}
