package app

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	Config map[string]map[string]string
	// Store *sessions.CookieStore
)

func Init() error {
	file, err := os.Open("config.json")
	defer file.Close()
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(file)

	json.Unmarshal([]byte(byteValue), &Config)
	gob.Register(map[string]interface{}{})

	return nil
}
