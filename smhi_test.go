package smhi_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/peterstark72/smhi"
)

func TestGetPointForecast(t *testing.T) {

	fc, err := smhi.GetPointForecast(55.5174044, 12.9986549)
	if err != nil {
		t.Errorf("Could not load : %s", err)
	}

	if len(fc) == 0 {
		t.Errorf("Got empty forecast : %s", err)
	}

}

func TestJSON(t *testing.T) {

	fc, err := smhi.GetPointForecast(55.5174044, 12.9986549)
	if err != nil {
		t.Errorf("Could not load : %s", err)
	}

	data, err := json.Marshal(fc)
	if err != nil {
		t.Errorf("Could not marshal JSON : %s", err)
	}

	os.Stdout.Write(data)
}
