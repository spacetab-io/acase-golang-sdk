package acaseSdk

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

const apiUrl = "http://test-www.acase.ru/xml/form.jsp"

type Api struct {
	BuyerId  string
	UserId   string
	Password string
	Language LanguageTypeEnum
}

func NewApi(auth Auth) *Api {
	var	lang LanguageTypeEnum
	if auth.Language == "RU" {
		lang = Ru
	} else {
		lang = En
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

func (a *Api) AdmUnit1Request(countryCode int, admUnitCode string, admUnitName string) (*AdmUnit1ListType, *AcaseResponseError) {
	req := &AdmUnit1RequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		Action: AdmUnit1ActionType{
			Name: List,
			Parameters: AdmUnit1ActionTypeParameters{
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
	resp := &AdmUnit1ResponseType{}
	err = xml.Unmarshal(respData, resp)
	FatalError(err)
	if resp.Error.Code != "" {
		rError := make([]RespError, 1)
		rError[0] = RespError{
			Code: resp.Error.Code,
			Message: resp.Error.Description,
		}
		res := ErrorResponse(rError)
		return nil, &res[0]
	}

	return &resp.AdmUnit1List, nil
}

func (a *Api) AdmUnit2Request(countryCode int, admUnit1Code string, admUnit1Name string, admUnit2Code string, admUnit2Name string) (*AdmUnit2ListType, *AcaseResponseError) {
	req := &AdmUnit2RequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		Action: AdmUnit2ActionType{
			Name: List,
			Parameters: AdmUnit2ActionTypeParameters{
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
	resp := &AdmUnit2ResponseType{}
	err = xml.Unmarshal(respData, resp)
	FatalError(err)
	if resp.Error.Code != "" {
		rError := make([]RespError, 1)
		rError[0] = RespError{
			Code: resp.Error.Code,
			Message: resp.Error.Description,
		}
		res := ErrorResponse(rError)
		return nil, &res[0]
	}

	return &resp.AdmUnit2List, nil
}