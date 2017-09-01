package main

import (
	"encoding/json"
	"io/ioutil"
)

func readConfig(file string, cfg interface{}) error {
	cfgbuf, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(cfgbuf, &cfg)
	if err != nil {
		return err
	}
	return err
}
