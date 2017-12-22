package acase_sdk

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

const apiUrl = ""

type Api struct {
	BuyerId  string
	UserId   string
	Password string
	Language string
}

func NewApi(auth Auth) *Api {
	return &Api{
		BuyerId:  auth.BuyerId,
		UserId:   auth.UserId,
		Password: auth.Password,
		Language: auth.Language,
	}
}

func deSerializeResponse(data []byte) *AcaseResponse {
	res := &AcaseResponse{}
	err := xml.Unmarshal(data, res)
	FatalError(err)
	return res
}

func serializeRequest(reqDetails *RequestDetails) *bytes.Buffer {
	bItem, err := xml.Marshal(reqDetails)
	FatalError(err)
	data := xml.Header + string(bItem)
	return bytes.NewBuffer([]byte(data))
}

func requestInternal(reqDetails *RequestDetails) (*AcaseResponse, *[]AcaseResponseError) {
	data := serializeRequest(reqDetails)

	req, err := http.NewRequest("POST", apiUrl, data)
	FatalError(err)
	req.Header.Add("Content-Type", "application/xml")

	client := &http.Client{}
	resp, err := client.Do(req)
	FatalError(err)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	respDetails := deSerializeResponse(body)
	if respDetails.Errors != nil {
		return nil, ErrorResponse(&respDetails.Errors)
	}
	return respDetails, nil
}
