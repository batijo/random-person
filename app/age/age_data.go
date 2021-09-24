package age

import (
	"encoding/json"
	"io/ioutil"

	wr "github.com/mroth/weightedrand"
)

type weightData struct {
	Age    int  `json:"age"`
	Weight uint `json:"weight"`
}

var ages *wr.Chooser

func LoadData(filepath string) error {
	var wd []weightData
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &wd)
	if err != nil {
		return err
	}
	ages, err = wr.NewChooser(getChoices(wd)...)
	if err != nil {
		return err
	}
	return nil
}

func fillWeights(wd *[]weightData) {
	var lastWeihgt = weightData{0, 0}
	i := 1
	for _, d := range *wd {
		for j := i; lastWeihgt.Age != d.Age-1; j++ {
			*wd = append(*wd, weightData{
				Age:    j,
				Weight: lastWeihgt.Weight,
			})
			lastWeihgt = weightData{j, lastWeihgt.Weight}
			i = j + 1
		}
		lastWeihgt = weightData{i, d.Weight}
	}
}

func getChoices(data []weightData) []wr.Choice {
	var choices []wr.Choice
	for _, wd := range data {
		choices = append(choices, wr.Choice{Item: wd.Age, Weight: wd.Weight})
	}
	return choices
}
