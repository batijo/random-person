package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadData(filename string, model interface{}) error {
	_, err := os.Stat(filename)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &model)
	if err != nil {
		return err
	}
	return nil
}
