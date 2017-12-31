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

func (a *Api) CityListRequest(countryCode, countryName, cityName string, cityCode int) (*acaseSts.CityListType, *AcaseResponseError) {
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

func (a *Api) CountryDescriptionRequest(countryCode int) (*acaseSts.CountryDescriptionType, *AcaseResponseError) {
	req := &acaseSts.CountryDescriptionRequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		CountryCode: countryCode,
	}
	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.CountryDescriptionType{}
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

func (a *Api) ClientCategoryListRequest(categoryCode int, categoryName string) (*acaseSts.ClientCategoryListType, *AcaseResponseError) {
	req := &acaseSts.ClientCategoryListRequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		ClientCategoryCode: categoryCode,
		ClientCategoryName: categoryName,
	}
	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.ClientCategoryListType{}
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

func (a *Api) CountryListRequest(countryCode int, countryName string) (*acaseSts.CountryListType, *AcaseResponseError) {
	req := &acaseSts.CountryListRequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		CountryCode: countryCode,
		CountryName: countryName,
	}
	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.CountryListType{}
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

func (a *Api) CurrencyListRequest(currencyCode int, currencyName, options string) (*acaseSts.CurrencyListResponseType, *AcaseResponseError) {
	req := &acaseSts.CurrencyListRequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		Code: currencyCode,
		Name: currencyName,
	}
	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.CurrencyListResponseType{}
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

func (a *Api) CustomerRequestCreate(fullName, zipCode, address, piAddress, inn, kpp, phone, name, buyerTypeName,
	countryName, cityName string, buyerTypeCode, countryCode, cityCode int) (*acaseSts.CustomerResponseCreateType, *AcaseResponseError) {

	req := &acaseSts.CustomerRequestCreateType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		ActionCreate: acaseSts.ActionCreateType{
			Parameters:acaseSts.CustomerRequestParametersType{
				Customer:&acaseSts.CustomerType{
					FullName:fullName,
					ZipCode:zipCode,
					Address:address,
					PIAddress:piAddress,
					INN:inn,
					KPP:kpp,
					Phone:phone,
					Name:name,
					BuyerType:acaseSts.BuyerTypeType{
						Name:buyerTypeName,
						Code:buyerTypeCode,
					},
					Country:acaseSts.CountryType{
						Code:countryCode,
						Name:countryName,
					},
					City:acaseSts.CityType{
						Code:cityCode,
						Name:cityName,
					},
				},
			},
		},
	}
	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.CustomerResponseCreateType{}
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

func (a *Api) CustomerRequestUpdate(fullName, zipCode, address, piAddress, inn, kpp, phone, name, buyerTypeName,
	countryName, cityName string, customerCode, buyerTypeCode, countryCode, cityCode int) (*acaseSts.CustomerResponseUpdateType, *AcaseResponseError) {

	req := &acaseSts.CustomerRequestUpdateType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		ActionUpdate: acaseSts.ActionUpdateType{
			Parameters:acaseSts.CustomerRequestParametersType{
				Customer:&acaseSts.CustomerType{
					Code:customerCode,
					FullName:fullName,
					ZipCode:zipCode,
					Address:address,
					PIAddress:piAddress,
					INN:inn,
					KPP:kpp,
					Phone:phone,
					Name:name,
					BuyerType:acaseSts.BuyerTypeType{
						Name:buyerTypeName,
						Code:buyerTypeCode,
					},
					Country:acaseSts.CountryType{
						Code:countryCode,
						Name:countryName,
					},
					City:acaseSts.CityType{
						Code:cityCode,
						Name:cityName,
					},
				},
			},
		},
	}
	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.CustomerResponseUpdateType{}
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

func (a *Api) CustomerRequestDelete(customerCode int) (*acaseSts.CustomerResponseDeleteType, *AcaseResponseError) {

	req := &acaseSts.CustomerRequestDeleteType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		ActionDelete: acaseSts.ActionDeleteType{
			Parameters:acaseSts.CustomerRequestParametersType{
				Customer:&acaseSts.CustomerType{
					Code:customerCode,
				},
			},
		},
	}
	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.CustomerResponseDeleteType{}
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

func (a *Api) CustomerRequestList(sort, actualOnly int) (*acaseSts.CustomerResponseListType, *AcaseResponseError) {
	req := &acaseSts.CustomerRequestListType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		ActionList: acaseSts.ActionListType{
			Parameters:acaseSts.CustomerRequestParametersType{
				Sort:sort,
				ShowActualOnly:actualOnly,
			},
		},
	}
	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.CustomerResponseListType{}
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

func (a *Api) CustomerRequestInfo(customerCode int) (*acaseSts.CustomerResponseInfoType, *AcaseResponseError) {
	req := &acaseSts.CustomerRequestInfoType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		ActionInfo: acaseSts.ActionInfoType{
			Parameters:acaseSts.CustomerRequestParametersType{
				CustomerCode:customerCode,
			},
		},
	}
	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.CustomerResponseInfoType{}
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

func (a *Api) HotelAmenityListRequest(hotelAmenityCode int, hotelAmenityName string) (*acaseSts.HotelAmenityListResponseType, *AcaseResponseError) {
	req := &acaseSts.HotelAmenityListRequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		HotelAmenityCode:hotelAmenityCode,
		HotelAmenityName:hotelAmenityName,
	}
	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.HotelAmenityListResponseType{}
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

func (a *Api) HotelDescriptionRequest(hotelCode, currencyCode int) (*acaseSts.HotelDescriptionResponseType, *AcaseResponseError) {
	req := &acaseSts.HotelDescriptionRequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		HotelCode:hotelCode,
		CurrencyCode:currencyCode,
	}
	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.HotelDescriptionResponseType{}
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

func (a *Api) HotelListRequest(hotelCode, countryCode, cityCode, hotelRatingCode int, hotelName, options string) (*acaseSts.HotelListResponseType, *AcaseResponseError) {
	req := &acaseSts.HotelListRequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		HotelCode:hotelCode,
		HotelName:hotelName,
		CountryCode:countryCode,
		CityCode:cityCode,
		HotelRatingCode:hotelRatingCode,
		Options:options,
	}
	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.HotelListResponseType{}
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

func (a *Api) HotelPricingRequest2(productCode, currency, whereToPay, numberOfGuests, meal, numberOfExtraBedsAdult,
    numberOfExtraBedsChild, numberOfExtraBedsInfant int,
	hotel, arrivalDate, departureDate, arrivalTime, departureTime, id, accommodationId string) (*acaseSts.HotelPricingResponse2Type, *AcaseResponseError) {

	req := &acaseSts.HotelPricingRequest2Type{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		Hotel:hotel,
		ProductCode:productCode,
		Currency:currency,
		WhereToPay:whereToPay,
		ArrivalDate:arrivalDate,
		DepartureDate:departureDate,
		ArrivalTime:arrivalTime,
		DepartureTime:departureTime,
		NumberOfGuests:numberOfGuests,
		NumberOfExtraBedsAdult:numberOfExtraBedsAdult,
		NumberOfExtraBedsChild:numberOfExtraBedsChild,
		NumberOfExtraBedsInfant:numberOfExtraBedsInfant,
		Meal:meal,
		Id:id,
		AccommodationId:accommodationId,
	}
	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)

	resp := &acaseSts.HotelPricingResponse2Type{}
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

