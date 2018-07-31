package tygelsjo_test

import (
	"net/http"
	"testing"

	"github.com/peterstark72/tygelsjo"
)

func TestForeCast(T *testing.T) {
	tygelsjo.GetWeather(new(http.Client))
}
