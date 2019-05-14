package test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	acaseSdk "github.com/tmconsulting/acase-golang-sdk"
	"gopkg.in/h2non/gock.v1"
)

var (
	au = &acaseSdk.Auth{
		BuyerId:  "BuyerId",
		UserId:   "UserId",
		Password: "Password",
		Language: "ru",
	}
	acApi  = acaseSdk.NewApi(*au, "http://test-www.acase.ru/xml/form.jsp")
	isMock = true
)

func getXml(filename string) ([]byte, error) {
	cwd, _ := os.Getwd()
	fullFileName := filepath.Join(cwd, "data", filename)
	data, err := ioutil.ReadFile(fullFileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getCustomErrorType() *acaseSdk.AcaseResponseError {
	er := &acaseSdk.AcaseResponseError{}
	er = nil
	return er
}

func gockOff() {
	gock.Off()
}

func testRequest(filename string, isError bool) error {
	if isMock || isError {
		data, err := getXml(filename)
		if err != nil {
			return err
		}
		gock.New("http://test-www.acase.ru").Post("/xml/form.jsp").Reply(200).XML(data)
	}
	return nil
}
