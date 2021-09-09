package email

import (
	"encoding/json"
	"io/ioutil"
)

type weightData struct {
	Data   string `json:"data"`
	Weight uint   `json:"weight"`
}

type weightsData struct {
	WeightsData []weightData `json:"weightsData"`
}

var (
	templates weightsData
	domains   weightsData
)

func (e *weightsData) loadDataFromFile(filepath string) error {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &e)
	if err != nil {
		return err
	}
	return nil
}

func LoadTemplates(filepath string) error {
	return templates.loadDataFromFile(filepath)
}
func LoadDomains(filepath string) error {
	return domains.loadDataFromFile(filepath)
}
