package acase_sdk

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/tmconsulting/acase-golang-sdk/acaseStructs"
)

const apiUrl = ""

type Api struct {
	BuyerId  string
	UserId   string
	Password string
	Language acaseStructs.LanguageTypeEnum
}

func NewApi(auth Auth) *Api {
	var	lang acaseStructs.LanguageTypeEnum
	if auth.Language == "RU" {
		lang = acaseStructs.Ru
	} else {
		lang = acaseStructs.En
	}
	return &Api{
		BuyerId:  auth.BuyerId,
		UserId:   auth.UserId,
		Password: auth.Password,
		Language: lang,
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

func requestInternal(data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer([]byte(data)))
	FatalError(err)
	req.Header.Add("Content-Type", "application/xml")

	client := &http.Client{}
	resp, err := client.Do(req)
	FatalError(err)
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (a *Api) AdmUnit1Request(countryCode int, admUnitCode string, admUnitName string) (*acaseStructs.AdmUnit1ListType, *AcaseResponseError) {
	req := &acaseStructs.AdmUnit1RequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		Action: acaseStructs.ActionType{
			Name: acaseStructs.List,
			Parameters: acaseStructs.AdmUnit1ActionTypeParameters{
				CountryCode: countryCode,
				AdmUnit1Code: admUnitCode,
				AdmUnit1Name: admUnitCode,
			},
		},
	}
	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseStructs.AdmUnit1ResponseType{}
	err = xml.Unmarshal(respData, resp)
	FatalError(err)
	if resp.Error.Code != 0 && resp.Error.Description != "" {
		rError := make([]RespError, 1)
		rError[0] = RespError{
			Code: string(resp.Error.Code),
			Message: resp.Error.Description,
		}
		res := ErrorResponse(rError)
		return nil, &res[0]
	}

	return &resp.AdmUnit1List, nil
}