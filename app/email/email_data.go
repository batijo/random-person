package email

import (
	"encoding/json"
	"io/ioutil"

	wr "github.com/mroth/weightedrand"
)

type weightData struct {
	Data   string `json:"data"`
	Weight uint   `json:"weight"`
}

var (
	templates *wr.Chooser
	domains   *wr.Chooser
)

func loadDataFromFile(filepath string) ([]weightData, error) {
	var wd []weightData
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return wd, err
	}
	err = json.Unmarshal(data, &wd)
	if err != nil {
		return wd, err
	}
	return wd, nil
}

func getChoices(data []weightData) []wr.Choice {
	var choices []wr.Choice
	for _, wd := range data {
		choices = append(choices, wr.Choice{Item: wd.Data, Weight: wd.Weight})
	}
	return choices
}

func LoadData(templatesPath, domainsPath string) error {
	templatesData, err := loadDataFromFile(templatesPath)
	if err != nil {
		return err
	}
	templates, err = wr.NewChooser(getChoices(templatesData)...)
	if err != nil {
		return err
	}
	domainsData, err := loadDataFromFile(domainsPath)
	if err != nil {
		return err
	}
	domains, err = wr.NewChooser(getChoices(domainsData)...)
	if err != nil {
		return err
	}
	return nil
}
