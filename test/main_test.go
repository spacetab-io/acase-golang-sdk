package test

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/h2non/gock.v1"
)

func main() {
	defer gock.Off()
}

func getXml(filename string) (interface{}, error) {
	cwd, _ := os.Getwd()
	fullFileName := filepath.Join(cwd, "data", filename)
	data, err := ioutil.ReadFile(fullFileName)
	if err != nil {
		return nil, err
	}

	var res map[string]map[string]interface{}
	if err := xml.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return data, nil
}

func testRequest(filename string) error {
	data, err := getXml(filename)
	if err != nil {
		return err
	}
	gock.New("").Post("").Reply(200).XML(data)
	return nil
}