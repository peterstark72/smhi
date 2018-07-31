package tygelsjo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const TimeLayout = "2006-01-02T15:04:05Z"

var Wsymb2 = map[int]string{
	1:  "Klart",                //Clear sky
	2:  "Mest klart",           //Nearly clear sky
	3:  "Växlande molnighet",   //Variable cloudiness
	4:  "Halvklart",            //Halfclear sky
	5:  "Molnig",               //Cloudy sky
	6:  "Mulet",                //	Overcast
	7:  "Dimma",                //	Fog
	8:  "Lätta Regnskurar",     //	Light rain showers
	9:  "Regnskurar",           //Moderate rain showers
	10: "Mycket regnskurar",    //Heavy rain showers
	11: "Åskstorm",             //	Thunderstorm
	12: "Lätt snöblandat regn", //	Light sleet showers
	13: "Snöblandat",           //Moderate sleet showers
	14: "Snöblandat",           //	Heavy sleet showers
	15: "Snö",                  //Light snow showers
	16: "Snö",                  //Moderate snow showers
	17: "Mycket snö",           //	Heavy snow showers
	18: "Ihållande regn",       //	Light rain
	19: "Regn",                 //Moderate rain
	20: "Mycket regn",          //	Heavy rain
	21: "Åska",                 //Thunder
	22: "Snöblandat",           //	Light sleet
	23: "Snöblandat",           //	Moderate sleet
	24: "Snöblandat",           //Heavy sleet
	25: "Lätt snö",             //Light snowfall
	26: "Snö",                  //Moderate snowfall
	27: "Mycket snö",           //	Heavy snowfall
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Parameter struct {
	Name      string    `json:"name"`
	LevelType string    `json:"leveltype"`
	Level     int       `json:"level"`
	Unit      string    `json:"unit"`
	Values    []float64 `json:"values"`
}

type Item struct {
	ValidTime  string      `json:"validTime"`
	Parameters []Parameter `json:"parameters"`
}

type Forecast struct {
	Geometry   Geometry
	TimeSeries []Item
}

type Weather struct {
	T, TCCMean, Wsymb2, R, WS, Gust float64
}

func GetWeather(client *http.Client) Weather {

	url := "https://opendata-download-metfcst.smhi.se/api/category/pmp3g/version/2/geotype/point/lon/12.9883/lat/55.51776/data.json"

	res, _ := client.Get(url)

	data, _ := ioutil.ReadAll(res.Body)

	var fc Forecast
	json.Unmarshal(data, &fc)

	w := Weather{}
	for _, p := range fc.TimeSeries[0].Parameters {
		switch p.Name {
		case "Wsymb2":
			w.Wsymb2 = p.Values[0]
		case "t":
			w.T = p.Values[0]
		case "tcc_mean":
			w.TCCMean = p.Values[0]
		case "r":
			w.R = p.Values[0]
		case "ws":
			w.WS = p.Values[0]
		case "gust":
			w.Gust = p.Values[0]
		}
	}

	return w
}
