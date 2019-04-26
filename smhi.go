//Package smhi wraps the SMHI Open Api into one Go API
package smhi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const GetPointForecastEndPoint = "https://opendata-download-metfcst.smhi.se/api/category/pmp3g/version/2/geotype/point/lon/%s/lat/%s/data.json"

//Coordinate is a GPS coordinate
type Coordinate [2]float64

//Geometry is SMHI Geometry
type Geometry struct {
	Type        string       `json:"type"`
	Coordinates []Coordinate `json:"coordinates"`
}

//Parameter is defined by https://opendata.smhi.se/apidocs/metfcst/parameters.html
type Parameter struct {
	Name      string    `json:"name"`
	LevelType string    `json:"leveltype"`
	Level     int       `json:"level"`
	Unit      string    `json:"unit"`
	Values    []float64 `json:"values"`
}

//Item is en items in a TimeSeries
type Item struct {
	ValidTime  string      `json:"validTime"`
	Parameters []Parameter `json:"parameters"`
}

//Forecast is result from https://opendata.smhi.se/apidocs/metfcst/get-forecast.html
type Forecast struct {
	Geometry   Geometry
	TimeSeries []Item
}

//Weather is simple forecast with a timestamp
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
	Time                  time.Time
	PointLocation         Coordinate
}

//GetPointForecast returns an array of Weather forecasts for given coordinate
func GetPointForecast(lat float64, lon float64) (result []Weather, err error) {

	//Convert lat, lon to str with 5 decimals
	lats, lons := strconv.FormatFloat(lat, 'f', 5, 64), strconv.FormatFloat(lon, 'f', 5, 64)

	//Get API
	url := fmt.Sprintf(GetPointForecastEndPoint, lons, lats)
	res, err := http.Get(url)
	if err != nil {
		return
	}

	//Parse data
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	var fc Forecast
	err = json.Unmarshal(data, &fc)
	if err != nil {
		return
	}

	for _, ts := range fc.TimeSeries {
		var w Weather
		w.Time, _ = time.Parse(time.RFC3339, ts.ValidTime)
		w.PointLocation = fc.Geometry.Coordinates[0]
		for _, p := range ts.Parameters {
			switch p.Name {
			case "Wsymb2":
				w.Wsymb2 = int(p.Values[0])
				w.Situation = Wsymb2[w.Wsymb2][0]
				w.Description = Wsymb2[w.Wsymb2][1]
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
