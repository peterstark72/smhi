package main

import (
	"fmt"

	"github.com/peterstark72/smhi"
)

func main() {

	fc := smhi.GetPointForecast(55.5174044, 12.9986549)

	for _, f := range fc[:20] {
		fmt.Printf("%s\t%15s\t%v\n", f.TimeSweden.Format("2006-01-02 15:04"), f.Situation, f.Temperature)
	}

}
