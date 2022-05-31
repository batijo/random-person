package age

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"

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
	fillWeights(&wd)
	rand.Seed(time.Now().UTC().UnixNano())
	ages, err = wr.NewChooser(getChoices(wd)...)
	if err != nil {
		return err
	}
	return nil
}

func fillWeights(wd *[]weightData) {
	var lastWeihgt = weightData{-1, 0}
	for _, d := range *wd {
		for j := lastWeihgt.Age + 1; lastWeihgt.Age != d.Age-1; j++ {
			if d.Age < 0 {
				continue
			}
			*wd = append(*wd, weightData{
				Age:    j,
				Weight: lastWeihgt.Weight,
			})
			lastWeihgt = weightData{j, lastWeihgt.Weight}
		}
		lastWeihgt = weightData{d.Age, d.Weight}
	}
}

func getChoices(data []weightData) []wr.Choice {
	var choices []wr.Choice
	for _, wd := range data {
		choices = append(choices, wr.Choice{Item: wd.Age, Weight: wd.Weight})
	}
	return choices
}
