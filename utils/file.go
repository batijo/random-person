package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadData(filename string, model interface{}) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
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
