package acase_sdk

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/tmconsulting/acase-golang-sdk/acaseStructs"
)

const apiUrl = "http://test-www.acase.ru/xml/form.jsp"

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
		Action: acaseStructs.AdmUnit1ActionType{
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

func (a *Api) AdmUnit2Request(countryCode int, admUnit1Code string, admUnit1Name string, admUnit2Code string, admUnit2Name string) (*acaseStructs.AdmUnit2ListType, *AcaseResponseError) {
	req := &acaseStructs.AdmUnit2RequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		Action: acaseStructs.AdmUnit2ActionType{
			Name: acaseStructs.List,
			Parameters: acaseStructs.AdmUnit2ActionTypeParameters{
				CountryCode: countryCode,
				AdmUnit1Code: admUnit1Code,
				AdmUnit1Name: admUnit1Code,
				AdmUnit2Code: admUnit2Code,
				AdmUnit2Name: admUnit2Name,
			},
		},
	}
	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseStructs.AdmUnit2ResponseType{}
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

	return &resp.AdmUnit2List, nil
}