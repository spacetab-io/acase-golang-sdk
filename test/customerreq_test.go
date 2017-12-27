package test

import (
	"testing"

	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
)

func TestCustomerCreateRequest_Ok(t *testing.T)  {
	testRequest("create_customer_response.xml")
	defer gock.Off()

	data, err := acApi.CustomerRequestCreate("Международный клуб Ромашка", "123456",
		"Тестовая ул. д.5/2 кв.5","","123456789","","+371 1 222 333",
		"Ромашка","Юридическое лицо","Латвия","Рига",
		3,30,24)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.ActionCreate.Parameters.Customer.FullName, "Международный клуб Ромашка")
	st.Expect(t, data.ActionCreate.Parameters.Customer.ZipCode, "123456")
	st.Expect(t, data.ActionCreate.Parameters.Customer.Address, "Тестовая ул. д.5/2 кв.5")
	st.Expect(t, data.ActionCreate.Parameters.Customer.PIAddress, "")
	st.Expect(t, data.ActionCreate.Parameters.Customer.INN, "123456789")
	st.Expect(t, data.ActionCreate.Parameters.Customer.KPP, "")
	st.Expect(t, data.ActionCreate.Parameters.Customer.Phone, "+371 1 222 333")
	st.Expect(t, data.ActionCreate.Parameters.Customer.Name, "Ромашка")
	st.Expect(t, data.ActionCreate.Parameters.Customer.BuyerType.Code, 3)
	st.Expect(t, data.ActionCreate.Parameters.Customer.BuyerType.Name, "Юридическое лицо")
	st.Expect(t, data.ActionCreate.Parameters.Customer.Country.Code, 30)
	st.Expect(t, data.ActionCreate.Parameters.Customer.Country.Name, "Латвия")
	st.Expect(t, data.ActionCreate.Parameters.Customer.City.Code, 24)
	st.Expect(t, data.ActionCreate.Parameters.Customer.City.Name, "Рига")

	st.Expect(t, data.Customer.FullName, "Международный клуб Ромашка")
	st.Expect(t, data.Customer.ZipCode, "123456")
	st.Expect(t, data.Customer.Address, "Тестовая ул. д.5/2 кв.5")
	st.Expect(t, data.Customer.PIAddress, "Латвия, 123456, Рига, Тестовая ул. д.5/2 кв.5")
	st.Expect(t, data.Customer.INN, "123456789")
	st.Expect(t, data.Customer.Phone, "+371 1 222 333")
	st.Expect(t, data.Customer.Code, 1322076)
	st.Expect(t, data.Customer.Name, "Ромашка")
	st.Expect(t, data.Customer.BuyerType.Code, 3)
	st.Expect(t, data.Customer.BuyerType.Name, "Юридическое лицо")
	st.Expect(t, data.Customer.Country.Code, 30)
	st.Expect(t, data.Customer.Country.Name, "Латвия")
	st.Expect(t, data.Customer.City.Code, 24)
	st.Expect(t, data.Customer.City.Name, "Рига")
	st.Expect(t, data.Customer.City.CityType.Code, 2)
	st.Expect(t, data.Customer.City.CityType.Name, "Город")
	st.Expect(t, data.Customer.City.AdmUnit1.Code, 1)
	st.Expect(t, data.Customer.City.AdmUnit1.Name, "Не определено")
	st.Expect(t, data.Customer.City.AdmUnit2.Code, 1)
	st.Expect(t, data.Customer.City.AdmUnit2.Name, "Не определено")
	st.Expect(t, data.Customer.Actual.Code, 1)
	st.Expect(t, data.Customer.Actual.Name, "Да")
	st.Expect(t, data.Customer.AllowModify.Code, 1)
	st.Expect(t, data.Customer.AllowModify.Name, "Да")
}

func TestCustomerRequest_Error(t *testing.T) {
	testRequest("create_customer_error_response.xml")
	defer gock.Off()

	_, err := acApi.CustomerRequestCreate("Международный клуб Ромашка", "123456",
		"Тестовая ул. д.5/2 кв.5","","123456789","","+371 1 222 333",
		"Ромашка","Юридическое лицо","Латвия","Рига",
		3,30,24)

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}

