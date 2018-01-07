package test

import (
	"testing"
	"gopkg.in/h2non/gock.v1"
	"github.com/nbio/st"
	"github.com/tmconsulting/acase-golang-sdk/acaseSts"
)

func getOrderRequestObject() *acaseSts.OrderRequestType {
	res := &acaseSts.OrderRequestType{}
	res.BuyerId="MyCompanyId"
	res.UserId="MyUserId"
	res.Password="MyPassword"
	res.Language="ru"
	res.Id=0
	res.ReferenceNumber=""
	res.Currency.Code=3
	res.WhereToPay.Code=3
	res.AccommodationList.Accommodation = make([]acaseSts.AccommodationType, 1)
	res.AccommodationList.Accommodation[0] = *(&acaseSts.AccommodationType{})
	res.AccommodationList.Accommodation[0].Id=0
	res.AccommodationList.Accommodation[0].ArrivalDate="21.11.2004"
	res.AccommodationList.Accommodation[0].ArrivalTime="10:00"
	res.AccommodationList.Accommodation[0].DepartureDate="22.11.2004"
	res.AccommodationList.Accommodation[0].DepartureTime="17:00"
	res.AccommodationList.Accommodation[0].NumberOfRooms=1
	res.AccommodationList.Accommodation[0].NumberOfGuests=1
	res.AccommodationList.Accommodation[0].AdditionalInfo="No Smoking room"
	res.AccommodationList.Accommodation[0].ReferenceNumber="12345"
	res.AccommodationList.Accommodation[0].ToBeCancelled.Code=2
	res.AccommodationList.Accommodation[0].ToBeChanged.Code=1
	res.AccommodationList.Accommodation[0].Hotel.Code=300035
	res.AccommodationList.Accommodation[0].Product.Code=302339
	res.AccommodationList.Accommodation[0].Product.Allotment.Code="0"
	res.AccommodationList.Accommodation[0].CriticalEarlyCheckIn.Code=1
	res.AccommodationList.Accommodation[0].CriticalLateCheckOut.Code=2
	res.AccommodationList.Accommodation[0].Persons.Items=make([]acaseSts.PersonType, 1)
	res.AccommodationList.Accommodation[0].Persons.Items[0].LastName="Smith"
	res.AccommodationList.Accommodation[0].Persons.Items[0].FirstName="John"
	res.AccommodationList.Accommodation[0].Persons.Items[0].Category.Code=1
	res.AccommodationList.Accommodation[0].Persons.Items[0].Citizenship.Code=0

	return res
}

func TestOrderRequest_AccommodationAndTransfers_Ok(t *testing.T) {
	testRequest("orderrequest_response_example.xml")
	defer gock.Off()

	item := getOrderRequestObject()
	data, err := acApi.OrderRequest(item)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.Success.Id, 496446)
}

func TestOrderRequest_Error(t *testing.T) {
	testRequest("orderrequest_error_example.xml")
	defer gock.Off()

	item := getOrderRequestObject()
	_, err := acApi.OrderRequest(item)

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}

