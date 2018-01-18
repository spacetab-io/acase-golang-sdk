package test

import (
	"testing"
	"github.com/tmconsulting/acase-golang-sdk/acaseSts"
	"github.com/nbio/st"
)

func getAccommodationAndTransfersRequestObject() *acaseSts.OrderInfoNotifyRequestType {
	res := &acaseSts.OrderInfoNotifyRequestType{
		ContractNumber: "0157/13",
		ContractDate: "01.07.2013",
		GuaranteeAmount: 0.00,
		NotifierId: "Academservice",
		Password: "322223",
		Language: "en",
		Id: 546797,
		ReferenceNumber: "CYXBW",
		ReferenceNumberName: "Ref No",
		RegistrationDate: "16.03.2005",
		DeadlineDate: "15.03.2005",
		Price: 142.00,
		TravelAgencyCommission: 0.00,
		Penalty: 0.00,
		StartDate: "16.03.2005",
		EndDate: "17.03.2005",
		InvoiceRule: 3,
		InvoiceId: 7413250,
		ConfirmId: 0,
	}
	res.Currency.Code = 3
	res.Currency.Name = "USD"
	res.Autonom.Code = 1
	res.Autonom.Name = "Автономный"
	res.VatLimit.Code = 1
	res.VatLimit.Name = "Без ограничений"
	res.TypeContract.Code = 3
	res.TypeContract.Name = "Агентский"
	res.WhereToPay.Code = 3
	res.WhereToPay.Name = "Оплата Академсервису согласно договору"
	res.SubjectToProcessing.Code = 1
	res.SubjectToProcessing.Name = "Да"
	res.FinancialCondition.Code = 1
	res.FinancialCondition.Name = "Не требуется"
	res.Status.Code = 41
	res.Status.Name = "Подтвержден"
	res.ContactPerson.Name = "Смирнов Павел"
	res.ContactPerson.Phone = ""
	res.ContactPerson.Fax = ""
	res.ContactPerson.Email = "smirnov@smirnov.ru"
	res.AccountDetails.IsGain.Code = 3
	res.AccountDetails.IsGain.Name = "-"
	res.Customer.FullName = "ООО Ромашка"
	res.Customer.ZipCode = "123456"
	res.Customer.Address = "Тестовая ул. д.5/2 оф. 5"
	res.Customer.PIAddress = "Россия, 123456, Москва, Тестовая ул. д.5/2 оф. 5"
	res.Customer.INN = "1234567890"
	res.Customer.KPP = "123456789"
	res.Customer.Phone = "+7 495 123 45 56"
	res.Customer.Name = "Ромашка"
	res.Customer.Code = 1322075
	res.Customer.BuyerType.Name = "Юридическое лицо"
	res.Customer.BuyerType.Code = 3
	res.Customer.Country.Name = "Россия"
	res.Customer.Country.Code = 9
	res.Customer.City.Name = "Москва"
	res.Customer.City.Code = 2
	res.Customer.City.CityType = &acaseSts.SimpleCodeNameType{}
	res.Customer.City.CityType.Name = "Город"
	res.Customer.City.CityType.Code = 2
	res.Customer.City.AdmUnit1 = &acaseSts.SimpleCodeNameType{}
	res.Customer.City.AdmUnit1.Name = "Не определено"
	res.Customer.City.AdmUnit1.Code = 1
	res.Customer.City.AdmUnit2 = &acaseSts.AdmUnit2Type{}
	res.Customer.City.AdmUnit2.Name = "Не определено"
	res.Customer.City.AdmUnit2.Code = 1
	res.Customer.Actual = &acaseSts.YesNoCodeType{}
	res.Customer.Actual.Name = "Да"
	res.Customer.Actual.Code = 1
	res.Customer.AllowModify = &acaseSts.YesNoCodeType{}
	res.Customer.AllowModify.Name = "Да"
	res.Customer.AllowModify.Code = 1
	res.IsCustomer.Name = "Yes"
	res.IsCustomer.Code = 1
	res.AccommodationList.Accommodation = make([]acaseSts.AccommodationType, 1)
	res.AccommodationList.Accommodation[0].Id = 738746
	res.AccommodationList.Accommodation[0].ArrivalDate = "16.03.2005"
	res.AccommodationList.Accommodation[0].ArrivalTime = "10:00"
	res.AccommodationList.Accommodation[0].DepartureDate = "17.03.2005"
	res.AccommodationList.Accommodation[0].DepartureTime = "17:00"
	res.AccommodationList.Accommodation[0].NumberOfNights = 1
	res.AccommodationList.Accommodation[0].NumberOfRooms = 1
	res.AccommodationList.Accommodation[0].NumberOfGuests = 1
	res.AccommodationList.Accommodation[0].AdditionalInfo = ""
	res.AccommodationList.Accommodation[0].SupplierInfo = ""
	res.AccommodationList.Accommodation[0].ConfirmationNumber = ""
	res.AccommodationList.Accommodation[0].Price = 142.00
	res.AccommodationList.Accommodation[0].VATIncludedInPrice = 0.00
	res.AccommodationList.Accommodation[0].TravelAgencyCommission = 0.00
	res.AccommodationList.Accommodation[0].Penalty = 0.00
	res.AccommodationList.Accommodation[0].DeadlineDate = "15.03.2005"
	res.AccommodationList.Accommodation[0].DeadlineTimeLoc = "15.03.2005 13:00"
	res.AccommodationList.Accommodation[0].DeadlineTimeSys = "15.03.2005 13:00"
	res.AccommodationList.Accommodation[0].DeadlineTimeUTC = "15.03.2005 10:00"
	res.AccommodationList.Accommodation[0].PossiblePenaltySize = 142.00
	res.AccommodationList.Accommodation[0].ToBeCancelled.Code=2
	res.AccommodationList.Accommodation[0].ToBeCancelled.Name="Нет"
	res.AccommodationList.Accommodation[0].ToBeChanged.Code=2
	res.AccommodationList.Accommodation[0].ToBeChanged.Name="Нет"
	res.AccommodationList.Accommodation[0].Hotel.Code=300061
	res.AccommodationList.Accommodation[0].Hotel.Name="Русь (Солнечный)"
	res.AccommodationList.Accommodation[0].ObjType.Code=2
	res.AccommodationList.Accommodation[0].ObjType.Name="Гостиница"
	res.AccommodationList.Accommodation[0].ObjType.NameWhere="гостинице"
	res.AccommodationList.Accommodation[0].Country.Code=9
	res.AccommodationList.Accommodation[0].Country.Name="Россия"
	res.AccommodationList.Accommodation[0].City.Code=2
	res.AccommodationList.Accommodation[0].City.Name="Москва"
	res.AccommodationList.Accommodation[0].Status.Code=41
	res.AccommodationList.Accommodation[0].Status.Name="Подтверждено"
	res.AccommodationList.Accommodation[0].AmendmentRejected.Code=2
	res.AccommodationList.Accommodation[0].AmendmentRejected.Name="Нет"
	res.AccommodationList.Accommodation[0].AllowCancel.Code=1
	res.AccommodationList.Accommodation[0].AllowCancel.Name="Да"
	res.AccommodationList.Accommodation[0].AllowModify.Code=1
	res.AccommodationList.Accommodation[0].AllowModify.Name="Да"
	res.AccommodationList.Accommodation[0].Product.Code=410167
	res.AccommodationList.Accommodation[0].Product.RoomName="Junior Suite"
	res.AccommodationList.Accommodation[0].Product.RoomCode=900153
	res.AccommodationList.Accommodation[0].Product.RateCode=900185
	res.AccommodationList.Accommodation[0].Product.RateName="FIT"
	res.AccommodationList.Accommodation[0].Product.RateDescription=""
	res.AccommodationList.Accommodation[0].Product.RateGroupCode=4
	res.AccommodationList.Accommodation[0].Product.RateGroupName="FIT"
	res.AccommodationList.Accommodation[0].Product.Allotment.Code="500235"
	res.AccommodationList.Accommodation[0].Product.Allotment.Name="РУСЬ-ДЛЯ ВСЕХ"
	res.AccommodationList.Accommodation[0].Meal.Code=0
	res.AccommodationList.Accommodation[0].Meal.Name=""
	res.AccommodationList.Accommodation[0].Persons.Items = make([]acaseSts.PersonType, 1)
	res.AccommodationList.Accommodation[0].Persons.Items[0].LastName="ИВАНОВ"
	res.AccommodationList.Accommodation[0].Persons.Items[0].FirstName="ИВАН ИВАНОВИЧ"
	res.AccommodationList.Accommodation[0].Persons.Items[0].Category.Code=1
	res.AccommodationList.Accommodation[0].Persons.Items[0].Category.Name="Господин"
	res.AccommodationList.Accommodation[0].Persons.Items[0].Citizenship.Code=0
	res.AccommodationList.Accommodation[0].Persons.Items[0].Citizenship.Name=""
	res.AccommodationList.Accommodation[0].IsEarlyCheckIn.Code=1
	res.AccommodationList.Accommodation[0].IsEarlyCheckIn.Name="Да"
	res.AccommodationList.Accommodation[0].CriticalEarlyCheckIn.Code=1
	res.AccommodationList.Accommodation[0].CriticalEarlyCheckIn.Name="Критично"
	res.AccommodationList.Accommodation[0].EarlyCheckInConfirmationStatus.Code=2
	res.AccommodationList.Accommodation[0].EarlyCheckInConfirmationStatus.Name="Не подтвержден"
	res.AccommodationList.Accommodation[0].IsHour.Code=2
	res.AccommodationList.Accommodation[0].IsHour.Name="Нет"
	res.AccommodationList.Accommodation[0].IsLateCheckOut.Code=1
	res.AccommodationList.Accommodation[0].IsLateCheckOut.Name="Да"
	res.AccommodationList.Accommodation[0].CriticalLateCheckOut.Code=2
	res.AccommodationList.Accommodation[0].CriticalLateCheckOut.Name="Некритично"
	res.AccommodationList.Accommodation[0].LateCheckOutConfirmationStatus.Code=2
	res.AccommodationList.Accommodation[0].LateCheckOutConfirmationStatus.Name="Не подтвержден"

	return res
}

func getAgencyAgreementsWithAdditionalBenefitsRequestObject() *acaseSts.OrderInfoNotifyRequestType {
	res := &acaseSts.OrderInfoNotifyRequestType{
		ContractNumber: "0157/13",
		ContractDate: "01.07.2013",
		GuaranteeAmount: 0.00,
		NotifierId: "Academservice",
		Password: "322223",
		Language: "en",
		Id: 546797,
		ReferenceNumber: "CYXBW",
		ReferenceNumberName: "Ref No",
		RegistrationDate: "16.03.2005",
		DeadlineDate: "15.03.2005",
		Price: 142.00,
		TravelAgencyCommission: 0.00,
		Penalty: 0.00,
		StartDate: "16.03.2005",
		EndDate: "17.03.2005",
		InvoiceRule: 3,
		InvoiceId: 7413250,
		ConfirmId: 0,
	}
	res.Currency.Code = 3
	res.Currency.Name = "USD"
	res.Autonom.Code = 1
	res.Autonom.Name = "Автономный"
	res.VatLimit.Code = 1
	res.VatLimit.Name = "Без ограничений"
	res.TypeContract.Code = 3
	res.TypeContract.Name = "Агентский"
	res.WhereToPay.Code = 3
	res.WhereToPay.Name = "Оплата Академсервису согласно договору"
	res.SubjectToProcessing.Code = 1
	res.SubjectToProcessing.Name = "Да"
	res.FinancialCondition.Code = 1
	res.FinancialCondition.Name = "Не требуется"
	res.Status.Code = 41
	res.Status.Name = "Подтвержден"
	res.ContactPerson.Name = "Смирнов Павел"
	res.ContactPerson.Phone = ""
	res.ContactPerson.Fax = ""
	res.ContactPerson.Email = "smirnov@smirnov.ru"
	res.AccountDetails.IsGain.Code = 3
	res.AccountDetails.IsGain.Name = "-"
	res.Customer.FullName = "ООО Ромашка"
	res.Customer.ZipCode = "123456"
	res.Customer.Address = "Тестовая ул. д.5/2 оф. 5"
	res.Customer.PIAddress = "Россия, 123456, Москва, Тестовая ул. д.5/2 оф. 5"
	res.Customer.INN = "1234567890"
	res.Customer.KPP = "123456789"
	res.Customer.Phone = "+7 495 123 45 56"
	res.Customer.Name = "Ромашка"
	res.Customer.Code = 1322075
	res.Customer.BuyerType.Name = "Юридическое лицо"
	res.Customer.BuyerType.Code = 3
	res.Customer.Country.Name = "Россия"
	res.Customer.Country.Code = 9
	res.Customer.City.Name = "Москва"
	res.Customer.City.Code = 2
	res.Customer.City.CityType = &acaseSts.SimpleCodeNameType{}
	res.Customer.City.CityType.Name = "Город"
	res.Customer.City.CityType.Code = 2
	res.Customer.City.AdmUnit1 = &acaseSts.SimpleCodeNameType{}
	res.Customer.City.AdmUnit1.Name = "Не определено"
	res.Customer.City.AdmUnit1.Code = 1
	res.Customer.City.AdmUnit2 = &acaseSts.AdmUnit2Type{}
	res.Customer.City.AdmUnit2.Name = "Не определено"
	res.Customer.City.AdmUnit2.Code = 1
	res.Customer.Actual = &acaseSts.YesNoCodeType{}
	res.Customer.Actual.Name = "Да"
	res.Customer.Actual.Code = 1
	res.Customer.AllowModify = &acaseSts.YesNoCodeType{}
	res.Customer.AllowModify.Name = "Да"
	res.Customer.AllowModify.Code = 1
	res.IsCustomer.Name = "Yes"
	res.IsCustomer.Code = 1
	res.AccommodationList.Accommodation = make([]acaseSts.AccommodationType, 1)
	res.AccommodationList.Accommodation[0].Id = 738746
	res.AccommodationList.Accommodation[0].ArrivalDate = "16.03.2005"
	res.AccommodationList.Accommodation[0].ArrivalTime = "10:00"
	res.AccommodationList.Accommodation[0].DepartureDate = "17.03.2005"
	res.AccommodationList.Accommodation[0].DepartureTime = "17:00"
	res.AccommodationList.Accommodation[0].NumberOfNights = 1
	res.AccommodationList.Accommodation[0].NumberOfRooms = 1
	res.AccommodationList.Accommodation[0].NumberOfGuests = 1
	res.AccommodationList.Accommodation[0].AdditionalInfo = ""
	res.AccommodationList.Accommodation[0].SupplierInfo = ""
	res.AccommodationList.Accommodation[0].ConfirmationNumber = ""
	res.AccommodationList.Accommodation[0].Price = 142.00
	res.AccommodationList.Accommodation[0].VATIncludedInPrice = 0.00
	res.AccommodationList.Accommodation[0].TravelAgencyCommission = 0.00
	res.AccommodationList.Accommodation[0].Gain = 1212.00
	res.AccommodationList.Accommodation[0].PenaltyGain = 270
	res.AccommodationList.Accommodation[0].Penalty = 2500.00
	res.AccommodationList.Accommodation[0].DeadlineDate = "15.03.2005"
	res.AccommodationList.Accommodation[0].DeadlineTimeLoc = "15.03.2005 13:00"
	res.AccommodationList.Accommodation[0].DeadlineTimeSys = "15.03.2005 13:00"
	res.AccommodationList.Accommodation[0].DeadlineTimeUTC = "15.03.2005 10:00"
	res.AccommodationList.Accommodation[0].PossiblePenaltySize = 142.00
	res.AccommodationList.Accommodation[0].ToBeCancelled.Code=2
	res.AccommodationList.Accommodation[0].ToBeCancelled.Name="Нет"
	res.AccommodationList.Accommodation[0].ToBeChanged.Code=2
	res.AccommodationList.Accommodation[0].ToBeChanged.Name="Нет"
	res.AccommodationList.Accommodation[0].Hotel.Code=300061
	res.AccommodationList.Accommodation[0].Hotel.Name="Русь (Солнечный)"
	res.AccommodationList.Accommodation[0].ObjType.Code=2
	res.AccommodationList.Accommodation[0].ObjType.Name="Гостиница"
	res.AccommodationList.Accommodation[0].ObjType.NameWhere="гостинице"
	res.AccommodationList.Accommodation[0].Country.Code=9
	res.AccommodationList.Accommodation[0].Country.Name="Россия"
	res.AccommodationList.Accommodation[0].City.Code=2
	res.AccommodationList.Accommodation[0].City.Name="Москва"
	res.AccommodationList.Accommodation[0].Status.Code=41
	res.AccommodationList.Accommodation[0].Status.Name="Подтверждено"
	res.AccommodationList.Accommodation[0].AmendmentRejected.Code=2
	res.AccommodationList.Accommodation[0].AmendmentRejected.Name="Нет"
	res.AccommodationList.Accommodation[0].AllowCancel.Code=1
	res.AccommodationList.Accommodation[0].AllowCancel.Name="Да"
	res.AccommodationList.Accommodation[0].AllowModify.Code=1
	res.AccommodationList.Accommodation[0].AllowModify.Name="Да"
	res.AccommodationList.Accommodation[0].AllowSetGain = &acaseSts.YesNoCodeType{Code:1,Name:"Да"}
	res.AccommodationList.Accommodation[0].Product.Code=410167
	res.AccommodationList.Accommodation[0].Product.RoomName="Junior Suite"
	res.AccommodationList.Accommodation[0].Product.RoomCode=900153
	res.AccommodationList.Accommodation[0].Product.RateCode=900185
	res.AccommodationList.Accommodation[0].Product.RateName="FIT"
	res.AccommodationList.Accommodation[0].Product.RateDescription=""
	res.AccommodationList.Accommodation[0].Product.RateGroupCode=4
	res.AccommodationList.Accommodation[0].Product.RateGroupName="FIT"
	res.AccommodationList.Accommodation[0].Product.Allotment.Code="500235"
	res.AccommodationList.Accommodation[0].Product.Allotment.Name="РУСЬ-ДЛЯ ВСЕХ"
	res.AccommodationList.Accommodation[0].Meal.Code=0
	res.AccommodationList.Accommodation[0].Meal.Name=""
	res.AccommodationList.Accommodation[0].Persons.Items = make([]acaseSts.PersonType, 1)
	res.AccommodationList.Accommodation[0].Persons.Items[0].LastName="ИВАНОВ"
	res.AccommodationList.Accommodation[0].Persons.Items[0].FirstName="ИВАН ИВАНОВИЧ"
	res.AccommodationList.Accommodation[0].Persons.Items[0].Category.Code=1
	res.AccommodationList.Accommodation[0].Persons.Items[0].Category.Name="Господин"
	res.AccommodationList.Accommodation[0].Persons.Items[0].Citizenship.Code=0
	res.AccommodationList.Accommodation[0].Persons.Items[0].Citizenship.Name=""
	res.AccommodationList.Accommodation[0].Penalties.Items = make([]acaseSts.PenaltyType, 2)
	res.AccommodationList.Accommodation[0].Penalties.Items[0].Id = 1
	res.AccommodationList.Accommodation[0].Penalties.Items[0].VATIncludedInPrice = 0.00
	res.AccommodationList.Accommodation[0].Penalties.Items[0].FullVATInPrice = 1
	res.AccommodationList.Accommodation[0].Penalties.Items[0].Price = 1000.00
	res.AccommodationList.Accommodation[0].Penalties.Items[0].Gain = 150.00
	res.AccommodationList.Accommodation[0].Penalties.Items[0].Reason.Code = 3
	res.AccommodationList.Accommodation[0].Penalties.Items[0].Reason.Name = "Штраф за позднее изменение"
	res.AccommodationList.Accommodation[0].Penalties.Items[0].AllowSetGain.Code = 1
	res.AccommodationList.Accommodation[0].Penalties.Items[0].AllowSetGain.Name = "Да"
	res.AccommodationList.Accommodation[0].Penalties.Items[0].Id = 2
	res.AccommodationList.Accommodation[0].Penalties.Items[0].VATIncludedInPrice = 0.00
	res.AccommodationList.Accommodation[0].Penalties.Items[0].FullVATInPrice = 1
	res.AccommodationList.Accommodation[0].Penalties.Items[0].Price = 1500.00
	res.AccommodationList.Accommodation[0].Penalties.Items[0].Gain = 120.00
	res.AccommodationList.Accommodation[0].Penalties.Items[0].Reason.Code = 3
	res.AccommodationList.Accommodation[0].Penalties.Items[0].Reason.Name = "Штраф за позднее изменение"
	res.AccommodationList.Accommodation[0].Penalties.Items[0].AllowSetGain.Code = 1
	res.AccommodationList.Accommodation[0].Penalties.Items[0].AllowSetGain.Name = "Да"
	res.AccommodationList.Accommodation[0].IsEarlyCheckIn.Code=1
	res.AccommodationList.Accommodation[0].IsEarlyCheckIn.Name="Да"
	res.AccommodationList.Accommodation[0].CriticalEarlyCheckIn.Code=1
	res.AccommodationList.Accommodation[0].CriticalEarlyCheckIn.Name="Критично"
	res.AccommodationList.Accommodation[0].EarlyCheckInConfirmationStatus.Code=2
	res.AccommodationList.Accommodation[0].EarlyCheckInConfirmationStatus.Name="Не подтвержден"
	res.AccommodationList.Accommodation[0].IsHour.Code=2
	res.AccommodationList.Accommodation[0].IsHour.Name="Нет"
	res.AccommodationList.Accommodation[0].IsLateCheckOut.Code=1
	res.AccommodationList.Accommodation[0].IsLateCheckOut.Name="Да"
	res.AccommodationList.Accommodation[0].CriticalLateCheckOut.Code=2
	res.AccommodationList.Accommodation[0].CriticalLateCheckOut.Name="Некритично"
	res.AccommodationList.Accommodation[0].LateCheckOutConfirmationStatus.Code=2
	res.AccommodationList.Accommodation[0].LateCheckOutConfirmationStatus.Name="Не подтвержден"

	return res
}

/*
func TestGenerateXml_AccommodationAndTransfers_Ok(t *testing.T) {
	item := getAccommodationAndTransfersRequestObject()
	bItem, err := xml.MarshalIndent(item, "", "    ")
	if err != nil {
		t.Error(err)
	}
	t.Log(xml.Header + string(bItem))
}
*/

func TestOrderNotifyRequest_AccommodationAndTransfers_Ok(t *testing.T) {
	testRequest("ordinfonotify_accommtransfers_response_example.xml", false)
	defer gockOff()

	item := getAccommodationAndTransfersRequestObject()
	_, err := acApi.OrderInfoNotifyRequest(item)
	er := getCustomErrorType()
	st.Expect(t, err, er)
}

func TestOrderNotifyRequest_AgencyAgreementsWithAdditionalBenefits_Ok(t *testing.T) {
	testRequest("orderinfonotify_agencyagreements_response_example.xml", false)
	defer gockOff()

	item := getAgencyAgreementsWithAdditionalBenefitsRequestObject()
	_, err := acApi.OrderInfoNotifyRequest(item)
	er := getCustomErrorType()
	st.Expect(t, err, er)
}

func TestOrderNotifyRequest_Error(t *testing.T) {
	testRequest("orderinfonotify_error_example.xml", true)
	defer gockOff()

	item := getAgencyAgreementsWithAdditionalBenefitsRequestObject()
	_, err := acApi.OrderInfoNotifyRequest(item)

	st.Expect(t, err.Message, "Доступ запрещен !")
}
