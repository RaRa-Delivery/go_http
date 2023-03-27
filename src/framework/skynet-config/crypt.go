package skynet_config

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Crypt struct {
	GeneratorAlgo   string
	GeneratorSecret string
	IVHex           string
	KeyHex          string
}

func (c *Crypt) Decrypt(cryptText string) (string /*key*/, string /*secret*/, string /*baseUrl*/, error) {
	fmt.Println(c.GeneratorSecret)
	//key := []byte(c.GeneratorSecret)
	key, _ := hex.DecodeString(c.KeyHex)
	ciphertext, _ := hex.DecodeString(cryptText)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", "", "", err
	}
	if len(ciphertext) < aes.BlockSize {
		return "", "", "", errors.New("ciphertext too short")
	}
	iv, _ := hex.DecodeString(c.IVHex)
	if len(ciphertext)%aes.BlockSize != 0 {
		return "", "", "", errors.New("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	var result map[string]map[string]string
	slimJson := strings.TrimSpace(string(ciphertext))
	_ = json.Unmarshal([]byte(slimJson), &result)
	payload := result["payload"]
	return payload["key"], payload["secret"], payload["baseUrl"], nil
}
