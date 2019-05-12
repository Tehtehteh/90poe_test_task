package parser

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"t_task/proto"
)

func createFixtureFile(fileName, fixture string) error {
	fd, err := os.Create(fileName)
	if err != nil {
		return err
	}
	_, err = fd.WriteString(fixture)
	if err != nil {
		return err
	}
	err = fd.Close()
	if err != nil {
		return err
	}
	return nil
}

const fixtureFileName = "fixture.json"

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("cleaing fixture.json in case it has not been cleaned")
	return func(t *testing.T) {
		if _, err := os.Stat(fixtureFileName); err == nil {
			err = os.Remove(fixtureFileName)
			if err != nil {
				t.Log("error removing file:", err)
			}
			t.Log("successfully removed file")
		}
	}
}

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

func TestReadPortsFromFileInvalid(t *testing.T) {
	fixture := `
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
	`
	tearDown := setupTestCase(t)
	defer tearDown(t)
	err := createFixtureFile(fixtureFileName, fixture)
	if err != nil {
		t.Errorf("error creating fixture file: %s", err)
		return
	}
	_, err = ReadPortsFromFile(fixtureFileName, nil)
	if err == nil {
		t.Errorf("should error on malformed JSON. err: %s", err)
		return
	}
}

func TestReadPortsFromFileValid(t *testing.T) {
	const fixtureFileName = "fixture.json"
	const fixture = `
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
	`
	err := createFixtureFile(fixtureFileName, fixture)
	if err != nil {
		t.Errorf("error creating fixture file: %s", err)
		return
	}
	tearDown := setupTestCase(t)
	defer tearDown(t)
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
	result, err := ReadPortsFromFile(fixtureFileName, nil)
	if err != nil {
		t.Errorf("should not error on valid JSON. err: %s", err)
		return
	}
	if !reflect.DeepEqual(result[0], p) {
		t.Errorf("should unmarshal into same entity")
		return
	}
}

func TestReadPortsFromReaderWithCb(t *testing.T) {
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
		ID:      "QEQ",
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
	cb := func(p *proto.Port) error {
		p.ID = "QEQ"
		return nil
	}
	result, err := ReadPortsFromReader(mockReader, cb)
	if err != nil {
		t.Errorf("should not error on valid JSON. err: %s", err)
		return
	}
	if !reflect.DeepEqual(result[0], p) {
		t.Errorf("should have same entity (callback was executed and changed ID to `QEQ`).")
		return
	}
}
