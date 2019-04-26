package main

import (
	"fmt"
	"time"

	"github.com/peterstark72/smhi"
)

func main() {

	cet, _ := time.LoadLocation("Europe/Copenhagen")

	fc, err := smhi.GetPointForecast(55.5174044, 12.9986549)
	if err != nil {
		return
	}

	for _, f := range fc[:20] {
		fmt.Printf("%s\t%15s\t%v°C\t%s\n", f.Time.In(cet).Format("2006-01-02 15:04"), f.Situation, f.Temperature, f.Precipitation)
	}

}
