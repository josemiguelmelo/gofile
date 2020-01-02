package gofile

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// HandleFileRead function to apply to the file content
type HandleFileRead func(fileContent []byte) (interface{}, error)

// UsingFile open file and apply function over its content. return some data and error
func UsingFile(fileLocation string, handler HandleFileRead) (interface{}, error) {
	file, err := os.Open(fileLocation)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)

	return handler(byteValue)
}

// ParseJSONFile reads a file and parses it into the provided interface
func ParseJSONFile(file string, resultObj interface{}) error {
	_, err := UsingFile(file, func(fileContent []byte) (interface{}, error) {
		err := json.Unmarshal(fileContent, resultObj)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	return err
}

// WriteJSONFile writes data into a JSON file
func WriteJSONFile(fileLocation string, data interface{}) error {
	fileContent, _ := json.MarshalIndent(data, "", " ")

	return ioutil.WriteFile(fileLocation, fileContent, 0644)
}
