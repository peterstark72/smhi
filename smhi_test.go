package smhi_test

import (
	"fmt"

	"github.com/peterstark72/smhi"
)

func ExampleGetPointForecast() {

	fc := smhi.GetPointForecast(55.5174044, 12.9986549)

	fmt.Println(fc)
}
