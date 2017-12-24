package acaseSdk

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/tmconsulting/acase-golang-sdk/acaseSts"
)

const apiUrl = "http://test-www.acase.ru/xml/form.jsp"

type Api struct {
	BuyerId  string
	UserId   string
	Password string
	Language acaseSts.LanguageTypeEnum
}

func NewApi(auth Auth) *Api {
	var	lang acaseSts.LanguageTypeEnum
	if auth.Language == "RU" {
		lang = acaseSts.Ru
	} else {
		lang = acaseSts.En
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

func (a *Api) AdmUnit1Request(countryCode int, admUnitCode, admUnitName string) (*acaseSts.AdmUnit1ListType, *AcaseResponseError) {
	req := &acaseSts.AdmUnit1RequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		Action: acaseSts.AdmUnit1ActionType{
			Name: acaseSts.List,
			Parameters: acaseSts.AdmUnit1ActionTypeParameters{
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
	resp := &acaseSts.AdmUnit1ResponseType{}
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

func (a *Api) AdmUnit2Request(countryCode int, admUnit1Code, admUnit1Name, admUnit2Code, admUnit2Name string) (*acaseSts.AdmUnit2ListType, *AcaseResponseError) {
	req := &acaseSts.AdmUnit2RequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		Action: acaseSts.AdmUnit2ActionType{
			Name: acaseSts.List,
			Parameters: acaseSts.AdmUnit2ActionTypeParameters{
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
	resp := &acaseSts.AdmUnit2ResponseType{}
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

func (a *Api) CitizenshipListRequest() (*acaseSts.CitizenshipListType, *AcaseResponseError) {
	req := &acaseSts.CitizenshipListRequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
	}
	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.CitizenshipListType{}
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

	return resp, nil
}

func (a *Api) CityDescriptionRequest(cityCode string) (*acaseSts.CityDescriptionType, *AcaseResponseError) {
	req := &acaseSts.CityDescriptionRequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		CityCode: cityCode,
	}
	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.CityDescriptionType{}
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

	return resp, nil
}

func (a *Api) CityListRequest(countryCode, countryName, cityCode, cityName string) (*acaseSts.CityListType, *AcaseResponseError) {
	req := &acaseSts.CityListRequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		CountryCode: countryCode,
		CountryName: countryName,
		CityCode: cityCode,
		CityName: cityName,
	}
	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.CityListType{}
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

	return resp, nil
}