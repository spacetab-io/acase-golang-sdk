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

func getOrderRequestAgencyAgreementsWithAdditionalBenefitsObject() *acaseSts.OrderRequestType {
	res := &acaseSts.OrderRequestType{}
	res.BuyerId="MyCompanyId"
	res.UserId="MyUserId"
	res.Password="MyPassword"
	res.Language="ru"
	res.Id=4512345
	res.ReferenceNumber=""
	res.Currency.Code=2
	res.WhereToPay.Code=3
	res.AccommodationList.Accommodation = make([]acaseSts.AccommodationType, 2)
	res.AccommodationList.Accommodation[0].Id=5567891
	res.AccommodationList.Accommodation[0].ArrivalDate="17.05.2017"
	res.AccommodationList.Accommodation[0].ArrivalTime="14:00"
	res.AccommodationList.Accommodation[0].DepartureDate="18.05.2017"
	res.AccommodationList.Accommodation[0].DepartureTime="12:00"
	res.AccommodationList.Accommodation[0].NumberOfRooms=1
	res.AccommodationList.Accommodation[0].NumberOfGuests=1
	res.AccommodationList.Accommodation[0].Gain=1212.00
	res.AccommodationList.Accommodation[0].AdditionalInfo="No Smoking room"
	res.AccommodationList.Accommodation[0].ReferenceNumber=""
	res.AccommodationList.Accommodation[0].ToBeCancelled.Code=2
	res.AccommodationList.Accommodation[0].ToBeChanged.Code=1
	res.AccommodationList.Accommodation[0].Hotel.Code=9800032
	res.AccommodationList.Accommodation[0].Product.Code=1134385
	res.AccommodationList.Accommodation[0].Product.Allotment.Code="9500001+1499424+9500001+2"
	res.AccommodationList.Accommodation[0].CriticalEarlyCheckIn.Code=3
	res.AccommodationList.Accommodation[0].CriticalLateCheckOut.Code=3
	res.AccommodationList.Accommodation[0].Persons.Items=make([]acaseSts.PersonType, 1)
	res.AccommodationList.Accommodation[0].Persons.Items[0].LastName="Test"
	res.AccommodationList.Accommodation[0].Persons.Items[0].FirstName="t"
	res.AccommodationList.Accommodation[0].Persons.Items[0].Category.Code=1
	res.AccommodationList.Accommodation[0].Persons.Items[0].Citizenship.Code=0
	res.AccommodationList.Accommodation[0].Penalties.Items=make([]acaseSts.PenaltyType, 2)
	res.AccommodationList.Accommodation[0].Penalties.Items[0].Id=5614785
	res.AccommodationList.Accommodation[0].Penalties.Items[0].Gain=120.00
	res.AccommodationList.Accommodation[0].Penalties.Items[1].Id=5614786
	res.AccommodationList.Accommodation[0].Penalties.Items[1].Gain=150.00
	res.AccommodationList.Accommodation[1].Id=0
	res.AccommodationList.Accommodation[1].ArrivalDate="17.05.2017"
	res.AccommodationList.Accommodation[1].ArrivalTime="14:00"
	res.AccommodationList.Accommodation[1].DepartureDate="18.05.2017"
	res.AccommodationList.Accommodation[1].DepartureTime="12:00"
	res.AccommodationList.Accommodation[1].NumberOfRooms=1
	res.AccommodationList.Accommodation[1].NumberOfGuests=1
	res.AccommodationList.Accommodation[1].Gain=1212.00
	res.AccommodationList.Accommodation[1].AdditionalInfo="No Smoking room"
	res.AccommodationList.Accommodation[1].ReferenceNumber=""
	res.AccommodationList.Accommodation[1].ToBeCancelled.Code=2
	res.AccommodationList.Accommodation[1].ToBeChanged.Code=1
	res.AccommodationList.Accommodation[1].Hotel.Code=9800032
	res.AccommodationList.Accommodation[1].Product.Code=1134385
	res.AccommodationList.Accommodation[1].Product.Allotment.Code="9500001+1499424+9500001+2"
	res.AccommodationList.Accommodation[1].CriticalEarlyCheckIn.Code=3
	res.AccommodationList.Accommodation[1].CriticalLateCheckOut.Code=3
	res.AccommodationList.Accommodation[1].Persons.Items=make([]acaseSts.PersonType, 1)
	res.AccommodationList.Accommodation[1].Persons.Items[0].LastName="Test"
	res.AccommodationList.Accommodation[1].Persons.Items[0].FirstName="tt"
	res.AccommodationList.Accommodation[1].Persons.Items[0].Category.Code=1
	res.AccommodationList.Accommodation[1].Persons.Items[0].Citizenship.Code=0

	return res
}

func getOrderAwocRequestObject() *acaseSts.OrderAwocRequestType {
	res := &acaseSts.OrderAwocRequestType{}
	res.BuyerId="MyCompanyId"
	res.UserId="MyUserId"
	res.Password="MyPassword"
	res.Language="ru"
	res.Id=0
	res.ReferenceNumber=""
	res.Currency.Code=2
	res.WhereToPay.Code=3
	res.AwocList.Items = make([]acaseSts.AwocType, 1)
	res.AwocList.Items[0].Id=0
	res.AwocList.Items[0].Subject=""
	res.AwocList.Items[0].ArrivalDate="01.09.2013"
	res.AwocList.Items[0].DepartureDate="03.09.2013"
	res.AwocList.Items[0].DepartureTime="17:00"
	res.AwocList.Items[0].ReferenceNumber=""
	res.AwocList.Items[0].ToBeCancelled.Code=2
	res.AwocList.Items[0].ToBeChanged.Code=1
	res.AwocList.Items[0].Hotel.Code=1300725
	res.AwocList.Items[0].DetailList.Items=make([]acaseSts.DetailType, 1)
	res.AwocList.Items[0].DetailList.Items[0].Id=0
	res.AwocList.Items[0].DetailList.Items[0].Subject="Двухместный номер с большой кроватью. Завтрак."
	res.AwocList.Items[0].DetailList.Items[0].Quantity=1
	res.AwocList.Items[0].DetailList.Items[0].Duration=2.00
	res.AwocList.Items[0].Persons.Items=make([]acaseSts.PersonType, 2)
	res.AwocList.Items[0].Persons.Items[0].LastName="Петров"
	res.AwocList.Items[0].Persons.Items[0].FirstName="Илья"
	res.AwocList.Items[0].Persons.Items[0].Category.Code=1
	res.AwocList.Items[0].Persons.Items[0].Citizenship.Code=0
	res.AwocList.Items[0].Persons.Items[1].LastName="Петрова"
	res.AwocList.Items[0].Persons.Items[1].FirstName="Елена"
	res.AwocList.Items[0].Persons.Items[1].Category.Code=2
	res.AwocList.Items[0].Persons.Items[1].Citizenship.Code=0

	return res
}

func TestOrderAwocRequest_Ok(t *testing.T) {
	testRequest("orderrequest_response_example.xml")
	defer gock.Off()

	item := getOrderAwocRequestObject()
	data, err := acApi.OrderAwocRequest(item)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.Success.Id, 496446)
}

func TestOrderRequest_AgencyAgreementsWithAdditionalBenefits_Ok(t *testing.T) {
	testRequest("orderrequest_response_example.xml")
	defer gock.Off()

	item := getOrderRequestAgencyAgreementsWithAdditionalBenefitsObject()
	data, err := acApi.OrderRequest(item)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.Success.Id, 496446)
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

