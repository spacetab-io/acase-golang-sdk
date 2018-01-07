package test

import (
	"testing"
	"gopkg.in/h2non/gock.v1"
	"github.com/tmconsulting/acase-golang-sdk/acaseSts"
	"github.com/nbio/st"
)

func getRequestPriceOfferRequestObject() *acaseSts.OrderInfoAwocNotifyRequestType {
	res := &acaseSts.OrderInfoAwocNotifyRequestType{}
	res.Id=2736033
	res.ReferenceNumber=""
	res.ReferenceNumberName="Ref No"
	res.RegistrationDate="12.08.2013"
	res.DeadlineDate="30.08.2013"
	res.Price=0.00
	res.TravelAgencyCommission=0.00
	res.Penalty=0.00
	res.StartDate="01.09.2013"
	res.EndDate="03.09.2013"
	res.InvoiceRule=2
	res.InvoiceId=0
	res.ConfirmId=0
	res.Currency.Code=2
	res.Currency.Name="RUB"
	res.WhereToPay.Code=3
	res.WhereToPay.Name="Оплата ООО \"АКАДЕМСЕРВИС\" согласно договору"
	res.Status.Code=5
	res.Status.Name="Предложение цены"
	res.ContactPerson.Name="test"
	res.ContactPerson.Phone=""
	res.ContactPerson.Fax=""
	res.ContactPerson.Email="test-mail@mail.ru"
	res.Customer.FullName="ООО Ромашка"
	res.Customer.ZipCode="123456"
	res.Customer.Address="Тестовая ул. д.5/2 оф. 5"
	res.Customer.PIAddress="Россия, 123456, Москва, Тестовая ул. д.5/2 оф. 5"
	res.Customer.INN="1234567890"
	res.Customer.KPP="123456789"
	res.Customer.Phone="+7 495 123 45 56"
	res.Customer.Name="Ромашка"
	res.Customer.Code=1322075
	res.Customer.BuyerType.Name="Юридическое лицо"
	res.Customer.BuyerType.Code=3
	res.Customer.Country.Name="Россия"
	res.Customer.Country.Code=9
	res.Customer.City.Name="Москва"
	res.Customer.City.Code=2
	res.Customer.City.CityType = &acaseSts.CityTypeType{}
	res.Customer.City.CityType.Name="Город"
	res.Customer.City.CityType.Code=2
	res.Customer.City.AdmUnit1 = &acaseSts.AdmUnit1Type{}
	res.Customer.City.AdmUnit1.Name="Не определено"
	res.Customer.City.AdmUnit1.Code=1
	res.Customer.City.AdmUnit2 = &acaseSts.AdmUnit2Type{}
	res.Customer.City.AdmUnit2.Name="Не определено"
	res.Customer.City.AdmUnit2.Code=1
	res.Customer.Actual = &acaseSts.ActualType{}
	res.Customer.Actual.Name="Да"
	res.Customer.Actual.Code=1
	res.Customer.AllowModify = &acaseSts.AllowModifyType{}
	res.Customer.AllowModify.Name="Да"
	res.Customer.AllowModify.Code=1
	res.IsCustomer.Code = 1
	res.IsCustomer.Name = "Да"
	res.AwocList.Items = make([]acaseSts.AwocType, 2)
	res.AwocList.Items[0].Id=3807880
	res.AwocList.Items[0].ParentId=0
	res.AwocList.Items[0].Subject="Гостиница Уют (Лапшинка) с 01.09.2013 по 03.09.2013, ПЕТРОВ ИЛЬЯ, ПЕТРОВА ЕЛЕНА"
	res.AwocList.Items[0].ArrivalDate="01.09.2013"
	res.AwocList.Items[0].DepartureDate="03.09.2013"
	res.AwocList.Items[0].NumberOfNights=2
	res.AwocList.Items[0].ArrivalTime="12:00"
	res.AwocList.Items[0].DepartureTime="12:00"
	res.AwocList.Items[0].Price=7132.00
	res.AwocList.Items[0].VATIncludedInPrice=1087.93
	res.AwocList.Items[0].TravelAgencyCommission=0.00
	res.AwocList.Items[0].DeadlineDate="30.08.2013"
	res.AwocList.Items[0].PossiblePenaltySize=3105.00
	res.AwocList.Items[0].Penalty=0.00
	res.AwocList.Items[0].ConfirmationNumber=""
	res.AwocList.Items[0].ReferenceNumber=""
	res.AwocList.Items[0].SupplierInfo=""
	res.AwocList.Items[0].ToBeCancelled.Code=2
	res.AwocList.Items[0].ToBeCancelled.Name="Нет"
	res.AwocList.Items[0].ToBeChanged.Code=2
	res.AwocList.Items[0].ToBeChanged.Name="Нет"
	res.AwocList.Items[0].ToBeSelected.Code=2
	res.AwocList.Items[0].ToBeSelected.Name="Нет"
	res.AwocList.Items[0].Hotel.Code=1300725
	res.AwocList.Items[0].Hotel.Name="Уют Внуково МО"
	res.AwocList.Items[0].Country.Code=9
	res.AwocList.Items[0].Country.Name="Россия"
	res.AwocList.Items[0].AdmUnit1.Code=1
	res.AwocList.Items[0].AdmUnit1.Name="Не определено"
	res.AwocList.Items[0].AdmUnit2.Code=1
	res.AwocList.Items[0].AdmUnit2.Name="Не определено"
	res.AwocList.Items[0].TypeOfPlace.Code=2
	res.AwocList.Items[0].TypeOfPlace.Name="Город"
	res.AwocList.Items[0].City.Code=5254
	res.AwocList.Items[0].City.Name="Лапшинка"
	res.AwocList.Items[0].Status.Code=5
	res.AwocList.Items[0].Status.Name="Предложение цены"
	res.AwocList.Items[0].AllowCancel.Code=1
	res.AwocList.Items[0].AllowCancel.Name="Да"
	res.AwocList.Items[0].AllowModify.Code=2
	res.AwocList.Items[0].AllowModify.Name="Нет"
	res.AwocList.Items[0].AllowSelect.Code=1
	res.AwocList.Items[0].AllowSelect.Name="Да"
	res.AwocList.Items[0].DetailList.Items = make([]acaseSts.DetailType, 2)
	res.AwocList.Items[0].DetailList.Items[0].Id=3807881
	res.AwocList.Items[0].DetailList.Items[0].Subject="Стандартный двухместный номер с большой кроватью."
	res.AwocList.Items[0].DetailList.Items[0].Quantity=1
	res.AwocList.Items[0].DetailList.Items[0].Duration=2.00
	res.AwocList.Items[0].DetailList.Items[0].Price=5980.00
	res.AwocList.Items[0].DetailList.Items[0].TravelAgencyCommission=0.00
	res.AwocList.Items[0].DetailList.Items[0].VATIncludedInPrice=912.20
	res.AwocList.Items[0].DetailList.Items[1].Id=3807886
	res.AwocList.Items[0].DetailList.Items[1].Subject="Завтрак ШС"
	res.AwocList.Items[0].DetailList.Items[1].Quantity=2
	res.AwocList.Items[0].DetailList.Items[1].Duration=2.00
	res.AwocList.Items[0].DetailList.Items[1].Price=1152.00
	res.AwocList.Items[0].DetailList.Items[1].TravelAgencyCommission=0.00
	res.AwocList.Items[0].DetailList.Items[1].VATIncludedInPrice=175.73
	res.AwocList.Items[0].Persons.Items = make([]acaseSts.PersonType, 2)
	res.AwocList.Items[0].Persons.Items[0].LastName="ПЕТРОВ"
	res.AwocList.Items[0].Persons.Items[0].FirstName="ИЛЬЯ"
	res.AwocList.Items[0].Persons.Items[0].Company=""
	res.AwocList.Items[0].Persons.Items[0].Position=""
	res.AwocList.Items[0].Persons.Items[0].Category.Code=1
	res.AwocList.Items[0].Persons.Items[0].Category.Name="Господин"
	res.AwocList.Items[0].Persons.Items[0].Citizenship.Code=1
	res.AwocList.Items[0].Persons.Items[0].Citizenship.Name="Hе имеет значения"
	res.AwocList.Items[0].Persons.Items[1].LastName="ПЕТРОВА"
	res.AwocList.Items[0].Persons.Items[1].FirstName="ЕЛЕНА"
	res.AwocList.Items[0].Persons.Items[1].Company=""
	res.AwocList.Items[0].Persons.Items[1].Position=""
	res.AwocList.Items[0].Persons.Items[1].Category.Code=2
	res.AwocList.Items[0].Persons.Items[1].Category.Name="Госпожа"
	res.AwocList.Items[0].Persons.Items[1].Citizenship.Code=1
	res.AwocList.Items[0].Persons.Items[1].Citizenship.Name="Hе имеет значения"
	res.AwocList.Items[1].Id=3807882
	res.AwocList.Items[1].ParentId=3807880
	res.AwocList.Items[1].Subject="Гостиница Уют (Лапшинка) с 01.09.2013 по 03.09.2013, ПЕТРОВ ИЛЬЯ, ПЕТРОВА ЕЛЕНА"
	res.AwocList.Items[1].ArrivalDate="01.09.2013"
	res.AwocList.Items[1].DepartureDate="03.09.2013"
	res.AwocList.Items[1].NumberOfNights=2
	res.AwocList.Items[1].ArrivalTime="12:00"
	res.AwocList.Items[1].DepartureTime="12:00"
	res.AwocList.Items[1].Price=7708.00
	res.AwocList.Items[1].VATIncludedInPrice=1175.79
	res.AwocList.Items[1].TravelAgencyCommission=0.00
	res.AwocList.Items[1].DeadlineDate="30.08.2013"
	res.AwocList.Items[1].PossiblePenaltySize=3393.00
	res.AwocList.Items[1].Penalty=0.00
	res.AwocList.Items[1].ConfirmationNumber=""
	res.AwocList.Items[1].ReferenceNumber=""
	res.AwocList.Items[1].SupplierInfo=""
	res.AwocList.Items[1].ToBeCancelled.Code=2
	res.AwocList.Items[1].ToBeCancelled.Name="Нет"
	res.AwocList.Items[1].ToBeChanged.Code=2
	res.AwocList.Items[1].ToBeChanged.Name="Нет"
	res.AwocList.Items[1].ToBeSelected.Code=2
	res.AwocList.Items[1].ToBeSelected.Name="Нет"
	res.AwocList.Items[1].Hotel.Code=1300725
	res.AwocList.Items[1].Hotel.Name="Уют Внуково МО"
	res.AwocList.Items[1].Country.Code=9
	res.AwocList.Items[1].Country.Name="Россия"
	res.AwocList.Items[1].AdmUnit1.Code=1
	res.AwocList.Items[1].AdmUnit1.Name="Не определено"
	res.AwocList.Items[1].AdmUnit2.Code=1
	res.AwocList.Items[1].AdmUnit2.Name="Не определено"
	res.AwocList.Items[1].TypeOfPlace.Code=2
	res.AwocList.Items[1].TypeOfPlace.Name="Город"
	res.AwocList.Items[1].City.Code=5254
	res.AwocList.Items[1].City.Name="Лапшинка"
	res.AwocList.Items[1].Status.Code=5
	res.AwocList.Items[1].Status.Name="Предложение цены"
	res.AwocList.Items[1].AllowCancel.Code=1
	res.AwocList.Items[1].AllowCancel.Name="Да"
	res.AwocList.Items[1].AllowModify.Code=2
	res.AwocList.Items[1].AllowModify.Name="Нет"
	res.AwocList.Items[1].AllowSelect.Code=1
	res.AwocList.Items[1].AllowSelect.Name="Да"
	res.AwocList.Items[1].DetailList.Items = make([]acaseSts.DetailType, 2)
	res.AwocList.Items[1].DetailList.Items[0].Id=3807883
	res.AwocList.Items[1].DetailList.Items[0].Subject="Улучшенный двухместный номер с большой кроватью."
	res.AwocList.Items[1].DetailList.Items[0].Quantity=1
	res.AwocList.Items[1].DetailList.Items[0].Duration=2.00
	res.AwocList.Items[1].DetailList.Items[0].Price=6556.00
	res.AwocList.Items[1].DetailList.Items[0].TravelAgencyCommission=0.00
	res.AwocList.Items[1].DetailList.Items[0].VATIncludedInPrice=1000.06
	res.AwocList.Items[1].DetailList.Items[1].Id=3807887
	res.AwocList.Items[1].DetailList.Items[1].Subject="Завтрак ШС"
	res.AwocList.Items[1].DetailList.Items[1].Quantity=2
	res.AwocList.Items[1].DetailList.Items[1].Duration=2.00
	res.AwocList.Items[1].DetailList.Items[1].Price=1152.00
	res.AwocList.Items[1].DetailList.Items[1].TravelAgencyCommission=0.00
	res.AwocList.Items[1].DetailList.Items[1].VATIncludedInPrice=175.73
	res.AwocList.Items[1].Persons.Items = make([]acaseSts.PersonType, 2)
	res.AwocList.Items[1].Persons.Items[0].LastName="ПЕТРОВ"
	res.AwocList.Items[1].Persons.Items[0].FirstName="ИЛЬЯ"
	res.AwocList.Items[1].Persons.Items[0].Company=""
	res.AwocList.Items[1].Persons.Items[0].Position=""
	res.AwocList.Items[1].Persons.Items[0].Category.Code=1
	res.AwocList.Items[1].Persons.Items[0].Category.Name="Господин"
	res.AwocList.Items[1].Persons.Items[0].Citizenship.Code=1
	res.AwocList.Items[1].Persons.Items[0].Citizenship.Name="Hе имеет значения"
	res.AwocList.Items[1].Persons.Items[1].LastName="ПЕТРОВА"
	res.AwocList.Items[1].Persons.Items[1].FirstName="ЕЛЕНА"
	res.AwocList.Items[1].Persons.Items[1].Company=""
	res.AwocList.Items[1].Persons.Items[1].Position=""
	res.AwocList.Items[1].Persons.Items[1].Category.Code=2
	res.AwocList.Items[1].Persons.Items[1].Category.Name="Госпожа"
	res.AwocList.Items[1].Persons.Items[1].Citizenship.Code=1
	res.AwocList.Items[1].Persons.Items[1].Citizenship.Name="Hе имеет значения"

	return res
}

func getRequestPriceRequestObject() *acaseSts.OrderInfoAwocNotifyRequestType {
	res := &acaseSts.OrderInfoAwocNotifyRequestType{}
	res.Id=2736033
	res.ReferenceNumber=""
	res.ReferenceNumberName="Ref No"
	res.RegistrationDate="12.08.2013"
	res.DeadlineDate="01.09.2013"
	res.Price=0.00
	res.TravelAgencyCommission=0.00
	res.Penalty=0.00
	res.StartDate="01.09.2013"
	res.EndDate="03.09.2013"
	res.InvoiceRule=2
	res.InvoiceId=0
	res.ConfirmId=2974355
	res.Currency.Code=2
	res.Currency.Name="RUB"
	res.WhereToPay.Code=3
	res.WhereToPay.Name="Оплата ООО \"АКАДЕМСЕРВИС\" согласно договору"
	res.Status.Code=1
	res.Status.Name="Запрос цены"
	res.ContactPerson.Name="test"
	res.ContactPerson.Phone=""
	res.ContactPerson.Fax=""
	res.ContactPerson.Email="test-mail@mail.ru"
	res.Customer.FullName="ООО Ромашка"
	res.Customer.ZipCode="123456"
	res.Customer.Address="Тестовая ул. д.5/2 оф. 5"
	res.Customer.PIAddress="Россия, 123456, Москва, Тестовая ул. д.5/2 оф. 5"
	res.Customer.INN="1234567890"
	res.Customer.KPP="123456789"
	res.Customer.Phone="+7 495 123 45 56"
	res.Customer.Name="Ромашка"
	res.Customer.Code=1322075
	res.Customer.BuyerType.Name="Юридическое лицо"
	res.Customer.BuyerType.Code=3
	res.Customer.Country.Name="Россия"
	res.Customer.Country.Code=9
	res.Customer.City.Name="Москва"
	res.Customer.City.Code=2
	res.Customer.City.CityType = &acaseSts.CityTypeType{}
	res.Customer.City.CityType.Name="Город"
	res.Customer.City.CityType.Code=2
	res.Customer.City.AdmUnit1 = &acaseSts.AdmUnit1Type{}
	res.Customer.City.AdmUnit1.Name="Не определено"
	res.Customer.City.AdmUnit1.Code=1
	res.Customer.City.AdmUnit2 = &acaseSts.AdmUnit2Type{}
	res.Customer.City.AdmUnit2.Name="Не определено"
	res.Customer.City.AdmUnit2.Code=1
	res.Customer.Actual = &acaseSts.ActualType{}
	res.Customer.Actual.Name="Да"
	res.Customer.Actual.Code=1
	res.Customer.AllowModify = &acaseSts.AllowModifyType{}
	res.Customer.AllowModify.Name="Да"
	res.Customer.AllowModify.Code=1
	res.IsCustomer.Code = 1
	res.IsCustomer.Name = "Да"
	res.AwocList.Items = make([]acaseSts.AwocType, 1)
	res.AwocList.Items[0].Id=3807880
	res.AwocList.Items[0].ParentId=0
	res.AwocList.Items[0].Subject="Гостиница Уют (Лапшинка) с 01.09.2013 по 03.09.2013, ПЕТРОВ ИЛЬЯ, ПЕТРОВА ЕЛЕНА"
	res.AwocList.Items[0].ArrivalDate="01.09.2013"
	res.AwocList.Items[0].DepartureDate="03.09.2013"
	res.AwocList.Items[0].NumberOfNights=2
	res.AwocList.Items[0].ArrivalTime=""
	res.AwocList.Items[0].DepartureTime=""
	res.AwocList.Items[0].Price=0.00
	res.AwocList.Items[0].VATIncludedInPrice=0.00
	res.AwocList.Items[0].TravelAgencyCommission=0.00
	res.AwocList.Items[0].DeadlineDate=""
	res.AwocList.Items[0].PossiblePenaltySize=0.00
	res.AwocList.Items[0].Penalty=0.00
	res.AwocList.Items[0].ConfirmationNumber=""
	res.AwocList.Items[0].ReferenceNumber=""
	res.AwocList.Items[0].SupplierInfo=""
	res.AwocList.Items[0].ToBeCancelled.Code=2
	res.AwocList.Items[0].ToBeCancelled.Name="Нет"
	res.AwocList.Items[0].ToBeChanged.Code=2
	res.AwocList.Items[0].ToBeChanged.Name="Нет"
	res.AwocList.Items[0].ToBeSelected.Code=2
	res.AwocList.Items[0].ToBeSelected.Name="Нет"
	res.AwocList.Items[0].Hotel.Code=1300725
	res.AwocList.Items[0].Hotel.Name="Уют Внуково МО"
	res.AwocList.Items[0].Country.Code=9
	res.AwocList.Items[0].Country.Name="Россия"
	res.AwocList.Items[0].AdmUnit1.Code=1
	res.AwocList.Items[0].AdmUnit1.Name="Не определено"
	res.AwocList.Items[0].AdmUnit2.Code=1
	res.AwocList.Items[0].AdmUnit2.Name="Не определено"
	res.AwocList.Items[0].TypeOfPlace.Code=2
	res.AwocList.Items[0].TypeOfPlace.Name="Город"
	res.AwocList.Items[0].City.Code=5254
	res.AwocList.Items[0].City.Name="Лапшинка"
	res.AwocList.Items[0].Status.Code=1
	res.AwocList.Items[0].Status.Name="Запрос цены"
	res.AwocList.Items[0].AllowCancel.Code=1
	res.AwocList.Items[0].AllowCancel.Name="Да"
	res.AwocList.Items[0].AllowModify.Code=2
	res.AwocList.Items[0].AllowModify.Name="Нет"
	res.AwocList.Items[0].AllowSelect.Code=2
	res.AwocList.Items[0].AllowSelect.Name="Нет"
	res.AwocList.Items[0].DetailList.Items = make([]acaseSts.DetailType, 1)
	res.AwocList.Items[0].DetailList.Items[0].Id=3807881
	res.AwocList.Items[0].DetailList.Items[0].Subject="Двухместный номер с большой кроватью. Завтрак."
	res.AwocList.Items[0].DetailList.Items[0].Quantity=1
	res.AwocList.Items[0].DetailList.Items[0].Duration=2.00
	res.AwocList.Items[0].DetailList.Items[0].Price=0.00
	res.AwocList.Items[0].DetailList.Items[0].TravelAgencyCommission=0.00
	res.AwocList.Items[0].DetailList.Items[0].VATIncludedInPrice=0.00
	res.AwocList.Items[0].Persons.Items = make([]acaseSts.PersonType, 2)
	res.AwocList.Items[0].Persons.Items[0].LastName="ПЕТРОВ"
	res.AwocList.Items[0].Persons.Items[0].FirstName="ИЛЬЯ"
	res.AwocList.Items[0].Persons.Items[0].Company=""
	res.AwocList.Items[0].Persons.Items[0].Position=""
	res.AwocList.Items[0].Persons.Items[0].Category.Code=1
	res.AwocList.Items[0].Persons.Items[0].Category.Name="Господин"
	res.AwocList.Items[0].Persons.Items[0].Citizenship.Code=1
	res.AwocList.Items[0].Persons.Items[0].Citizenship.Name="Hе имеет значения"
	res.AwocList.Items[0].Persons.Items[1].LastName="ПЕТРОВА"
	res.AwocList.Items[0].Persons.Items[1].FirstName="ЕЛЕНА"
	res.AwocList.Items[0].Persons.Items[1].Company=""
	res.AwocList.Items[0].Persons.Items[1].Position=""
	res.AwocList.Items[0].Persons.Items[1].Category.Code=2
	res.AwocList.Items[0].Persons.Items[1].Category.Name="Госпожа"
	res.AwocList.Items[0].Persons.Items[1].Citizenship.Code=1
	res.AwocList.Items[0].Persons.Items[1].Citizenship.Name="Hе имеет значения"

	return res
}

func getRequestReservationRequestObject() *acaseSts.OrderInfoAwocNotifyRequestType {
	res := &acaseSts.OrderInfoAwocNotifyRequestType{}
	res.Id=2736033
	res.ReferenceNumber=""
	res.ReferenceNumberName="Ref No"
	res.RegistrationDate="12.08.2013"
	res.DeadlineDate="30.08.2013"
	res.Price=7708.00
	res.TravelAgencyCommission=0.00
	res.Penalty=0.00
	res.StartDate="01.09.2013"
	res.EndDate="03.09.2013"
	res.InvoiceRule=2
	res.InvoiceId=4693145
	res.ConfirmId=0
	res.Currency.Code=2
	res.Currency.Name="RUB"
	res.WhereToPay.Code=3
	res.WhereToPay.Name="Оплата ООО \"АКАДЕМСЕРВИС\" согласно договору"
	res.Status.Code=11
	res.Status.Name="Получен"
	res.ContactPerson.Name="test"
	res.ContactPerson.Phone=""
	res.ContactPerson.Fax=""
	res.ContactPerson.Email="test-mail@mail.ru"
	res.Customer.FullName="ООО Ромашка"
	res.Customer.ZipCode="123456"
	res.Customer.Address="Тестовая ул. д.5/2 оф. 5"
	res.Customer.PIAddress="Россия, 123456, Москва, Тестовая ул. д.5/2 оф. 5"
	res.Customer.INN="1234567890"
	res.Customer.KPP="123456789"
	res.Customer.Phone="+7 495 123 45 56"
	res.Customer.Name="Ромашка"
	res.Customer.Code=1322075
	res.Customer.BuyerType.Name="Юридическое лицо"
	res.Customer.BuyerType.Code=3
	res.Customer.Country.Name="Россия"
	res.Customer.Country.Code=9
	res.Customer.City.Name="Москва"
	res.Customer.City.Code=2
	res.Customer.City.CityType = &acaseSts.CityTypeType{}
	res.Customer.City.CityType.Name="Город"
	res.Customer.City.CityType.Code=2
	res.Customer.City.AdmUnit1 = &acaseSts.AdmUnit1Type{}
	res.Customer.City.AdmUnit1.Name="Не определено"
	res.Customer.City.AdmUnit1.Code=1
	res.Customer.City.AdmUnit2 = &acaseSts.AdmUnit2Type{}
	res.Customer.City.AdmUnit2.Name="Не определено"
	res.Customer.City.AdmUnit2.Code=1
	res.Customer.Actual = &acaseSts.ActualType{}
	res.Customer.Actual.Name="Да"
	res.Customer.Actual.Code=1
	res.Customer.AllowModify = &acaseSts.AllowModifyType{}
	res.Customer.AllowModify.Name="Да"
	res.Customer.AllowModify.Code=1
	res.IsCustomer.Code = 1
	res.IsCustomer.Name = "Да"
	res.AwocList.Items = make([]acaseSts.AwocType, 2)
	res.AwocList.Items[0].Id=3807880
	res.AwocList.Items[0].ParentId=0
	res.AwocList.Items[0].Subject="Гостиница Уют (Лапшинка) с 01.09.2013 по 03.09.2013, ПЕТРОВ ИЛЬЯ, ПЕТРОВА ЕЛЕНА"
	res.AwocList.Items[0].ArrivalDate="01.09.2013"
	res.AwocList.Items[0].DepartureDate="03.09.2013"
	res.AwocList.Items[0].NumberOfNights=2
	res.AwocList.Items[0].ArrivalTime="12:00"
	res.AwocList.Items[0].DepartureTime="12:00"
	res.AwocList.Items[0].Price=7132.00
	res.AwocList.Items[0].VATIncludedInPrice=1087.93
	res.AwocList.Items[0].TravelAgencyCommission=0.00
	res.AwocList.Items[0].DeadlineDate="30.08.2013"
	res.AwocList.Items[0].PossiblePenaltySize=3105.00
	res.AwocList.Items[0].Penalty=0.00
	res.AwocList.Items[0].ConfirmationNumber=""
	res.AwocList.Items[0].ReferenceNumber=""
	res.AwocList.Items[0].SupplierInfo=""
	res.AwocList.Items[0].ToBeCancelled.Code=2
	res.AwocList.Items[0].ToBeCancelled.Name="Нет"
	res.AwocList.Items[0].ToBeChanged.Code=2
	res.AwocList.Items[0].ToBeChanged.Name="Нет"
	res.AwocList.Items[0].ToBeSelected.Code=2
	res.AwocList.Items[0].ToBeSelected.Name="Нет"
	res.AwocList.Items[0].Hotel.Code=1300725
	res.AwocList.Items[0].Hotel.Name="Уют Внуково МО"
	res.AwocList.Items[0].Country.Code=9
	res.AwocList.Items[0].Country.Name="Россия"
	res.AwocList.Items[0].AdmUnit1.Code=1
	res.AwocList.Items[0].AdmUnit1.Name="Не определено"
	res.AwocList.Items[0].AdmUnit2.Code=1
	res.AwocList.Items[0].AdmUnit2.Name="Не определено"
	res.AwocList.Items[0].TypeOfPlace.Code=2
	res.AwocList.Items[0].TypeOfPlace.Name="Город"
	res.AwocList.Items[0].City.Code=5254
	res.AwocList.Items[0].City.Name="Лапшинка"
	res.AwocList.Items[0].Status.Code=73
	res.AwocList.Items[0].Status.Name="Аннулировано"
	res.AwocList.Items[0].AllowCancel.Code=2
	res.AwocList.Items[0].AllowCancel.Name="Нет"
	res.AwocList.Items[0].AllowModify.Code=2
	res.AwocList.Items[0].AllowModify.Name="Нет"
	res.AwocList.Items[0].AllowSelect.Code=2
	res.AwocList.Items[0].AllowSelect.Name="Нет"
	res.AwocList.Items[0].DetailList.Items = make([]acaseSts.DetailType, 2)
	res.AwocList.Items[0].DetailList.Items[0].Id=3807881
	res.AwocList.Items[0].DetailList.Items[0].Subject="Стандартный двухместный номер с большой кроватью."
	res.AwocList.Items[0].DetailList.Items[0].Quantity=1
	res.AwocList.Items[0].DetailList.Items[0].Duration=2.00
	res.AwocList.Items[0].DetailList.Items[0].Price=5980.00
	res.AwocList.Items[0].DetailList.Items[0].TravelAgencyCommission=0.00
	res.AwocList.Items[0].DetailList.Items[0].VATIncludedInPrice=912.20
	res.AwocList.Items[0].DetailList.Items[1].Id=3807886
	res.AwocList.Items[0].DetailList.Items[1].Subject="Завтрак ШС"
	res.AwocList.Items[0].DetailList.Items[1].Quantity=2
	res.AwocList.Items[0].DetailList.Items[1].Duration=2.00
	res.AwocList.Items[0].DetailList.Items[1].Price=1152.00
	res.AwocList.Items[0].DetailList.Items[1].TravelAgencyCommission=0.00
	res.AwocList.Items[0].DetailList.Items[1].VATIncludedInPrice=175.73
	res.AwocList.Items[0].Persons.Items = make([]acaseSts.PersonType, 2)
	res.AwocList.Items[0].Persons.Items[0].LastName="ПЕТРОВ"
	res.AwocList.Items[0].Persons.Items[0].FirstName="ИЛЬЯ"
	res.AwocList.Items[0].Persons.Items[0].Company=""
	res.AwocList.Items[0].Persons.Items[0].Position=""
	res.AwocList.Items[0].Persons.Items[0].Category.Code=1
	res.AwocList.Items[0].Persons.Items[0].Category.Name="Господин"
	res.AwocList.Items[0].Persons.Items[0].Citizenship.Code=1
	res.AwocList.Items[0].Persons.Items[0].Citizenship.Name="Hе имеет значения"
	res.AwocList.Items[0].Persons.Items[1].LastName="ПЕТРОВА"
	res.AwocList.Items[0].Persons.Items[1].FirstName="ЕЛЕНА"
	res.AwocList.Items[0].Persons.Items[1].Company=""
	res.AwocList.Items[0].Persons.Items[1].Position=""
	res.AwocList.Items[0].Persons.Items[1].Category.Code=2
	res.AwocList.Items[0].Persons.Items[1].Category.Name="Госпожа"
	res.AwocList.Items[0].Persons.Items[1].Citizenship.Code=1
	res.AwocList.Items[0].Persons.Items[1].Citizenship.Name="Hе имеет значения"
	res.AwocList.Items[1].Id=3807882
	res.AwocList.Items[1].ParentId=3807880
	res.AwocList.Items[1].Subject="Гостиница Уют (Лапшинка) с 01.09.2013 по 03.09.2013, ПЕТРОВ ИЛЬЯ, ПЕТРОВА ЕЛЕНА"
	res.AwocList.Items[1].ArrivalDate="01.09.2013"
	res.AwocList.Items[1].DepartureDate="03.09.2013"
	res.AwocList.Items[1].NumberOfNights=2
	res.AwocList.Items[1].ArrivalTime="12:00"
	res.AwocList.Items[1].DepartureTime="12:00"
	res.AwocList.Items[1].Price=7708.00
	res.AwocList.Items[1].VATIncludedInPrice=1175.79
	res.AwocList.Items[1].TravelAgencyCommission=0.00
	res.AwocList.Items[1].DeadlineDate="30.08.2013"
	res.AwocList.Items[1].PossiblePenaltySize=3393.00
	res.AwocList.Items[1].Penalty=0.00
	res.AwocList.Items[1].ConfirmationNumber=""
	res.AwocList.Items[1].ReferenceNumber=""
	res.AwocList.Items[1].SupplierInfo=""
	res.AwocList.Items[1].ToBeCancelled.Code=2
	res.AwocList.Items[1].ToBeCancelled.Name="Нет"
	res.AwocList.Items[1].ToBeChanged.Code=2
	res.AwocList.Items[1].ToBeChanged.Name="Нет"
	res.AwocList.Items[1].ToBeSelected.Code=2
	res.AwocList.Items[1].ToBeSelected.Name="Нет"
	res.AwocList.Items[1].Hotel.Code=1300725
	res.AwocList.Items[1].Hotel.Name="Уют Внуково МО"
	res.AwocList.Items[1].Country.Code=9
	res.AwocList.Items[1].Country.Name="Россия"
	res.AwocList.Items[1].AdmUnit1.Code=1
	res.AwocList.Items[1].AdmUnit1.Name="Не определено"
	res.AwocList.Items[1].AdmUnit2.Code=1
	res.AwocList.Items[1].AdmUnit2.Name="Не определено"
	res.AwocList.Items[1].TypeOfPlace.Code=2
	res.AwocList.Items[1].TypeOfPlace.Name="Город"
	res.AwocList.Items[1].City.Code=5254
	res.AwocList.Items[1].City.Name="Лапшинка"
	res.AwocList.Items[1].Status.Code=11
	res.AwocList.Items[1].Status.Name="Запрос получен"
	res.AwocList.Items[1].AllowCancel.Code=1
	res.AwocList.Items[1].AllowCancel.Name="Да"
	res.AwocList.Items[1].AllowModify.Code=2
	res.AwocList.Items[1].AllowModify.Name="Нет"
	res.AwocList.Items[1].AllowSelect.Code=2
	res.AwocList.Items[1].AllowSelect.Name="Нет"
	res.AwocList.Items[1].DetailList.Items = make([]acaseSts.DetailType, 2)
	res.AwocList.Items[1].DetailList.Items[0].Id=3807883
	res.AwocList.Items[1].DetailList.Items[0].Subject="Улучшенный двухместный номер с большой кроватью."
	res.AwocList.Items[1].DetailList.Items[0].Quantity=1
	res.AwocList.Items[1].DetailList.Items[0].Duration=2.00
	res.AwocList.Items[1].DetailList.Items[0].Price=6556.00
	res.AwocList.Items[1].DetailList.Items[0].TravelAgencyCommission=0.00
	res.AwocList.Items[1].DetailList.Items[0].VATIncludedInPrice=1000.06
	res.AwocList.Items[1].DetailList.Items[1].Id=3807887
	res.AwocList.Items[1].DetailList.Items[1].Subject="Завтрак ШС"
	res.AwocList.Items[1].DetailList.Items[1].Quantity=2
	res.AwocList.Items[1].DetailList.Items[1].Duration=2.00
	res.AwocList.Items[1].DetailList.Items[1].Price=1152.00
	res.AwocList.Items[1].DetailList.Items[1].TravelAgencyCommission=0.00
	res.AwocList.Items[1].DetailList.Items[1].VATIncludedInPrice=175.73
	res.AwocList.Items[1].Persons.Items = make([]acaseSts.PersonType, 2)
	res.AwocList.Items[1].Persons.Items[0].LastName="ПЕТРОВ"
	res.AwocList.Items[1].Persons.Items[0].FirstName="ИЛЬЯ"
	res.AwocList.Items[1].Persons.Items[0].Company=""
	res.AwocList.Items[1].Persons.Items[0].Position=""
	res.AwocList.Items[1].Persons.Items[0].Category.Code=1
	res.AwocList.Items[1].Persons.Items[0].Category.Name="Господин"
	res.AwocList.Items[1].Persons.Items[0].Citizenship.Code=1
	res.AwocList.Items[1].Persons.Items[0].Citizenship.Name="Hе имеет значения"
	res.AwocList.Items[1].Persons.Items[1].LastName="ПЕТРОВА"
	res.AwocList.Items[1].Persons.Items[1].FirstName="ЕЛЕНА"
	res.AwocList.Items[1].Persons.Items[1].Company=""
	res.AwocList.Items[1].Persons.Items[1].Position=""
	res.AwocList.Items[1].Persons.Items[1].Category.Code=2
	res.AwocList.Items[1].Persons.Items[1].Category.Name="Госпожа"
	res.AwocList.Items[1].Persons.Items[1].Citizenship.Code=1
	res.AwocList.Items[1].Persons.Items[1].Citizenship.Name="Hе имеет значения"

	return res
}

/*
func TestGenerateXml_AwocNotify_Ok(t *testing.T) {
	item := getRequestPriceOfferRequestObject()
	bItem, err := xml.MarshalIndent(item, "", "    ")
	if err != nil {
		t.Error(err)
	}
	t.Log(xml.Header + string(bItem))
}
*/
func TestOrderNotifyRequest_RequestReservation_Ok(t *testing.T) {
	testRequest("ordernotify_hotelswrates_response.xml")
	defer gock.Off()

	item := getRequestReservationRequestObject()
	_, err := acApi.OrderInfoAwocNotifyRequest(item)
	er := getCustomErrorType()
	st.Expect(t, err, er)
}

func TestOrderNotifyRequest_RequestPriceRequest_Ok(t *testing.T) {
	testRequest("ordernotify_hotelswrates_response.xml")
	defer gock.Off()

	item := getRequestPriceRequestObject()
	_, err := acApi.OrderInfoAwocNotifyRequest(item)
	er := getCustomErrorType()
	st.Expect(t, err, er)
}


func TestOrderNotifyRequest_RequestPriceOffer_Ok(t *testing.T) {
	testRequest("ordernotify_hotelswrates_response.xml")
	defer gock.Off()

	item := getRequestPriceOfferRequestObject()
	_, err := acApi.OrderInfoAwocNotifyRequest(item)
	er := getCustomErrorType()
	st.Expect(t, err, er)
}

func TestOrderAwocNotifyRequest_Error(t *testing.T) {
	testRequest("orderinfonotify_error_example.xml")
	defer gock.Off()

	item := getRequestPriceOfferRequestObject()
	_, err := acApi.OrderInfoAwocNotifyRequest(item)

	st.Expect(t, err.Message, "Доступ запрещен !")
}

