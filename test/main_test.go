package test

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tmconsulting/acase-golang-sdk"
	"gopkg.in/h2non/gock.v1"
)

var (
	au = &acaseSdk.Auth{
		BuyerId: "BuyerId",
		UserId: "UserId",
		Password: "Password",
		Language: "Language",
	}
	acApi = acaseSdk.NewApi(*au)
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
	gock.New("http://test-www.acase.ru").Post("/xml/form.jsp").Reply(200).XML(data)
	return nil
}


