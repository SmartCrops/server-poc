package waterplanner

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "api.openweathermap.org/data/2.5/forecast"
	appid   = "3d5b7e2657ecaa08a9628246422495ff"
	lat     = "-34.603722"
	lng     = "-58.381592"
)

func getWeather(lat, lng string) (weather, error) {
	resp, err := http.Get(fmt.Sprintf("https://%s?lat=%s&lon=%s&appid=%s", baseURL, lat, lng, appid))
	if err != nil {
		return weather{}, fmt.Errorf("failed to call the api: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return weather{}, fmt.Errorf("failed to read the body: %w", err)
	}

	w := weather{}
	if err = json.Unmarshal(data, &w); err != nil {
		return weather{}, fmt.Errorf("failed to unmarshall the body: %w", err)
	}

	return w, nil
}

func (w weather) itWillRainIn24h() bool {
	for i := 0; i < 8; i++ {
		if w.List[0].Weather[0].Main == "Rain" {
			return true
		}
	}
	return false
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
			Speed float64 `json:"speed"`
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
