package waterplanner

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseUrl = "api.openweathermap.org/data/2.5/forecast"
	appid   = "3d5b7e2657ecaa08a9628246422495ff"
)

func GetAccumulatedRainVolume(hours int, lat float64, lon float64) (float64, error) {
	if hours < 0 || hours > 120 {
		return 0, errors.New("Weather data available for 0 - 120 hours!")
	}

	resp, err := http.Get(fmt.Sprintf("https://%s?lat=%f&lon=%f&appid=%s", baseUrl, lat, lon, appid))
	if err != nil {
		return 0, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	m := weather{}
	err = json.Unmarshal(data, &m)

	index := hours / 3
	accumulatedRainVolume := 0.0

	for i := 0; i < index; i++ {
		accumulatedRainVolume += m.List[i].Rain.ThreeH
	}

	return accumulatedRainVolume, nil
}

type weather struct {
	Cod     string `json:"cod"`
	Message int    `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []struct {
		Dt   int `json:"dt"`
		Main struct {
			Temp      float64 `json:"temp"`
			FeelsLike float64 `json:"feels_like"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  int     `json:"pressure"`
			SeaLevel  int     `json:"sea_level"`
			GrndLevel int     `json:"grnd_level"`
			Humidity  int     `json:"humidity"`
			TempKf    float64 `json:"temp_kf"`
		} `json:"main"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Wind struct {
			Speed int     `json:"speed"`
			Deg   int     `json:"deg"`
			Gust  float64 `json:"gust"`
		} `json:"wind"`
		Visibility int `json:"visibility"`
		Pop        int `json:"pop"`
		Sys        struct {
			Pod string `json:"pod"`
		} `json:"sys"`
		DtTxt string `json:"dt_txt"`
		Rain  struct {
			ThreeH float64 `json:"3h"`
		} `json:"rain,omitempty"`
	} `json:"list"`
	City struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"coord"`
		Country    string `json:"country"`
		Population int    `json:"population"`
		Timezone   int    `json:"timezone"`
		Sunrise    int    `json:"sunrise"`
		Sunset     int    `json:"sunset"`
	} `json:"city"`
}
