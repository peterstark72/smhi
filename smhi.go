package smhi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const TimeLayout = "2006-01-02T15:04:05Z"

const GetPointForecastEndPoint = "https://opendata-download-metfcst.smhi.se/api/category/pmp3g/version/2/geotype/point/lon/%s/lat/%s/data.json"

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
	TotalCloudCover       int
	RelativeHumidity      float64
	WindSpeed             float64
	WindDirection         int
	Temperature           float64
	Wsymb2                int
	Description           string
	Situation             string
	PrecipitationCategory int
	Precipitation         string
	MinPrecipitation      float64
	MaxPrecipitation      float64
	Icon                  string
	Time                  time.Time
}

func GetPointForecast(lat float64, lon float64) (result []Weather) {

	latstr, lonstr := strconv.FormatFloat(lat, 'f', 5, 64), strconv.FormatFloat(lon, 'f', 5, 64)

	url := fmt.Sprintf(GetPointForecastEndPoint, lonstr, latstr)

	res, _ := http.Get(url)

	data, _ := ioutil.ReadAll(res.Body)

	var fc Forecast
	json.Unmarshal(data, &fc)

	for _, ts := range fc.TimeSeries {
		var w Weather
		w.Time, _ = time.Parse(time.RFC3339, ts.ValidTime)
		for _, p := range ts.Parameters {
			switch p.Name {
			case "Wsymb2":
				w.Wsymb2 = int(p.Values[0])
				w.Situation = Wsymb2[w.Wsymb2][0]
				w.Description = Wsymb2[w.Wsymb2][1]
				w.Icon = fmt.Sprintf(Wsymb2Icons, w.Wsymb2)
			case "t":
				w.Temperature = p.Values[0]
			case "tcc_mean":
				w.TotalCloudCover = int(p.Values[0])
			case "r":
				w.RelativeHumidity = p.Values[0]
			case "ws":
				w.WindSpeed = p.Values[0]
			case "pcat":
				w.PrecipitationCategory = int(p.Values[0])
				w.Precipitation = Pcat[w.PrecipitationCategory]
			case "wd":
				w.WindDirection = int(p.Values[0])
			case "pmin":
				w.MinPrecipitation = p.Values[0]
			case "pmax":
				w.MaxPrecipitation = p.Values[0]
			}
		}
		result = append(result, w)
	}

	return
}
