package waterplanner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	baseUrl = "api.openweathermap.org/data/2.5/forecast"
	appid   = "3d5b7e2657ecaa08a9628246422495ff"
	lat     = "-34.603722"
	lon     = "-58.381592"
)

func main() {
	resp, err := http.Get(fmt.Sprintf("https://%s?lat=%s&lon=%s&appid=%s", baseUrl, lat, lon, appid))
	if err != nil {
		log.Fatal("Error:", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	m := weather{}
	json.Unmarshal(data, &m)

	fmt.Println(m.List[0].Rain.ThreeH) // Rain volume accumulated in last 3 hours (List[1] would be in next 3h etc.)
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
