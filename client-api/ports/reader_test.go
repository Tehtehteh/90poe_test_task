package ports

import (
	"reflect"
	"strings"
	"testing"

	"t_task/proto"
)

func TestReadPortsFromReader(t *testing.T) {
	mockReader := strings.NewReader(`
		{
		  "0xFFF": {
			"name": "0xFFF",
			"city": "0xFFF",
			"country": "Will not find this",
			"alias": [],
			"regions": [],
			"coordinates": [
			  13.37,
			  13.37
			],
			"province": "0xFFF",
			"timezone": "0xFFF",
			"unlocs": [
			  "0xFFF"
			],
			"code": "1337"
		  }
		}
	`)
	p := proto.Port{
		ID:      "0xFFF",
		Name:    "0xFFF",
		City:    "0xFFF",
		Country: "Will not find this",
		Alias:   []string{},
		Coordinates: []float32{
			13.37,
			13.37,
		},
		Province: "0xFFF",
		Unlocs:   []string{"0xFFF"},
		Code:     "1337",
		Regions:  []string{},
		Timezone: "0xFFF",
	}
	result, err := ReadPortsFromReader(mockReader, nil)
	if err != nil {
		t.Errorf("should not error on valid JSON. err: %s", err)
		return
	}
	if !reflect.DeepEqual(result[0], p) {
		t.Errorf("should unmarshal into same entity")
		return
	}
}

func TestReadPortsFromReaderMalformed(t *testing.T) {
	mockReader := strings.NewReader(`
		{
		  "0xFFF"{/"qe": "0xFFF",
			"city": "0xFFF",
			"country": "Will not find this",
			"alias": [],
			"regions": [],
			"coordinates": [
			  13.37,
			  13.37
			],
			"province": "0xFFF",
			"timezone": "0xFFF",
			"unlocs": [
			  "0xFFF"
			],
			"code": "1337"
		  }
		}
	`)
	_, err := ReadPortsFromReader(mockReader, nil)
	if err == nil {
		t.Errorf("should error on malformed JSON. err: %s", err)
		return
	}
}
