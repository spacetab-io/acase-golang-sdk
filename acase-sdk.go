package acaseSdk

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/tmconsulting/acase-golang-sdk/acaseSts"
	"log"
)

type Api struct {
	BuyerId		string
	UserId		string
	Password	string
	Language	acaseSts.LanguageTypeEnum
	ApiUrl		string
}

func NewApi(auth Auth, apiUrl string) *Api {
	var	lang acaseSts.LanguageTypeEnum
	if auth.Language == "RU" {
		lang = acaseSts.Ru
	} else {
		lang = acaseSts.En
	}
	return &Api{
		BuyerId:auth.BuyerId,
		UserId:auth.UserId,
		Password:auth.Password,
		Language:lang,
		ApiUrl:apiUrl,
	}
}

func (a *Api) requestInternal(data []byte) (*[]byte, error) {
	req, err := http.NewRequest("POST", a.ApiUrl, bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Print(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/xml")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return &res, nil
}

func (a *Api) processRequest(item interface{}, res interface{}) *AcaseResponseError {
	bItem, err := xml.Marshal(item)
	respData, err := a.requestInternal([]byte(xml.Header + string(bItem)))
	if err != nil {
		log.Print(err)
		return &AcaseResponseError{
			Code:"Unknown",
			Message:err.Error(),
		}
	}
	err = xml.Unmarshal(*respData, res)
	if err != nil {
		log.Print(err)
		return &AcaseResponseError{
			Code:"Unknown",
			Message:err.Error(),
		}
	}
	return nil
}

func (a *Api) AdmUnit1Request(countryCode int, admUnitCode, admUnitName string) (*acaseSts.AdmUnit1ListType, *AcaseResponseError) {
	req := &acaseSts.AdmUnit1RequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
		Action: acaseSts.AdmUnit1ActionType{
			Name: acaseSts.List,
			Parameters: acaseSts.AdmUnit1ActionTypeParameters{
				CountryCode: countryCode,
				AdmUnit1Code: admUnitCode,
				AdmUnit1Name: admUnitCode,
			},
		},
	}
	resp := &acaseSts.AdmUnit1ResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}
	return &resp.AdmUnit1List, nil
}

func (a *Api) AdmUnit2Request(countryCode int, admUnit1Code, admUnit1Name, admUnit2Code, admUnit2Name string) (*acaseSts.AdmUnit2ListType, *AcaseResponseError) {
	req := &acaseSts.AdmUnit2RequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
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

	resp := &acaseSts.AdmUnit2ResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return &resp.AdmUnit2List, nil
}

func (a *Api) CitizenshipListRequest() (*acaseSts.CitizenshipListType, *AcaseResponseError) {
	req := &acaseSts.CitizenshipListRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
	}
	resp := &acaseSts.CitizenshipListType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) CityDescriptionRequest(cityCode string) (*acaseSts.CityDescriptionType, *AcaseResponseError) {
	req := &acaseSts.CityDescriptionRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
		CityCode: cityCode,
	}
	resp := &acaseSts.CityDescriptionType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) CityListRequest(countryCode, countryName, cityName string, cityCode int) (*acaseSts.CityListType, *AcaseResponseError) {
	req := &acaseSts.CityListRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
		CountryCode: countryCode,
		CountryName: countryName,
		CityCode: cityCode,
		CityName: cityName,
	}
	resp := &acaseSts.CityListType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) CountryDescriptionRequest(countryCode int) (*acaseSts.CountryDescriptionType, *AcaseResponseError) {
	req := &acaseSts.CountryDescriptionRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
		CountryCode: countryCode,
	}
	resp := &acaseSts.CountryDescriptionType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) ClientCategoryListRequest(categoryCode int, categoryName string) (*acaseSts.ClientCategoryListType, *AcaseResponseError) {
	req := &acaseSts.ClientCategoryListRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
		ClientCategoryCode: categoryCode,
		ClientCategoryName: categoryName,
	}
	resp := &acaseSts.ClientCategoryListType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) CountryListRequest(countryCode int, countryName string) (*acaseSts.CountryListType, *AcaseResponseError) {
	req := &acaseSts.CountryListRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
		CountryCode: countryCode,
		CountryName: countryName,
	}
	resp := &acaseSts.CountryListType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) CurrencyListRequest(currencyCode int, currencyName, options string) (*acaseSts.CurrencyListResponseType, *AcaseResponseError) {
	req := &acaseSts.CurrencyListRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
		Code: currencyCode,
		Name: currencyName,
	}
	resp := &acaseSts.CurrencyListResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) CustomerRequestCreate(fullName, zipCode, address, piAddress, inn, kpp, phone, name, buyerTypeName,
	countryName, cityName string, buyerTypeCode, countryCode, cityCode int) (*acaseSts.CustomerResponseCreateType, *AcaseResponseError) {

	req := &acaseSts.CustomerRequestCreateType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
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
					BuyerType:acaseSts.SimpleCodeNameType{
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
	resp := &acaseSts.CustomerResponseCreateType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) CustomerRequestUpdate(fullName, zipCode, address, piAddress, inn, kpp, phone, name, buyerTypeName,
	countryName, cityName string, customerCode, buyerTypeCode, countryCode, cityCode int) (*acaseSts.CustomerResponseUpdateType, *AcaseResponseError) {

	req := &acaseSts.CustomerRequestUpdateType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
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
					BuyerType:acaseSts.SimpleCodeNameType{
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
	resp := &acaseSts.CustomerResponseUpdateType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) CustomerRequestDelete(customerCode int) (*acaseSts.CustomerResponseDeleteType, *AcaseResponseError) {

	req := &acaseSts.CustomerRequestDeleteType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
		ActionDelete: acaseSts.ActionDeleteType{
			Parameters:acaseSts.CustomerRequestParametersType{
				Customer:&acaseSts.CustomerType{
					Code:customerCode,
				},
			},
		},
	}
	resp := &acaseSts.CustomerResponseDeleteType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) CustomerRequestList(sort, actualOnly int) (*acaseSts.CustomerResponseListType, *AcaseResponseError) {
	req := &acaseSts.CustomerRequestListType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
		ActionList: acaseSts.ActionListType{
			Parameters:acaseSts.CustomerRequestParametersType{
				Sort:sort,
				ShowActualOnly:actualOnly,
			},
		},
	}
	resp := &acaseSts.CustomerResponseListType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) CustomerRequestInfo(customerCode int) (*acaseSts.CustomerResponseInfoType, *AcaseResponseError) {
	req := &acaseSts.CustomerRequestInfoType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
		ActionInfo: acaseSts.ActionInfoType{
			Parameters:acaseSts.CustomerRequestParametersType{
				CustomerCode:customerCode,
			},
		},
	}
	resp := &acaseSts.CustomerResponseInfoType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) HotelAmenityListRequest(hotelAmenityCode int, hotelAmenityName string) (*acaseSts.HotelAmenityListResponseType, *AcaseResponseError) {
	req := &acaseSts.HotelAmenityListRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
		HotelAmenityCode:hotelAmenityCode,
		HotelAmenityName:hotelAmenityName,
	}
	resp := &acaseSts.HotelAmenityListResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) HotelDescriptionRequest(hotelCode, currencyCode int) (*acaseSts.HotelDescriptionResponseType, *AcaseResponseError) {
	req := &acaseSts.HotelDescriptionRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
		HotelCode:hotelCode,
		CurrencyCode:currencyCode,
	}
	resp := &acaseSts.HotelDescriptionResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) HotelListRequest(hotelCode, countryCode, cityCode, hotelRatingCode int, hotelName, options string) (*acaseSts.HotelListResponseType, *AcaseResponseError) {
	req := &acaseSts.HotelListRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
		HotelCode:hotelCode,
		HotelName:hotelName,
		CountryCode:countryCode,
		CityCode:cityCode,
		HotelRatingCode:hotelRatingCode,
		Options:options,
	}
	resp := &acaseSts.HotelListResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) HotelPricingRequest2(productCode, currency, whereToPay, numberOfGuests, meal, numberOfExtraBedsAdult,
    numberOfExtraBedsChild, numberOfExtraBedsInfant, hotel  int,
	arrivalDate, departureDate, arrivalTime, departureTime, id, accommodationId string) (*acaseSts.HotelPricingResponse2Type, *AcaseResponseError) {

	req := &acaseSts.HotelPricingRequest2Type{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
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
	resp := &acaseSts.HotelPricingResponse2Type{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) HotelProductRequest(currency, whereToPay, numberOfGuests, numberOfExtraBedsAdult,
	numberOfExtraBedsChild, numberOfExtraBedsInfant, hotel  int,
	arrivalDate, departureDate, id, accommodationId string) (*acaseSts.HotelProductResponseType, *AcaseResponseError) {

	req := &acaseSts.HotelProductRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
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
	resp := &acaseSts.HotelProductResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) HotelSearchRequest(arrivalDate, departureDate, options, hotelName, destListCode string,
	freeSaleOnly, hotelCategory, currency, whereToPay, numberOfGuests, hotelCode, distance, distTypeCode, distCode, guestsAdults, city int,
	priceFrom, priceTo float64, starCodes, minorAges []int) (*acaseSts.HotelSearchResponseType, *AcaseResponseError) {

	req := &acaseSts.HotelSearchRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
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
	resp := &acaseSts.HotelSearchResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) MealRequest(mealCode, mealTypeCode []int, mealName []string) (*acaseSts.MealResponseType, *AcaseResponseError) {
	req := &acaseSts.MealRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
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
	resp := &acaseSts.MealResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) MealTypeRequest(mealTypeCode []int, mealName []string) (*acaseSts.MealTypeResponseType, *AcaseResponseError) {
	req := &acaseSts.MealTypeRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
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
	resp := &acaseSts.MealTypeResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) ObjectRequest(objectTypeCode, objectSubTypeCode, cityCode []int) (*acaseSts.ObjectResponseType, *AcaseResponseError) {
	req := &acaseSts.ObjectRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
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
	resp := &acaseSts.ObjectResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) ObjectSubTypeRequest(objectSubTypeCode []int) (*acaseSts.ObjectSubTypeResponseType, *AcaseResponseError) {
	req := &acaseSts.ObjectSubTypeRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
	}

	if objectSubTypeCode != nil && len(objectSubTypeCode) > 0 {
		for _, item := range objectSubTypeCode {
			pt := &acaseSts.ObjSubTypeParamType{ObjectTypeCode:item}
			at := &acaseSts.ObjSubTypeActionType{Name:"LIST", Parameters:*pt}
			req.Action = append(req.Action, *at)
		}
	}
	resp := &acaseSts.ObjectSubTypeResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) ObjectTypeRequest(objectTypeCode []int) (*acaseSts.ObjectTypeResponseType, *AcaseResponseError) {
	req := &acaseSts.ObjectTypeRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
	}

	if objectTypeCode != nil && len(objectTypeCode) > 0 {
		for _, item := range objectTypeCode {
			pt := &acaseSts.ParametersObjType{ObjectTypeCode:item}
			at := &acaseSts.ObjectTypeActionType{Name:"LIST", Parameters:*pt}
			req.Action = append(req.Action, *at)
		}
	}
	resp := &acaseSts.ObjectTypeResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) ObjTypeListRequest(objTypeCode, objTypeName string) (*acaseSts.ObjTypeListResponseType, *AcaseResponseError) {
	req := &acaseSts.ObjTypeListRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
		Code:objTypeCode,
		Name:objTypeName,
	}
	resp := &acaseSts.ObjTypeListResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) OrderDocRequest(actionName acaseSts.OrderDocActionName, taskId, docId, code int) (*acaseSts.OrderDocsResponseType, *AcaseResponseError) {
	req := &acaseSts.OrderDocsRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
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
	resp := &acaseSts.OrderDocsResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) OrderInfoNotifyRequest(item *acaseSts.OrderInfoNotifyRequestType) (*acaseSts.OrderInfoNotifyResponseType, *AcaseResponseError) {
	item.Password = a.Password
	resp := &acaseSts.OrderInfoNotifyResponseType{}
	err := a.processRequest(item, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) OrderInfoAwocNotifyRequest(item *acaseSts.OrderInfoAwocNotifyRequestType) (*acaseSts.OrderInfoNotifyResponseType, *AcaseResponseError) {
	resp := &acaseSts.OrderInfoNotifyResponseType{}
	err := a.processRequest(item, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) OrderListRequest(arrivalDateFrom, arrivalDateTo, departureDateFrom, departureDateTo, deadlineDateFrom,
	deadlineDateTo, registrationDateFrom, registrationDateTo, accommodationDateFrom, accommodationDateTo, hotelName,
	lastName string, hotel int) (*acaseSts.OrderListResponseType, *AcaseResponseError) {

	req := &acaseSts.OrderListRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
		ArrivalDateFrom:arrivalDateFrom,
		ArrivalDateTo:arrivalDateTo,
		DepartureDateFrom:departureDateFrom,
		DepartureDateTo:departureDateTo,
		DeadlineDateFrom:deadlineDateFrom,
		DeadlineDateTo:deadlineDateTo,
		RegistrationDateFrom:registrationDateFrom,
		RegistrationDateTo:registrationDateTo,
		AccommodationDateFrom:accommodationDateFrom,
		AccommodationDateTo:accommodationDateTo,
		HotelName:hotelName,
		LastName:lastName,
		Hotel:hotel,
	}
	resp := &acaseSts.OrderListResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) OrderRequest(item *acaseSts.OrderRequestType) (*acaseSts.OrderResponseType, *AcaseResponseError) {

	item.Language = a.Language
	item.Password = a.Password
	item.UserId = a.UserId
	item.BuyerId = a.BuyerId

	resp := &acaseSts.OrderResponseType{}
	err := a.processRequest(item, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) OrderAwocRequest(item *acaseSts.OrderAwocRequestType) (*acaseSts.OrderResponseType, *AcaseResponseError) {

	item.Language = a.Language
	item.Password = a.Password
	item.UserId = a.UserId
	item.BuyerId = a.BuyerId

	resp := &acaseSts.OrderResponseType{}
	err := a.processRequest(item, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) RateGroupRequest(items []string) (*acaseSts.RateGroupResponseType, *AcaseResponseError) {

	req := &acaseSts.RateGroupRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
	}

	if items != nil && len(items) > 0 {
		req.ActionList.Parameters = make([]acaseSts.RateGroupParameterType, len(items))
		for i, item := range items {
			req.ActionList.Parameters[i].Name = item
		}
	} else {
		req.ActionList.Parameters = make([]acaseSts.RateGroupParameterType, 0)
	}
	resp := &acaseSts.RateGroupResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) RouteRequest(fromName, toName string, fromCode, toCode, fromTypeCode, toTypeCode int) (*acaseSts.RouteResponseType, *AcaseResponseError) {

	req := &acaseSts.RouteRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
	}
	req.Action.Name = "LIST"
	if fromName != "" {
		if toCode > 0 && toTypeCode > 0 {
			req.Action.Parameters.EndPoint.Code = toCode
			req.Action.Parameters.EndPoint.Type = toTypeCode
		}
		req.Action.Parameters.StartPoint.Name = fromName
	} else if toName != "" {
		if fromCode > 0 && fromTypeCode > 0 {
			req.Action.Parameters.StartPoint.Code = toCode
			req.Action.Parameters.StartPoint.Type = toTypeCode
		}
		req.Action.Parameters.EndPoint.Name = toName
	}
	resp := &acaseSts.RouteResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) PenaltyReasonRequest() (*acaseSts.PenaltyReasonResponseType, *AcaseResponseError) {

	req := &acaseSts.PenaltyReasonRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
	}
	req.Action.Name = "LISTPENALTY"
	resp := &acaseSts.PenaltyReasonResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) SpecialOfferTypeRequest(code int, name string) (*acaseSts.SpecialOfferTypeResponseType, *AcaseResponseError) {

	req := &acaseSts.SpecialOfferTypeRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
	}
	if code > 0 {
		req.ActionList.Parameters.Code = code
	}
	if name != "" {
		req.ActionList.Parameters.Name = name
	}
	resp := &acaseSts.SpecialOfferTypeResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) StarListRequest(code int, name, options string) (*acaseSts.StarListResponseType, *AcaseResponseError) {

	req := &acaseSts.StarListRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
		SimpleCodeNameType:acaseSts.SimpleCodeNameType{
			Code:code,
			Name:name,
		},
		Options:options,
	}
	resp := &acaseSts.StarListResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) StatusListRequest() (*acaseSts.StatusListResponseType, *AcaseResponseError) {

	req := &acaseSts.StatusListRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
	}
	resp := &acaseSts.StatusListResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) TypeOfPlaceRequest(typeOfPlaceCode int, typeOfPlaceName string) (*acaseSts.TypeOfPlaceResponseType, *AcaseResponseError) {

	req := &acaseSts.TypeOfPlaceRequestType{
		Credentials:acaseSts.Credentials{
			Language: a.Language,
			Password: a.Password,
			UserId: a.UserId,
			BuyerId: a.BuyerId,
		},
	}
	req.Action.Name = "LIST"
	req.Action.Parameters.TypeOfPlaceCode = typeOfPlaceCode
	req.Action.Parameters.TypeOfPlaceName = typeOfPlaceName
	resp := &acaseSts.TypeOfPlaceResponseType{}
	err := a.processRequest(req, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}
