package backends

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func AbiParsor(s string) (abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(s))
	return parsed, err
}

func GetStringFromFile(filepath string, key string) (string, error) {
	obj, err := getContractFromFile(filepath)
	if err != nil {
		return "", err
	}
	b, err := json.MarshalIndent(obj[key], "", "")

	// refine
	// strings.Replace(string(b), "\n", " ", -1)
	s := string(b)
	if s[:1] == "\"" && s[len(s)-1:] == "\"" {
		s = s[1:]
		s = s[:len(s)-1]
	}
	return s, err
}

func getContractFromFile(filepath string) (map[string]interface{}, error) {
	data, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(data)
	var objmap map[string]interface{}
	if err := json.Unmarshal(byteValue, &objmap); err != nil {
		log.Fatal(err)
	}
	return objmap, nil
}

func RemoveHexPrefix(s string) string {
	if s[:2] == "0x" {
		return s[2:]
	} else {
		return s
	}
}
