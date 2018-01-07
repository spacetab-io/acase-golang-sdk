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
    numberOfExtraBedsChild, numberOfExtraBedsInfant, hotel  int,
	arrivalDate, departureDate, arrivalTime, departureTime, id, accommodationId string) (*acaseSts.HotelPricingResponse2Type, *AcaseResponseError) {

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

func (a *Api) HotelProductRequest(currency, whereToPay, numberOfGuests, numberOfExtraBedsAdult,
	numberOfExtraBedsChild, numberOfExtraBedsInfant, hotel  int,
	arrivalDate, departureDate, id, accommodationId string) (*acaseSts.HotelProductResponseType, *AcaseResponseError) {

	req := &acaseSts.HotelProductRequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		Hotel:hotel,
		Currency:currency,
		WhereToPay:whereToPay,
		ArrivalDate:arrivalDate,
		DepartureDate:departureDate,
		Id:id,
		AccommodationId:accommodationId,
		NumberOfGuests:numberOfGuests,
		NumberOfExtraBedsAdult:numberOfExtraBedsAdult,
		NumberOfExtraBedsChild:numberOfExtraBedsChild,
		NumberOfExtraBedsInfant:numberOfExtraBedsInfant,
	}
	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)

	resp := &acaseSts.HotelProductResponseType{}
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

func (a *Api) HotelSearchRequest(arrivalDate, departureDate, options, hotelName, destListCode string,
	freeSaleOnly, hotelCategory, currency, whereToPay, numberOfGuests, hotelCode, distance, distTypeCode, distCode, guestsAdults, city int,
	priceFrom, priceTo float64, starCodes, minorAges []int) (*acaseSts.HotelSearchResponseType, *AcaseResponseError) {

	req := &acaseSts.HotelSearchRequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		ArrivalDate:arrivalDate,
		DepartureDate:departureDate,
		City:city,
		PriceFrom:priceFrom,
		PriceTo:priceTo,
		FreeSaleOnly:freeSaleOnly,
		HotelCategory:hotelCategory,
		Currency:currency,
		WhereToPay:whereToPay,
		NumberOfGuests:numberOfGuests,
		Options:options,
		HotelCode:hotelCode,
		HotelName:hotelName,
	}
	if starCodes != nil && len(starCodes) > 0 {
		for _, starCode := range starCodes {
			req.StarList.Items = append(req.StarList.Items, *(&acaseSts.SimpleCodeType{Code:starCode}))
		}
	}
	if destListCode != "" {
		req.DestinationList.Code = destListCode
		req.DestinationList.Distance = distance
		req.DestinationList.TypeCode = acaseSts.DestinationTypeEnum(distTypeCode)
	} else if distance > 0 {
		req.Destination.TypeCode = acaseSts.DestinationTypeEnum(distTypeCode)
		req.Destination.Distance = distance
		req.Destination.Code = distCode
	}

	if guestsAdults > 0 && minorAges != nil && len(minorAges) > 0 {
		req.Guests = &acaseSts.GuestsType{NumberOfAdults:guestsAdults}
		for _, minorAge := range minorAges {
			req.Guests.MinorAgeList.Items = append(req.Guests.MinorAgeList.Items, *(&acaseSts.MinorType{Age:minorAge}))
		}
	}

	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)

	resp := &acaseSts.HotelSearchResponseType{}
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

func (a *Api) MealRequest(mealCode, mealTypeCode []int, mealName []string) (*acaseSts.MealResponseType, *AcaseResponseError) {
	req := &acaseSts.MealRequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
	}

	if mealCode != nil && mealTypeCode != nil && mealName != nil {
		if len(mealCode) > 0 {
			if len(mealCode) != len(mealTypeCode) && len(mealCode) != len(mealName) {
				res := &AcaseResponseError{Code:"0", Message:"Length of parameters lists does not match"}
				return nil, res
			}
		}

		for i, item := range mealCode {
			pt := &acaseSts.ParametersType{MealCode:item, MealTypeCode:mealTypeCode[i], MealName:mealName[i]}
			at := &acaseSts.ActionType{Name:"LIST", Parameters:*pt}
			req.Action = append(req.Action, *at)
		}
	}

	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.MealResponseType{}
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

func (a *Api) MealTypeRequest(mealTypeCode []int, mealName []string) (*acaseSts.MealTypeResponseType, *AcaseResponseError) {
	req := &acaseSts.MealTypeRequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
	}

	if mealTypeCode != nil && mealName != nil {
		if len(mealTypeCode) > 0 {
			if len(mealTypeCode) != len(mealName) {
				res := &AcaseResponseError{Code:"0", Message:"Length of parameters lists does not match"}
				return nil, res
			}
		}

		for i, item := range mealTypeCode {
			pt := &acaseSts.ParametersType{MealTypeCode:item, MealName:mealName[i]}
			at := &acaseSts.ActionType{Name:"LIST", Parameters:*pt}
			req.Action = append(req.Action, *at)
		}
	}

	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.MealTypeResponseType{}
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

func (a *Api) ObjectRequest(objectTypeCode, objectSubTypeCode, cityCode []int) (*acaseSts.ObjectResponseType, *AcaseResponseError) {
	req := &acaseSts.ObjectRequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
	}

	if objectTypeCode != nil && objectSubTypeCode != nil && cityCode != nil {
		if len(objectTypeCode) > 0 {
			if len(objectTypeCode) != len(objectSubTypeCode) &&  len(objectTypeCode) != len(cityCode) {
				res := &AcaseResponseError{Code:"0", Message:"Length of parameters lists does not match"}
				return nil, res
			}
		}

		for i, item := range objectTypeCode {
			pt := &acaseSts.ObjectParametersType{ObjectTypeCode:item, ObjectSubTypeCode:objectSubTypeCode[i], CityCode:cityCode[i]}
			at := &acaseSts.ObjectActionType{Name:"LISTMETROSTYLE", Parameters:*pt}
			req.Action = append(req.Action, *at)
		}
	}

	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.ObjectResponseType{}
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

func (a *Api) ObjectSubTypeRequest(objectSubTypeCode []int) (*acaseSts.ObjectSubTypeResponseType, *AcaseResponseError) {
	req := &acaseSts.ObjectSubTypeRequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
	}

	if objectSubTypeCode != nil && len(objectSubTypeCode) > 0 {
		for _, item := range objectSubTypeCode {
			pt := &acaseSts.ObjSubTypeParamType{ObjectTypeCode:item}
			at := &acaseSts.ObjSubTypeActionType{Name:"LIST", Parameters:*pt}
			req.Action = append(req.Action, *at)
		}
	}

	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.ObjectSubTypeResponseType{}
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

func (a *Api) ObjectTypeRequest(objectTypeCode []int) (*acaseSts.ObjectTypeResponseType, *AcaseResponseError) {
	req := &acaseSts.ObjectTypeRequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
	}

	if objectTypeCode != nil && len(objectTypeCode) > 0 {
		for _, item := range objectTypeCode {
			pt := &acaseSts.ParametersObjType{ObjectTypeCode:item}
			at := &acaseSts.ObjectTypeActionType{Name:"LIST", Parameters:*pt}
			req.Action = append(req.Action, *at)
		}
	}

	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.ObjectTypeResponseType{}
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

func (a *Api) ObjTypeListRequest(objTypeCode, objTypeName string) (*acaseSts.ObjTypeListResponseType, *AcaseResponseError) {
	req := &acaseSts.ObjTypeListRequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
		Code:objTypeCode,
		Name:objTypeName,
	}
	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.ObjTypeListResponseType{}
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

func (a *Api) OrderDocRequest(actionName acaseSts.OrderDocActionName, taskId, docId, code int) (*acaseSts.OrderDocsResponseType, *AcaseResponseError) {
	req := &acaseSts.OrderDocsRequestType{
		Language: a.Language,
		Password: a.Password,
		UserId: a.UserId,
		BuyerId: a.BuyerId,
	}
	req.Action = make([]acaseSts.OrderDocActionType, 1)
	req.Action[0].Name = string(actionName)

	switch actionName {
		case acaseSts.TASKADD:
			req.Action[0].Parameters.DocId = docId
			req.Action[0].Parameters.DocType = &acaseSts.OrderDocParamsDocType{Code:code}
		case acaseSts.TASKSTATUS, acaseSts.TASKRESPONSE:
			req.Action[0].Parameters.TaskId = taskId
	}

	bItem, err := xml.Marshal(req)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.OrderDocsResponseType{}
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

func (a *Api) OrderInfoNotifyRequest(item *acaseSts.OrderInfoNotifyRequestType) (*acaseSts.OrderInfoNotifyResponseType, *AcaseResponseError) {
	item.Password = a.Password
	bItem, err := xml.Marshal(item)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.OrderInfoNotifyResponseType{}
	err = xml.Unmarshal(respData, resp)
	FatalError(err)
	if resp.Error.Code != "" || resp.Error.Description != "" {
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

func (a *Api) OrderInfoAwocNotifyRequest(item *acaseSts.OrderInfoAwocNotifyRequestType) (*acaseSts.OrderInfoNotifyResponseType, *AcaseResponseError) {
	bItem, err := xml.Marshal(item)
	FatalError(err)
	respData, err := requestInternal([]byte(xml.Header + string(bItem)))
	FatalError(err)
	resp := &acaseSts.OrderInfoNotifyResponseType{}
	err = xml.Unmarshal(respData, resp)
	FatalError(err)
	if resp.Error.Code != "" || resp.Error.Description != "" {
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

