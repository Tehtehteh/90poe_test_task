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

type PortReader func(dataSource string) ([]proto.Port, error)

func ReadPortsFromFile(filename string) ([]proto.Port, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("error reading from file %s: %s", filename, err)
		return nil, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("error closing file in defer callback: %s", err)
		}
	}()
	reader := bufio.NewReader(file)
	return ReadPortsFromReader(reader)
}

func ReadPortsFromReader(reader io.Reader) ([]proto.Port, error) {

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
		// Buffer here at position of opening `{`
		var port proto.Port
		err = dec.Decode(&port)
		if err != nil {
			log.Printf("error decoding to port data type: %s", err)
			return nil, err
		}
		id := fmt.Sprint(t)
		port.ID = id
		result = append(result, port)
	}
	return result, nil
}
