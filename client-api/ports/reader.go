package ports

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"t_task/proto"
)

type OnPortUnmarshalledCallback func(p proto.Port) error

type PortReader func(dataSource string) ([]proto.Port, error)

// Reads JSON-valid file, parses it and produces array of `proto.Port`
func ReadPortsFromFile(filename string, cb OnPortUnmarshalledCallback) ([]proto.Port, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("error reading from file %s: %s", filename, err)
		return nil, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("error closing file in defer fn: %s", err)
		}
	}()
	reader := bufio.NewReader(file)
	return ReadPortsFromReader(reader, cb)
}

// Creates `json.Decoder` from `reader` and reads from it by chunks.
func ReadPortsFromReader(reader io.Reader, cb OnPortUnmarshalledCallback) ([]proto.Port, error) {

	dec := json.NewDecoder(reader)
	var result []proto.Port
	// Reading first token which will be json.Delim `{`
	t, err := dec.Token()
	if err != nil {
		log.Printf("error reading first token: %s", err)
		return nil, err
	}
	for dec.More() {
		// Reading second token which will be string (code). Eks: `AEAJM`
		t, err = dec.Token()
		if err != nil {
			log.Printf("error reading string code: %s", err)
			return nil, err
		}
		// Buffer here is at position of opening `{`, so it's safe to call unmarshal
		var port proto.Port
		err = dec.Decode(&port)
		if err != nil {
			log.Printf("error decoding to port data type: %s", err)
			return nil, err
		}
		id := fmt.Sprint(t)
		port.ID = id
		if cb != nil {
			// Callback with freshly created `proto.Port` instance.
			err = cb(port)
			if err != nil {
				log.Printf("error executing `onPortUnmarshalledCallback`: %s", err)
			}
		}
		result = append(result, port)
	}
	return result, nil
}
