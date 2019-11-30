package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Weather struct {
	CurrentCondition []Condition `json:"current_condition"`
}

type Condition struct {
	TempC      string `json:"temp_C"`
	TempF      string `json:"temp_F"`
	FeelsLikeC string `json:"FeelsLikeC"`
	FeelsLikeF string `json:"FeelsLikeF"`
}

func GetWeather(location string) (Weather, error) {
	url := fmt.Sprintf("https://www.wttr.in/%s?format=j1", location)
	res, err := http.Get(url)
	if err != nil {
		return Weather{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Weather{}, err
	}

	var wttr Weather
	if err = json.Unmarshal(body, &wttr); err != nil {
		return Weather{}, err
	}

	if len(wttr.CurrentCondition) == 0 {
		return Weather{}, errors.New("no data found")
	}

	return wttr, nil
}
