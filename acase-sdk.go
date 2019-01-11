package acaseSdk

import (
	"bytes"
	"context"
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
	EventListener

	ApiUrl string

	credentials acaseSts.Credentials
}

func NewApi(auth Auth, apiUrl string) *Api {
	var lang acaseSts.LanguageTypeEnum

	if strings.ToLower(auth.Language) == "ru" {
		lang = acaseSts.Ru
	} else {
		lang = acaseSts.En
	}

	return (&Api{
		ApiUrl: apiUrl,

		credentials: acaseSts.Credentials{
			BuyerId:  auth.BuyerId,
			UserId:   auth.UserId,
			Password: auth.Password,
			Language: lang,
		},
	}).init()
}

func (a *Api) init() *Api {
	a.EventListener.Init()

	return a
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

func (a *Api) requestInternal(ctx context.Context, requestName, body string) ([]byte, error) {
	reqData := []byte(body)
	reqContentType := "application/x-www-form-urlencoded"

	req, err := http.NewRequest("POST", a.ApiUrl, bytes.NewBuffer(reqData))
	if err != nil {
		log.Print(err)
		return nil, err
	}

	req.Header.Add("Content-Type", reqContentType)
	req.Header.Add("Content-Length", string(len(body)))

	a.EventListener.raiseEvent(BeforeRequestSend, ctx, requestName, reqContentType, reqData)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	if err == nil {
		respContentType := resp.Header.Get("Content-Type")
		if len(respContentType) == 0 {
			respContentType = "text/xml"
		} else {
			parts := strings.Split(respContentType, "; ")
			respContentType = strings.TrimSpace(parts[0])
		}

		a.EventListener.raiseEvent(AfterResponseReceive, ctx, requestName, respContentType, respData)
	}

	return respData, nil
}

func (a *Api) formValuesRequest(ctx context.Context, requestName string, items map[string]string, res interface{}) *AcaseResponseError {
	v := url.Values{}
	v.Set("RequestName", requestName)
	v.Add("CompanyId", a.credentials.BuyerId)
	v.Add("UserId", a.credentials.UserId)
	v.Add("Password", a.credentials.Password)
	v.Add("Language", string(a.credentials.Language))

	for key, value := range items {
		v.Add(key, value)
	}

	//vl := v.Encode()
	//log.Print(vl)
	respData, err := a.requestInternal(ctx, requestName, v.Encode())
	if err != nil {
		log.Print(err)
		return &AcaseResponseError{
			Code:    "Unknown",
			Message: err.Error(),
		}
	}
	//log.Print(string(*respData))
	err = xml.Unmarshal(respData, res)
	if err != nil {
		log.Print(err)
		return &AcaseResponseError{
			Code:    "Unknown",
			Message: err.Error(),
		}
	}
	return nil
}

func (a *Api) processRequest(ctx context.Context, item interface{}, res interface{}) *AcaseResponseError {
	requestName, er := getTypeName(item)
	if er != nil {
		log.Print(er)
		return &AcaseResponseError{
			Code:    "Unknown",
			Message: er.Error(),
		}
	}

	bItem, err := xml.Marshal(item)
	if err != nil {
		log.Print(err)
		return &AcaseResponseError{
			Code:    "Unknown",
			Message: err.Error(),
		}
	}

	body := map[string]string{
		"XML": string(bItem),
	}

	return a.formValuesRequest(ctx, requestName, body, res)
}

func (a *Api) copyCredentials() acaseSts.Credentials {
	return a.credentials
}

func (a *Api) AdmUnit1Request(ctx context.Context, countryCode int, admUnitCode, admUnitName string) (*acaseSts.AdmUnit1ListType, *AcaseResponseError) {
	request := &acaseSts.AdmUnit1RequestType{
		Credentials: a.copyCredentials(),

		Action: acaseSts.AdmUnit1ActionType{
			Name: acaseSts.List,
			Parameters: acaseSts.AdmUnit1ActionTypeParameters{
				CountryCode:  countryCode,
				AdmUnit1Code: admUnitCode,
				AdmUnit1Name: admUnitCode,
			},
		},
	}

	resp := &acaseSts.AdmUnit1ResponseType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return &resp.AdmUnit1List, nil
}

func (a *Api) AdmUnit2Request(ctx context.Context, countryCode, admUnit1Code, admUnit2Code int, admUnit1Name, admUnit2Name string) (*acaseSts.AdmUnit2ListType, *AcaseResponseError) {
	request := &acaseSts.AdmUnit2RequestType{
		Credentials: a.copyCredentials(),

		Action: acaseSts.AdmUnit2ActionType{
			Name: acaseSts.List,
			Parameters: acaseSts.AdmUnit2ActionTypeParameters{
				CountryCode:  countryCode,
				AdmUnit1Code: admUnit1Code,
				AdmUnit1Name: admUnit1Name,
				AdmUnit2Code: admUnit2Code,
				AdmUnit2Name: admUnit2Name,
			},
		},
	}

	resp := &acaseSts.AdmUnit2ResponseType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return &resp.AdmUnit2List, nil
}

func (a *Api) CitizenshipListRequest(ctx context.Context) (*acaseSts.CitizenshipListType, *AcaseResponseError) {
	request := &acaseSts.CitizenshipListRequestType{
		Credentials: a.copyCredentials(),
	}

	resp := &acaseSts.CitizenshipListType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) CityDescriptionRequest(ctx context.Context, cityCode int64) (*acaseSts.CityDescriptionType, *AcaseResponseError) {
	request := map[string]string{
		"CityCode": strconv.FormatInt(cityCode, 10),
	}

	resp := &acaseSts.CityDescriptionType{}

	if err := a.formValuesRequest(ctx, "CityDescriptionRequest", request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) CityListRequest(ctx context.Context, countryName, cityName string, countryCode, cityCode int64) (*acaseSts.CityListType, *AcaseResponseError) {
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

	if err := a.formValuesRequest(ctx, "CityListRequest", request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) CountryDescriptionRequest(ctx context.Context, countryCode int64) (*acaseSts.CountryDescriptionType, *AcaseResponseError) {
	request := map[string]string{
		"CountryCode": strconv.FormatInt(countryCode, 10),
	}

	resp := &acaseSts.CountryDescriptionType{}

	if err := a.formValuesRequest(ctx, "CountryDescriptionRequest", request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) ClientCategoryListRequest(ctx context.Context, categoryCode int, categoryName string) (*acaseSts.ClientCategoryListType, *AcaseResponseError) {
	request := &acaseSts.ClientCategoryListRequestType{
		Credentials: a.copyCredentials(),

		ClientCategoryCode: categoryCode,
		ClientCategoryName: categoryName,
	}

	resp := &acaseSts.ClientCategoryListType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) CountryListRequest(ctx context.Context, countryCode int64, countryName string) (*acaseSts.CountryListType, *AcaseResponseError) {
	request := make(map[string]string)
	if countryCode > 0 {
		request["CountryCode"] = strconv.FormatInt(countryCode, 10)
	}
	if countryName != "" {
		request["CountryName"] = countryName
	}

	resp := &acaseSts.CountryListType{}

	if err := a.formValuesRequest(ctx, "CountryListRequest", request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) CurrencyListRequest(ctx context.Context, currencyCode int, currencyName, options string) (*acaseSts.CurrencyListResponseType, *AcaseResponseError) {
	request := &acaseSts.CurrencyListRequestType{
		Credentials: a.copyCredentials(),

		Code: currencyCode,
		Name: currencyName,
	}

	resp := &acaseSts.CurrencyListResponseType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) CustomerRequestCreate(ctx context.Context, fullName, zipCode, address, piAddress, inn, kpp, phone, name, buyerTypeName,
	countryName, cityName string, buyerTypeCode, countryCode, cityCode int) (*acaseSts.CustomerResponseCreateType, *AcaseResponseError) {

	request := &acaseSts.CustomerRequestCreateType{
		Credentials: a.copyCredentials(),

		ActionCreate: acaseSts.ActionCreateType{
			Parameters: acaseSts.CustomerRequestParametersType{
				Customer: &acaseSts.CustomerType{
					FullName:  fullName,
					ZipCode:   zipCode,
					Address:   address,
					PIAddress: piAddress,
					INN:       inn,
					KPP:       kpp,
					Phone:     phone,
					Name:      name,
					BuyerType: acaseSts.SimpleCodeNameType{
						Name: buyerTypeName,
						Code: buyerTypeCode,
					},
					Country: acaseSts.CountryType{
						Code: countryCode,
						Name: countryName,
					},
					City: acaseSts.CityType{
						Code: cityCode,
						Name: cityName,
					},
				},
			},
		},
	}

	resp := &acaseSts.CustomerResponseCreateType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) CustomerRequestUpdate(ctx context.Context, fullName, zipCode, address, piAddress, inn, kpp, phone, name, buyerTypeName,
	countryName, cityName string, customerCode, buyerTypeCode, countryCode, cityCode int) (*acaseSts.CustomerResponseUpdateType, *AcaseResponseError) {

	request := &acaseSts.CustomerRequestUpdateType{
		Credentials: a.copyCredentials(),

		ActionUpdate: acaseSts.ActionUpdateType{
			Parameters: acaseSts.CustomerRequestParametersType{
				Customer: &acaseSts.CustomerType{
					Code:      customerCode,
					FullName:  fullName,
					ZipCode:   zipCode,
					Address:   address,
					PIAddress: piAddress,
					INN:       inn,
					KPP:       kpp,
					Phone:     phone,
					Name:      name,
					BuyerType: acaseSts.SimpleCodeNameType{
						Name: buyerTypeName,
						Code: buyerTypeCode,
					},
					Country: acaseSts.CountryType{
						Code: countryCode,
						Name: countryName,
					},
					City: acaseSts.CityType{
						Code: cityCode,
						Name: cityName,
					},
				},
			},
		},
	}

	resp := &acaseSts.CustomerResponseUpdateType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) CustomerRequestDelete(ctx context.Context, customerCode int) (*acaseSts.CustomerResponseDeleteType, *AcaseResponseError) {
	request := &acaseSts.CustomerRequestDeleteType{
		Credentials: a.copyCredentials(),

		ActionDelete: acaseSts.ActionDeleteType{
			Parameters: acaseSts.CustomerRequestParametersType{
				Customer: &acaseSts.CustomerType{
					Code: customerCode,
				},
			},
		},
	}

	resp := &acaseSts.CustomerResponseDeleteType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) CustomerRequestList(ctx context.Context, sort, actualOnly int) (*acaseSts.CustomerResponseListType, *AcaseResponseError) {
	request := &acaseSts.CustomerRequestListType{
		Credentials: a.copyCredentials(),

		ActionList: acaseSts.ActionListType{
			Parameters: acaseSts.CustomerRequestParametersType{
				Sort:           sort,
				ShowActualOnly: actualOnly,
			},
		},
	}

	resp := &acaseSts.CustomerResponseListType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) CustomerRequestInfo(ctx context.Context, customerCode int) (*acaseSts.CustomerResponseInfoType, *AcaseResponseError) {
	request := &acaseSts.CustomerRequestInfoType{
		Credentials: a.copyCredentials(),

		ActionInfo: acaseSts.ActionInfoType{
			Parameters: acaseSts.CustomerRequestParametersType{
				CustomerCode: customerCode,
			},
		},
	}

	resp := &acaseSts.CustomerResponseInfoType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) HotelAmenityListRequest(ctx context.Context, hotelAmenityCode int64, hotelAmenityName string) (*acaseSts.HotelAmenityListResponseType, *AcaseResponseError) {
	request := make(map[string]string)
	if hotelAmenityCode > 0 {
		request["HotelAmenityCode"] = strconv.FormatInt(hotelAmenityCode, 10)
	}
	if hotelAmenityName != "" {
		request["HotelAmenityName"] = hotelAmenityName
	}

	resp := &acaseSts.HotelAmenityListResponseType{}

	if err := a.formValuesRequest(ctx, "HotelAmenityListRequest", request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) HotelDescriptionRequest(ctx context.Context, hotelCode, currencyCode int64) (*acaseSts.HotelDescriptionResponseType, *AcaseResponseError) {
	request := map[string]string{
		"HotelCode": strconv.FormatInt(hotelCode, 10),
	}

	if hotelCode <= 0 {
		return nil, &AcaseResponseError{Code: "0", Message: "Hotel code must be set"}
	}
	if currencyCode > 0 {
		request["CurrencyCode"] = strconv.FormatInt(currencyCode, 10)
	}

	resp := &acaseSts.HotelDescriptionResponseType{}

	if err := a.formValuesRequest(ctx, "HotelDescriptionRequest", request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) HotelListRequest(ctx context.Context, hotelCode, countryCode, cityCode, hotelRatingCode int64, hotelName, options string) (*acaseSts.HotelListResponseType, *AcaseResponseError) {
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

	if err := a.formValuesRequest(ctx, "HotelListRequest", request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) HotelPricingRequest2(ctx context.Context, productCode, currency, whereToPay, numberOfGuests, meal, numberOfExtraBedsAdult,
	numberOfExtraBedsChild, numberOfExtraBedsInfant, hotel int,
	arrivalDate, departureDate, arrivalTime, departureTime, id, accommodationId string) (*acaseSts.HotelPricingResponse2Type, *AcaseResponseError) {

	request := &acaseSts.HotelPricingRequest2Type{
		Credentials: a.copyCredentials(),

		Hotel:                   hotel,
		ProductCode:             productCode,
		Currency:                currency,
		WhereToPay:              whereToPay,
		ArrivalDate:             arrivalDate,
		DepartureDate:           departureDate,
		ArrivalTime:             arrivalTime,
		DepartureTime:           departureTime,
		NumberOfGuests:          numberOfGuests,
		NumberOfExtraBedsAdult:  numberOfExtraBedsAdult,
		NumberOfExtraBedsChild:  numberOfExtraBedsChild,
		NumberOfExtraBedsInfant: numberOfExtraBedsInfant,
		Meal:                    meal,
		Id:                      id,
		AccommodationId:         accommodationId,
	}

	resp := &acaseSts.HotelPricingResponse2Type{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) HotelProductRequest(ctx context.Context, currency, whereToPay, numberOfGuests, numberOfExtraBedsAdult,
	numberOfExtraBedsChild, numberOfExtraBedsInfant, hotel int,
	arrivalDate, departureDate, id, accommodationId string) (*acaseSts.HotelProductResponseType, *AcaseResponseError) {

	request := &acaseSts.HotelProductRequestType{
		Credentials: a.copyCredentials(),

		Hotel:                   hotel,
		Currency:                currency,
		WhereToPay:              whereToPay,
		ArrivalDate:             arrivalDate,
		DepartureDate:           departureDate,
		Id:                      id,
		AccommodationId:         accommodationId,
		NumberOfGuests:          numberOfGuests,
		NumberOfExtraBedsAdult:  numberOfExtraBedsAdult,
		NumberOfExtraBedsChild:  numberOfExtraBedsChild,
		NumberOfExtraBedsInfant: numberOfExtraBedsInfant,
	}

	resp := &acaseSts.HotelProductResponseType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) HotelSearchRequest(ctx context.Context, arrivalDate, departureDate, options, hotelName, destListCode string,
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
		Credentials: a.copyCredentials(),

		ArrivalDate:    arrivalDate,
		DepartureDate:  departureDate,
		City:           city,
		PriceFrom:      prFrom,
		PriceTo:        prTo,
		FreeSaleOnly:   freeSaleOnly,
		HotelCategory:  hotelCategory,
		Currency:       currency,
		WhereToPay:     whereToPay,
		NumberOfGuests: numberOfGuests,
		Options:        options,
		HotelCode:      hotelCode,
		HotelName:      hotelName,
	}

	if starCodes != nil && len(starCodes) > 0 {
		for _, starCode := range starCodes {
			request.StarList.Items = append(request.StarList.Items, *(&acaseSts.SimpleCodeType{Code: starCode}))
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
		request.Guests = &acaseSts.GuestsType{NumberOfAdults: guestsAdults}
		for _, minorAge := range minorAges {
			request.Guests.MinorAgeList.Items = append(request.Guests.MinorAgeList.Items, *(&acaseSts.MinorType{Age: minorAge}))
		}
	}

	resp := &acaseSts.HotelSearchResponseType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) MealRequest(ctx context.Context, mealCode, mealTypeCode int64, mealName string) (*acaseSts.MealResponseType, *AcaseResponseError) {
	request := &acaseSts.MealRequestType{
		Credentials: a.copyCredentials(),

		Action: acaseSts.ActionType{
			Name: "LIST",
			Parameters: acaseSts.ParametersType{
				MealName: mealName,
			},
		},
	}

	if mealTypeCode > 0 {
		request.Action.Parameters.MealTypeCode = strconv.FormatInt(mealTypeCode, 10)
	}
	if mealCode > 0 {
		request.Action.Parameters.MealCode = strconv.FormatInt(mealCode, 10)
	}

	resp := &acaseSts.MealResponseType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) MealTypeRequest(ctx context.Context, mealTypeCode int64, mealName string) (*acaseSts.MealTypeResponseType, *AcaseResponseError) {
	request := &acaseSts.MealTypeRequestType{
		Credentials: a.copyCredentials(),

		Action: acaseSts.MealTypeActionType{
			Name: "LIST",
			Parameters: acaseSts.MealTypeParametersType{
				MealTypeName: mealName,
			},
		},
	}

	if mealTypeCode > 0 {
		request.Action.Parameters.MealTypeCode = strconv.FormatInt(mealTypeCode, 10)
	}

	resp := &acaseSts.MealTypeResponseType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) ObjectRequest(ctx context.Context, objectTypeCode, objectSubTypeCode, cityCode int64) (*acaseSts.ObjectResponseType, *AcaseResponseError) {
	request := &acaseSts.ObjectRequestType{
		Credentials: a.copyCredentials(),

		Action: acaseSts.ObjectActionType{
			Name: "LISTMETROSTYLE",
		},
	}

	if objectTypeCode > 0 {
		request.Action.Parameters.ObjectTypeCode = strconv.FormatInt(objectTypeCode, 10)
	}
	if objectSubTypeCode > 0 {
		request.Action.Parameters.ObjectSubTypeCode = strconv.FormatInt(objectSubTypeCode, 10)
	}
	if cityCode > 0 {
		request.Action.Parameters.CityCode = strconv.FormatInt(cityCode, 10)
	}

	resp := &acaseSts.ObjectResponseType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) ObjectSubTypeRequest(ctx context.Context, objectTypeCode int64) (*acaseSts.ObjectSubTypeResponseType, *AcaseResponseError) {
	request := &acaseSts.ObjectSubTypeRequestType{
		Credentials: a.copyCredentials(),

		Action: acaseSts.ObjSubTypeActionType{
			Name: "LIST",
		},
	}

	if objectTypeCode > 0 {
		request.Action.Parameters.ObjectTypeCode = strconv.FormatInt(objectTypeCode, 10)
	}

	resp := &acaseSts.ObjectSubTypeResponseType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) ObjectTypeRequest(ctx context.Context, objectTypeCode int64) (*acaseSts.GeoObjectTypeResponseType, *AcaseResponseError) {
	request := &acaseSts.ObjectTypeRequestType{
		Credentials: a.copyCredentials(),

		Action: acaseSts.ObjectTypeActionType{
			Name: "LIST",
		},
	}

	if objectTypeCode > 0 {
		request.Action.Parameters.ObjectTypeCode = strconv.FormatInt(objectTypeCode, 10)
	}

	resp := &acaseSts.GeoObjectTypeResponseType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) ObjTypeListRequest(ctx context.Context, objTypeCode, objTypeName string) (*acaseSts.ObjTypeListResponseType, *AcaseResponseError) {
	request := &acaseSts.ObjTypeListRequestType{
		Credentials: a.copyCredentials(),

		Code: objTypeCode,
		Name: objTypeName,
	}

	resp := &acaseSts.ObjTypeListResponseType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) OrderDocRequest(ctx context.Context, actionName acaseSts.OrderDocActionName, taskId, docId, code int) (*acaseSts.OrderDocsResponseType, *AcaseResponseError) {
	request := &acaseSts.OrderDocsRequestType{
		Credentials: a.copyCredentials(),

		Action: acaseSts.OrderDocActionType{
			Name: string(actionName),
		},
	}

	switch actionName {
	case acaseSts.TASKADD:
		request.Action.Parameters.DocId = docId
		request.Action.Parameters.DocType = &acaseSts.OrderDocParamsDocType{Code: code}
	case acaseSts.TASKSTATUS, acaseSts.TASKRESPONSE:
		request.Action.Parameters.TaskId = taskId
	}

	resp := &acaseSts.OrderDocsResponseType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) OrderInfoRequest(ctx context.Context, item *acaseSts.OrderInfoRequestType) (*acaseSts.OrderInfoResponseType, *AcaseResponseError) {
	item.Credentials = a.copyCredentials()

	resp := &acaseSts.OrderInfoResponseType{}

	err := a.processRequest(ctx, item, resp)
	if err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) OrderInfoNotifyRequest(ctx context.Context, item *acaseSts.OrderInfoNotifyRequestType) (*acaseSts.OrderInfoNotifyResponseType, *AcaseResponseError) {
	item.Password = a.credentials.Password

	resp := &acaseSts.OrderInfoNotifyResponseType{}

	if err := a.processRequest(ctx, item, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) OrderInfoAwocNotifyRequest(ctx context.Context, item *acaseSts.OrderInfoAwocNotifyRequestType) (*acaseSts.OrderInfoNotifyResponseType, *AcaseResponseError) {
	resp := &acaseSts.OrderInfoNotifyResponseType{}

	if err := a.processRequest(ctx, item, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) OrderListRequest(ctx context.Context, arrivalDateFrom, arrivalDateTo, departureDateFrom, departureDateTo, deadlineDateFrom,
	deadlineDateTo, registrationDateFrom, registrationDateTo, accommodationDateFrom, accommodationDateTo, hotelName,
	lastName string, hotel int) (*acaseSts.OrderListResponseType, *AcaseResponseError) {

	request := &acaseSts.OrderListRequestType{
		Credentials: a.copyCredentials(),

		ArrivalDateFrom:       arrivalDateFrom,
		ArrivalDateTo:         arrivalDateTo,
		DepartureDateFrom:     departureDateFrom,
		DepartureDateTo:       departureDateTo,
		DeadlineDateFrom:      deadlineDateFrom,
		DeadlineDateTo:        deadlineDateTo,
		RegistrationDateFrom:  registrationDateFrom,
		RegistrationDateTo:    registrationDateTo,
		AccommodationDateFrom: accommodationDateFrom,
		AccommodationDateTo:   accommodationDateTo,
		HotelName:             hotelName,
		LastName:              lastName,
		Hotel:                 hotel,
	}

	resp := &acaseSts.OrderListResponseType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) OrderRequest(ctx context.Context, item *acaseSts.OrderRequestType) (*acaseSts.OrderResponseType, *AcaseResponseError) {
	item.Credentials = a.copyCredentials()

	resp := &acaseSts.OrderResponseType{}

	if err := a.processRequest(ctx, item, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) OrderAwocRequest(ctx context.Context, item *acaseSts.OrderAwocRequestType) (*acaseSts.OrderResponseType, *AcaseResponseError) {
	item.Credentials = a.copyCredentials()

	resp := &acaseSts.OrderResponseType{}

	if err := a.processRequest(ctx, item, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) RateGroupRequest(ctx context.Context, items []string) (*acaseSts.RateGroupResponseType, *AcaseResponseError) {
	request := &acaseSts.RateGroupRequestType{
		Credentials: a.copyCredentials(),
	}

	request.ActionList.Parameters = make([]acaseSts.RateGroupParameterType, len(items))

	for i, item := range items {
		request.ActionList.Parameters[i].Name = item
	}

	resp := &acaseSts.RateGroupResponseType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) RouteRequest(ctx context.Context, fromName, toName string, fromCode, toCode, fromTypeCode, toTypeCode int) (*acaseSts.RouteResponseType, *AcaseResponseError) {
	request := &acaseSts.RouteRequestType{
		Credentials: a.copyCredentials(),

		Action: acaseSts.ActionRouteType{
			Name: "LIST",
		},
	}

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

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) PenaltyReasonRequest(ctx context.Context) (*acaseSts.PenaltyReasonResponseType, *AcaseResponseError) {
	request := &acaseSts.PenaltyReasonRequestType{
		Credentials: a.copyCredentials(),

		Action: acaseSts.PenaltyReasonActionType{
			Name: "LISTPENALTY",
		},
	}

	resp := &acaseSts.PenaltyReasonResponseType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) SpecialOfferTypeRequest(ctx context.Context, code int, name string) (*acaseSts.SpecialOfferTypeResponseType, *AcaseResponseError) {
	request := &acaseSts.SpecialOfferTypeRequestType{
		Credentials: a.copyCredentials(),
	}

	if code > 0 {
		request.ActionList.Parameters.Code = code
	}
	if name != "" {
		request.ActionList.Parameters.Name = name
	}

	resp := &acaseSts.SpecialOfferTypeResponseType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) StarListRequest(ctx context.Context, code int64, name, options string) (*acaseSts.StarListResponseType, *AcaseResponseError) {
	request := &acaseSts.StarListRequestType{
		Credentials: a.copyCredentials(),

		Name:    name,
		Options: options,
	}
	if code > 0 {
		request.Code = strconv.FormatInt(code, 10)
	}

	resp := &acaseSts.StarListResponseType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) StatusListRequest(ctx context.Context) (*acaseSts.StatusListResponseType, *AcaseResponseError) {
	request := &acaseSts.StatusListRequestType{
		Credentials: a.copyCredentials(),
	}

	resp := &acaseSts.StatusListResponseType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}

func (a *Api) TypeOfPlaceRequest(ctx context.Context, typeOfPlaceCode int64, typeOfPlaceName string) (*acaseSts.TypeOfPlaceResponseType, *AcaseResponseError) {
	request := &acaseSts.TypeOfPlaceRequestType{
		Credentials: a.copyCredentials(),

		Action: acaseSts.TypeOfPlaceRequestActionType{
			Name: "LIST",
		},
	}

	if typeOfPlaceCode > 0 {
		request.Action.Parameters.TypeOfPlaceCode = strconv.FormatInt(typeOfPlaceCode, 10)
	}

	request.Action.Parameters.TypeOfPlaceName = typeOfPlaceName

	resp := &acaseSts.TypeOfPlaceResponseType{}

	if err := a.processRequest(ctx, request, resp); err != nil {
		return nil, err
	}

	if acsErr := ResponseError(resp.BaseResponse); acsErr != nil {
		return nil, acsErr
	}

	return resp, nil
}
