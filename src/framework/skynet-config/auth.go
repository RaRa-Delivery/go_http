package skynet_config

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
	"time"
)

const AUTH_FILE = "auth.crypt"
const URL_FILE = "base.conf"
const AGENT_GRANT_TYPE = "config:agent"
const SKYNET_BASE_URL = "https://skynet.classplusapp.com"
const HANDSHAKE_URL = "/oauth2/v1/token"
const GENERATOR_ALGO = "aes-192-cbc"
const GENERATOR_SECRET = "9834hdbfwhebf3y84gr32yi4fgkh3frwfbldn2h3o2r81y487rt434yogfwhbaf982347o21efbwhvfwevfuwevfy3fg8g"
const GENERATOR_IVH = "39fdffe3c15bc07ac1bdd8b8dd759c8f"
const GENERATOR_KEYH = "5d6a137ec0f240d83bc29bf2c1c6182cd50db37e72d31253"

type Auth struct {
	key     string
	secret  string
	baseUrl string
	rootDir string
	token   string
	scopes  []string
	name    string
}

func (a *Auth) fetchAuthTokenAndAuthUserInfo() (string /*token*/, []string /*scopes*/, string /*name*/, string /*baseUrl*/, error) {
	// default
	a.baseUrl = SKYNET_BASE_URL
	err := a.fetchAuthConfig()
	if err != nil {
		return "", make([]string, 0, 0), "", "", err
	}
	if len(a.secret) == 0 {
		return "", make([]string, 0, 0), "", "", errors.New("Not logged in! Hence can not fetch.")
	}
	err = a.testConnection()
	if err != nil {
		return "", make([]string, 0, 0), "", "", err
	}
	return a.token, a.scopes, a.name, a.baseUrl, nil
}

func (a *Auth) getBaseDir() string {
	_ = a.getHomeDir()
	return path.Join(a.rootDir, ".skynet")
}

func (a *Auth) getAuthDir() string {
	return path.Join(a.getBaseDir(), "auth")
}

func (a *Auth) getBaseUrlConfFilePath() string {
	return path.Join(a.getBaseDir(), URL_FILE)
}

func (a *Auth) getAuthFilePath() string {
	return path.Join(a.getAuthDir(), AUTH_FILE)
}

func (a *Auth) getRawConfig() (string, error) {
	var configRaw string
	filePath := a.getAuthFilePath()
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	configRaw = string(data)
	return configRaw, nil
}

func (a *Auth) fetchAuthConfig() error {
	raw, err := a.getRawConfig()
	if err != nil {
		return err
	}
	if len(raw) > 10 { // min 10 chars for valid config
		crypt := Crypt{
			GeneratorAlgo:   GENERATOR_ALGO,
			GeneratorSecret: GENERATOR_SECRET,
			IVHex:           GENERATOR_IVH,
			KeyHex:          GENERATOR_KEYH,
		}
		key, secret, baseUrl, err := crypt.Decrypt(raw)
		if err != nil {
			return err
		}
		a.key = key
		a.secret = secret
		a.baseUrl = baseUrl
		return nil
	} else {
		return errors.New("Not Logged In Yet! Please authorize the machine first. Will roll back to falbacks.")
	}
}

func (a *Auth) testConnection() error {
	err := a.handshake()
	if err != nil {
		return err
	}
	fmt.Println("Logged In With Client: " + a.name)
	return nil
}

func (a *Auth) handshake() error {
	url := a.baseUrl + HANDSHAKE_URL
	date := time.Now().Unix() * 1000
	grantType := AGENT_GRANT_TYPE
	seconds := 3600
	checkStr := fmt.Sprintf("%v:%v:%v:%v:%v", grantType, a.key, seconds, date, a.secret)
	mac := hmac.New(sha256.New, []byte(a.secret))
	mac.Write([]byte(checkStr))
	assertion := hex.EncodeToString(mac.Sum(nil))
	// Make Call
	body := handshakeAuthRequestBody{
		GrantType:    grantType,
		ApiKey:       a.key,
		GrantSeconds: seconds,
		Ts:           date,
		Assertion:    assertion,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}
	var postBody string = string(bodyBytes)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(postBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//fmt.Println("H response Status:", resp.Status)
	//fmt.Println("H response Headers:", resp.Header)
	respBody, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("H response Body:", string(respBody))
	var result handshakeAuthResponseBody
	slimJson := strings.TrimSpace(string(respBody))
	_ = json.Unmarshal([]byte(slimJson), &result)
	a.token = result.Token
	a.scopes = result.Scopes
	a.name = result.Name
	return nil
}

func (a *Auth) getHomeDir() error {
	dir, err := homedir.Dir()
	if err != nil {
		return err
	}
	a.rootDir = dir
	return nil
}

type handshakeAuthRequestBody struct {
	GrantType    string `json:"grant_type"`
	ApiKey       string `json:"api_key"`
	GrantSeconds int    `json:"grant_seconds"`
	Ts           int64  `json:"ts"`
	Assertion    string `json:"assertion"`
}
type handshakeAuthResponseBody struct {
	Token     string   `json:"token"`
	Name      string   `json:"name"`
	Scopes    []string `json:"scopes"`
	Expiry    int64    `json:"expiry"`
	TokenType string   `json:"token_type"`
}
