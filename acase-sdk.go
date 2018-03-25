package acaseSdk

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/tmconsulting/acase-golang-sdk/acaseSts"
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
	if auth.Language == "RU" || auth.Language == "ru" {
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

func getTypeName(item interface{}) (string, error) {
	t := reflect.TypeOf(item)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	f, ext := t.FieldByName("XMLName")
	if !ext {
		return "", errors.New("XMLName property not found")
	}
	vl, ok := f.Tag.Lookup("xml")
	if !ok || len(vl) == 0 {
		return "", errors.New("XML tag not found")
	}
	vl = strings.Split(vl, ",")[0]
	return vl, nil
}

func (a *Api) requestInternal(body string) (*[]byte, error) {
	req, err := http.NewRequest("POST", a.ApiUrl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Print(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", string(len(body)))
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

func (a * Api) formValuesRequest(requestName string, items map[string]string, res interface{}) *AcaseResponseError {
	v := url.Values{}
	v.Set("RequestName", requestName)
	v.Add("CompanyId", a.BuyerId)
	v.Add("UserId", a.UserId)
	v.Add("Password", a.Password)
	v.Add("Language", string(a.Language))
	for key, value := range items {
		v.Add(key, value)
	}
	//vl := v.Encode()
	//log.Print(vl)
	respData, err := a.requestInternal(v.Encode())
	if err != nil {
		log.Print(err)
		return &AcaseResponseError{
			Code:"Unknown",
			Message:err.Error(),
		}
	}
	//log.Print(string(*respData))
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

func (a *Api) processRequest(item interface{}, res interface{}) *AcaseResponseError {
	requestName, er := getTypeName(item)
	if er != nil {
		log.Print(er)
		return &AcaseResponseError{
			Code:"Unknown",
			Message:er.Error(),
		}
	}
	bItem, err := xml.Marshal(item)
	if err != nil {
		log.Print(err)
		return &AcaseResponseError{
			Code:"Unknown",
			Message:err.Error(),
		}
	}
	body := make(map[string]string, 1)
	body["XML"] = string(bItem)
	return a.formValuesRequest(requestName, body, res)
}

func (a *Api) fillCredentials(item *acaseSts.Credentials) {
	item.BuyerId = a.BuyerId
	item.UserId = a.UserId
	item.Password = a.Password
	item.Language = a.Language
}

func (a *Api) AdmUnit1Request(countryCode int, admUnitCode, admUnitName string) (*acaseSts.AdmUnit1ListType, *AcaseResponseError) {
	request := &acaseSts.AdmUnit1RequestType{
		Action: acaseSts.AdmUnit1ActionType{
			Name: acaseSts.List,
			Parameters: acaseSts.AdmUnit1ActionTypeParameters{
				CountryCode: countryCode,
				AdmUnit1Code: admUnitCode,
				AdmUnit1Name: admUnitCode,
			},
		},
	}
	a.fillCredentials(&request.Credentials)
	resp := &acaseSts.AdmUnit1ResponseType{}
	err := a.processRequest(request, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}
	return &resp.AdmUnit1List, nil
}

func (a *Api) AdmUnit2Request(countryCode, admUnit1Code, admUnit2Code int, admUnit1Name, admUnit2Name string) (*acaseSts.AdmUnit2ListType, *AcaseResponseError) {
	request := &acaseSts.AdmUnit2RequestType{
		Action: acaseSts.AdmUnit2ActionType{
			Name: acaseSts.List,
			Parameters: acaseSts.AdmUnit2ActionTypeParameters{
				CountryCode: countryCode,
				AdmUnit1Code: admUnit1Code,
				AdmUnit1Name: admUnit1Name,
				AdmUnit2Code: admUnit2Code,
				AdmUnit2Name: admUnit2Name,
			},
		},
	}
	a.fillCredentials(&request.Credentials)
	resp := &acaseSts.AdmUnit2ResponseType{}
	err := a.processRequest(request, resp)
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
	request := &acaseSts.CitizenshipListRequestType{}
	a.fillCredentials(&request.Credentials)
	resp := &acaseSts.CitizenshipListType{}
	err := a.processRequest(request, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}
	return resp, nil
}

func (a *Api) CityDescriptionRequest(cityCode int64) (*acaseSts.CityDescriptionType, *AcaseResponseError) {
	request := make(map[string]string, 1)
	request["CityCode"] = strconv.FormatInt(cityCode, 10)
	resp := &acaseSts.CityDescriptionType{}
	err := a.formValuesRequest("CityDescriptionRequest", request, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}
	return resp, nil
}

func (a *Api) CityListRequest(countryName, cityName string, countryCode, cityCode int64) (*acaseSts.CityListType, *AcaseResponseError) {
	request := make(map[string]string)
	if countryName != "" {
		request["CountryName"] = countryName
	}
	if cityName != "" {
		request["CityName"] = cityName
	}
	if countryCode > 0 {
		request["CountryCode"] = strconv.FormatInt(countryCode, 10)
	}
	if cityCode > 0 {
		request["CityCode"] = strconv.FormatInt(cityCode, 10)
	}
	resp := &acaseSts.CityListType{}
	err := a.formValuesRequest("CityListRequest", request, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}
	return resp, nil
}

func (a *Api) CountryDescriptionRequest(countryCode int64) (*acaseSts.CountryDescriptionType, *AcaseResponseError) {
	request := make(map[string]string, 1)
	request["CountryCode"] = strconv.FormatInt(countryCode, 10)
	resp := &acaseSts.CountryDescriptionType{}
	err := a.formValuesRequest("CountryDescriptionRequest", request, resp)
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
	request := &acaseSts.ClientCategoryListRequestType{
		ClientCategoryCode: categoryCode,
		ClientCategoryName: categoryName,
	}
	a.fillCredentials(&request.Credentials)
	resp := &acaseSts.ClientCategoryListType{}
	err := a.processRequest(request, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}
	return resp, nil
}

func (a *Api) CountryListRequest(countryCode int64, countryName string) (*acaseSts.CountryListType, *AcaseResponseError) {
	request := make(map[string]string)
	if countryCode > 0 {
		request["CountryCode"] = strconv.FormatInt(countryCode, 10)
	}
	if countryName != "" {
		request["CountryName"] = countryName
	}
	resp := &acaseSts.CountryListType{}
	err := a.formValuesRequest("CountryListRequest", request, resp)
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
	request := &acaseSts.CurrencyListRequestType{
		Code: currencyCode,
		Name: currencyName,
	}
	a.fillCredentials(&request.Credentials)
	resp := &acaseSts.CurrencyListResponseType{}
	err := a.processRequest(request, resp)
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

	request := &acaseSts.CustomerRequestCreateType{
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
	a.fillCredentials(&request.Credentials)
	resp := &acaseSts.CustomerResponseCreateType{}
	err := a.processRequest(request, resp)
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

	request := &acaseSts.CustomerRequestUpdateType{
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
	a.fillCredentials(&request.Credentials)
	resp := &acaseSts.CustomerResponseUpdateType{}
	err := a.processRequest(request, resp)
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

	request := &acaseSts.CustomerRequestDeleteType{
		ActionDelete: acaseSts.ActionDeleteType{
			Parameters:acaseSts.CustomerRequestParametersType{
				Customer:&acaseSts.CustomerType{
					Code:customerCode,
				},
			},
		},
	}
	a.fillCredentials(&request.Credentials)
	resp := &acaseSts.CustomerResponseDeleteType{}
	err := a.processRequest(request, resp)
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
	request := &acaseSts.CustomerRequestListType{
		ActionList: acaseSts.ActionListType{
			Parameters:acaseSts.CustomerRequestParametersType{
				Sort:sort,
				ShowActualOnly:actualOnly,
			},
		},
	}
	a.fillCredentials(&request.Credentials)
	resp := &acaseSts.CustomerResponseListType{}
	err := a.processRequest(request, resp)
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
	request := &acaseSts.CustomerRequestInfoType{
		ActionInfo: acaseSts.ActionInfoType{
			Parameters:acaseSts.CustomerRequestParametersType{
				CustomerCode:customerCode,
			},
		},
	}
	a.fillCredentials(&request.Credentials)
	resp := &acaseSts.CustomerResponseInfoType{}
	err := a.processRequest(request, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}
	return resp, nil
}

func (a *Api) HotelAmenityListRequest(hotelAmenityCode int64, hotelAmenityName string) (*acaseSts.HotelAmenityListResponseType, *AcaseResponseError) {
	request := make(map[string]string)
	if hotelAmenityCode > 0 {
		request["HotelAmenityCode"] = strconv.FormatInt(hotelAmenityCode, 10)
	}
	if hotelAmenityName != "" {
		request["HotelAmenityName"] = hotelAmenityName
	}
	resp := &acaseSts.HotelAmenityListResponseType{}
	err := a.formValuesRequest("HotelAmenityListRequest", request, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}
	return resp, nil
}

func (a *Api) HotelDescriptionRequest(hotelCode, currencyCode int64) (*acaseSts.HotelDescriptionResponseType, *AcaseResponseError) {
	request := make(map[string]string)
	if hotelCode <= 0 {
		return nil, &AcaseResponseError{Code:"0", Message:"Hotel code must be set"}
	}
	request["HotelCode"] = strconv.FormatInt(hotelCode, 10)
	if currencyCode > 0 {
		request["CurrencyCode"] = strconv.FormatInt(currencyCode, 10)
	}
	resp := &acaseSts.HotelDescriptionResponseType{}
	err := a.formValuesRequest("HotelDescriptionRequest", request, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}
	return resp, nil
}

func (a *Api) HotelListRequest(hotelCode, countryCode, cityCode, hotelRatingCode int64, hotelName, options string) (*acaseSts.HotelListResponseType, *AcaseResponseError) {
	request := make(map[string]string)
	if hotelCode > 0 {
		request["HotelCode"] = strconv.FormatInt(hotelCode, 10)
	}
	if countryCode > 0 {
		request["CountryCode"] = strconv.FormatInt(countryCode, 10)
	}
	if cityCode > 0 {
		request["CityCode"] = strconv.FormatInt(cityCode, 10)
	}
	if hotelRatingCode > 0 {
		request["HotelRatingCode"] = strconv.FormatInt(hotelRatingCode, 10)
	}
	if hotelName != "" {
		request["HotelName"] = hotelName
	}
	if options != "" {
		request["Opt"] = options
	}
	resp := &acaseSts.HotelListResponseType{}
	err := a.formValuesRequest("HotelListRequest", request, resp)
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

	request := &acaseSts.HotelPricingRequest2Type{
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
	a.fillCredentials(&request.Credentials)
	resp := &acaseSts.HotelPricingResponse2Type{}
	err := a.processRequest(request, resp)
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

	request := &acaseSts.HotelProductRequestType{
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
	a.fillCredentials(&request.Credentials)
	resp := &acaseSts.HotelProductResponseType{}
	err := a.processRequest(request, resp)
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

	prFrom := ""
	prTo := ""
	if priceFrom > 0 {
		prFrom = strconv.FormatFloat(priceFrom, 'f', 2, 64)
	}
	if priceTo > 0 {
		prTo = strconv.FormatFloat(priceTo, 'f', 2, 64)
	}

	request := &acaseSts.HotelSearchRequestType{
		ArrivalDate:arrivalDate,
		DepartureDate:departureDate,
		City:city,
		PriceFrom:prFrom,
		PriceTo:prTo,
		FreeSaleOnly:freeSaleOnly,
		HotelCategory:hotelCategory,
		Currency:currency,
		WhereToPay:whereToPay,
		NumberOfGuests:numberOfGuests,
		Options:options,
		HotelCode:hotelCode,
		HotelName:hotelName,
	}
	a.fillCredentials(&request.Credentials)
	if starCodes != nil && len(starCodes) > 0 {
		for _, starCode := range starCodes {
			request.StarList.Items = append(request.StarList.Items, *(&acaseSts.SimpleCodeType{Code:starCode}))
		}
	}
	if destListCode != "" {
		request.DestinationList.Code = destListCode
		request.DestinationList.Distance = distance
		request.DestinationList.TypeCode = acaseSts.DestinationTypeEnum(distTypeCode)
	} else if distance > 0 {
		request.Destination.TypeCode = acaseSts.DestinationTypeEnum(distTypeCode)
		request.Destination.Distance = distance
		request.Destination.Code = distCode
	}
	if guestsAdults > 0 && minorAges != nil && len(minorAges) > 0 {
		request.Guests = &acaseSts.GuestsType{NumberOfAdults:guestsAdults}
		for _, minorAge := range minorAges {
			request.Guests.MinorAgeList.Items = append(request.Guests.MinorAgeList.Items, *(&acaseSts.MinorType{Age:minorAge}))
		}
	}
	resp := &acaseSts.HotelSearchResponseType{}
	err := a.processRequest(request, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}
	return resp, nil
}

func (a *Api) MealRequest(mealCode, mealTypeCode int64, mealName string) (*acaseSts.MealResponseType, *AcaseResponseError) {
	request := &acaseSts.MealRequestType{}
	a.fillCredentials(&request.Credentials)
	request.Action.Name = "LIST"
	request.Action.Parameters.MealName = mealName
	if mealTypeCode > 0 {
		request.Action.Parameters.MealTypeCode = strconv.FormatInt(mealTypeCode, 10)
	} else {
		request.Action.Parameters.MealTypeCode = ""
	}
	if mealCode > 0 {
		request.Action.Parameters.MealCode = strconv.FormatInt(mealCode, 10)
	} else {
		request.Action.Parameters.MealCode = ""
	}
	resp := &acaseSts.MealResponseType{}
	err := a.processRequest(request, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}
	return resp, nil
}

func (a *Api) MealTypeRequest(mealTypeCode int64, mealName string) (*acaseSts.MealTypeResponseType, *AcaseResponseError) {
	request := &acaseSts.MealTypeRequestType{}
	request.Action.Name = "LIST"
	if mealTypeCode > 0 {
		request.Action.Parameters.MealTypeCode = strconv.FormatInt(mealTypeCode, 10)
	} else {
		request.Action.Parameters.MealTypeCode = ""
	}
	request.Action.Parameters.MealTypeName = mealName
	a.fillCredentials(&request.Credentials)
	resp := &acaseSts.MealTypeResponseType{}
	err := a.processRequest(request, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}
	return resp, nil
}

func (a *Api) ObjectRequest(objectTypeCode, objectSubTypeCode, cityCode int64) (*acaseSts.ObjectResponseType, *AcaseResponseError) {
	request := &acaseSts.ObjectRequestType{}
	a.fillCredentials(&request.Credentials)
	request.Action.Name = "LISTMETROSTYLE"
	if objectTypeCode > 0 {
		request.Action.Parameters.ObjectTypeCode = strconv.FormatInt(objectTypeCode, 10)
	} else {
		request.Action.Parameters.ObjectTypeCode = ""
	}
	if objectSubTypeCode > 0 {
		request.Action.Parameters.ObjectSubTypeCode = strconv.FormatInt(objectSubTypeCode, 10)
	} else {
		request.Action.Parameters.ObjectSubTypeCode = ""
	}
	if cityCode > 0 {
		request.Action.Parameters.CityCode = strconv.FormatInt(cityCode, 10)
	} else {
		request.Action.Parameters.CityCode = ""
	}
	resp := &acaseSts.ObjectResponseType{}
	err := a.processRequest(request, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}
	return resp, nil
}

func (a *Api) ObjectSubTypeRequest(objectTypeCode int64) (*acaseSts.ObjectSubTypeResponseType, *AcaseResponseError) {
	request := &acaseSts.ObjectSubTypeRequestType{}
	a.fillCredentials(&request.Credentials)
	request.Action.Name = "LIST"
	if objectTypeCode > 0 {
		request.Action.Parameters.ObjectTypeCode = strconv.FormatInt(objectTypeCode, 10)
	} else {
		request.Action.Parameters.ObjectTypeCode = ""
	}
	resp := &acaseSts.ObjectSubTypeResponseType{}
	err := a.processRequest(request, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}
	return resp, nil
}

func (a *Api) ObjectTypeRequest(objectTypeCode int64) (*acaseSts.GeoObjectTypeResponseType, *AcaseResponseError) {
	request := &acaseSts.ObjectTypeRequestType{}
	a.fillCredentials(&request.Credentials)
	request.Action.Name = "LIST"
	if objectTypeCode > 0 {
		request.Action.Parameters.ObjectTypeCode = strconv.FormatInt(objectTypeCode, 10)
	} else {
		request.Action.Parameters.ObjectTypeCode = ""
	}
	resp := &acaseSts.GeoObjectTypeResponseType{}
	err := a.processRequest(request, resp)
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
	request := &acaseSts.ObjTypeListRequestType{
		Code:objTypeCode,
		Name:objTypeName,
	}
	a.fillCredentials(&request.Credentials)
	resp := &acaseSts.ObjTypeListResponseType{}
	err := a.processRequest(request, resp)
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
	request := &acaseSts.OrderDocsRequestType{}
	a.fillCredentials(&request.Credentials)
	request.Action.Name = string(actionName)

	switch actionName {
		case acaseSts.TASKADD:
			request.Action.Parameters.DocId = docId
			request.Action.Parameters.DocType = &acaseSts.OrderDocParamsDocType{Code:code}
		case acaseSts.TASKSTATUS, acaseSts.TASKRESPONSE:
			request.Action.Parameters.TaskId = taskId
	}
	resp := &acaseSts.OrderDocsResponseType{}
	err := a.processRequest(request, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}
	return resp, nil
}

func (a *Api) OrderInfoRequest(item *acaseSts.OrderInfoRequestType) (*acaseSts.OrderInfoResponseType, *AcaseResponseError) {
	a.fillCredentials(&item.Credentials)
	resp := &acaseSts.OrderInfoResponseType{}
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

	request := &acaseSts.OrderListRequestType{
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
	a.fillCredentials(&request.Credentials)
	resp := &acaseSts.OrderListResponseType{}
	err := a.processRequest(request, resp)
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

	a.fillCredentials(&item.Credentials)
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

	a.fillCredentials(&item.Credentials)
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

	request := &acaseSts.RateGroupRequestType{}
	a.fillCredentials(&request.Credentials)

	if items != nil && len(items) > 0 {
		request.ActionList.Parameters = make([]acaseSts.RateGroupParameterType, len(items))
		for i, item := range items {
			request.ActionList.Parameters[i].Name = item
		}
	} else {
		request.ActionList.Parameters = make([]acaseSts.RateGroupParameterType, 0)
	}
	resp := &acaseSts.RateGroupResponseType{}
	err := a.processRequest(request, resp)
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

	request := &acaseSts.RouteRequestType{}
	a.fillCredentials(&request.Credentials)
	request.Action.Name = "LIST"
	if fromName != "" {
		if toCode > 0 && toTypeCode > 0 {
			request.Action.Parameters.EndPoint.Code = toCode
			request.Action.Parameters.EndPoint.Type = toTypeCode
		}
		request.Action.Parameters.StartPoint.Name = fromName
	} else if toName != "" {
		if fromCode > 0 && fromTypeCode > 0 {
			request.Action.Parameters.StartPoint.Code = toCode
			request.Action.Parameters.StartPoint.Type = toTypeCode
		}
		request.Action.Parameters.EndPoint.Name = toName
	}
	resp := &acaseSts.RouteResponseType{}
	err := a.processRequest(request, resp)
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

	request := &acaseSts.PenaltyReasonRequestType{}
	a.fillCredentials(&request.Credentials)
	request.Action.Name = "LISTPENALTY"
	resp := &acaseSts.PenaltyReasonResponseType{}
	err := a.processRequest(request, resp)
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

	request := &acaseSts.SpecialOfferTypeRequestType{}
	a.fillCredentials(&request.Credentials)
	if code > 0 {
		request.ActionList.Parameters.Code = code
	}
	if name != "" {
		request.ActionList.Parameters.Name = name
	}
	resp := &acaseSts.SpecialOfferTypeResponseType{}
	err := a.processRequest(request, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}
	return resp, nil
}

func (a *Api) StarListRequest(code int64, name, options string) (*acaseSts.StarListResponseType, *AcaseResponseError) {

	request := &acaseSts.StarListRequestType{
		Name:name,
		Options:options,
	}
	if code > 0 {
		request.Code = strconv.FormatInt(code, 10)
	} else {
		request.Code = ""
	}
	a.fillCredentials(&request.Credentials)
	resp := &acaseSts.StarListResponseType{}
	err := a.processRequest(request, resp)
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

	request := &acaseSts.StatusListRequestType{}
	a.fillCredentials(&request.Credentials)
	resp := &acaseSts.StatusListResponseType{}
	err := a.processRequest(request, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}
	return resp, nil
}

func (a *Api) TypeOfPlaceRequest(typeOfPlaceCode int64, typeOfPlaceName string) (*acaseSts.TypeOfPlaceResponseType, *AcaseResponseError) {

	request := &acaseSts.TypeOfPlaceRequestType{}
	a.fillCredentials(&request.Credentials)
	request.Action.Name = "LIST"
	if typeOfPlaceCode > 0 {
		request.Action.Parameters.TypeOfPlaceCode = strconv.FormatInt(typeOfPlaceCode, 10)
	} else {
		request.Action.Parameters.TypeOfPlaceCode = ""
	}
	request.Action.Parameters.TypeOfPlaceName = typeOfPlaceName
	resp := &acaseSts.TypeOfPlaceResponseType{}
	err := a.processRequest(request, resp)
	if err != nil {
		return nil, err
	}
	acsErr := ResponseError(resp.BaseResponse)
	if acsErr != nil {
		return nil, acsErr
	}
	return resp, nil
}
