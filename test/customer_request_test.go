package test

import (
	"testing"

	"github.com/nbio/st"
)

func TestCustomerCreateRequest_Ok(t *testing.T)  {
	testRequest("customer_create_response_example.xml", false)
	defer gockOff()

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
	testRequest("customer_error_example.xml", true)
	defer gockOff()

	_, err := acApi.CustomerRequestCreate("Международный клуб Ромашка", "123456",
		"Тестовая ул. д.5/2 кв.5","","123456789","","+371 1 222 333",
		"Ромашка","Юридическое лицо","Латвия","Рига",
		3,30,24)

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}

func TestCustomerUpdateRequest_Ok(t *testing.T)  {
	testRequest("customer_update_response_example.xml", false)
	defer gockOff()

	data, err := acApi.CustomerRequestUpdate("Международный клуб Ромашка", "123457",
		"Тестовая ул. д.5/2 кв.5","Латвия, 123457, Рига, Тестовая ул. д.5/2 кв.5",
		"123456789","","+371 1 22555 333","Ромашка","Юридическое лицо","Латвия","Рига",
		1322076,3,30,24)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.ActionUpdate.Parameters.CustomerCode, 1322076)
	st.Expect(t, data.ActionUpdate.Parameters.Customer.FullName, "Международный клуб Ромашка")
	st.Expect(t, data.ActionUpdate.Parameters.Customer.ZipCode, "123457")
	st.Expect(t, data.ActionUpdate.Parameters.Customer.Address, "Тестовая ул. д.5/2 кв.5")
	st.Expect(t, data.ActionUpdate.Parameters.Customer.PIAddress, "Латвия, 123457, Рига, Тестовая ул. д.5/2 кв.5")
	st.Expect(t, data.ActionUpdate.Parameters.Customer.INN, "123456789")
	st.Expect(t, data.ActionUpdate.Parameters.Customer.KPP, "")
	st.Expect(t, data.ActionUpdate.Parameters.Customer.Phone, "+371 1 22555 333")
	st.Expect(t, data.ActionUpdate.Parameters.Customer.Name, "Ромашка")
	st.Expect(t, data.ActionUpdate.Parameters.Customer.BuyerType.Code, 3)
	st.Expect(t, data.ActionUpdate.Parameters.Customer.BuyerType.Name, "Юридическое лицо")
	st.Expect(t, data.ActionUpdate.Parameters.Customer.Country.Code, 30)
	st.Expect(t, data.ActionUpdate.Parameters.Customer.Country.Name, "Латвия")
	st.Expect(t, data.ActionUpdate.Parameters.Customer.City.Code, 24)
	st.Expect(t, data.ActionUpdate.Parameters.Customer.City.Name, "Рига")

	st.Expect(t, data.Customer.FullName, "Международный клуб Ромашка")
	st.Expect(t, data.Customer.ZipCode, "123457")
	st.Expect(t, data.Customer.Address, "Тестовая ул. д.5/2 кв.5")
	st.Expect(t, data.Customer.PIAddress, "Латвия, 123457, Рига, Тестовая ул. д.5/2 кв.5")
	st.Expect(t, data.Customer.INN, "123456789")
	st.Expect(t, data.Customer.Phone, "+371 1 22555 333")
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

func TestCustomerDeleteRequest_Ok(t *testing.T)  {
	testRequest("customer_delete_response_example.xml", false)
	defer gockOff()

	data, err := acApi.CustomerRequestDelete(1322076)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.ActionDelete.Parameters.CustomerCode, 1322076)
}

func TestCustomerListRequest_Ok(t *testing.T)  {
	testRequest("customer_list_response_example.xml", false)
	defer gockOff()

	data, err := acApi.CustomerRequestList(1,1)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.ActionList.Parameters.Sort, 1)
	st.Expect(t, data.ActionList.Parameters.ShowActualOnly, 1)

	st.Expect(t, len(data.CustomerList.Customers), 2)

	st.Expect(t, data.CustomerList.Customers[0].FullName, "Международный клуб Ромашка")
	st.Expect(t, data.CustomerList.Customers[0].INN, "123456789")
	st.Expect(t, data.CustomerList.Customers[0].Code, 1322076)
	st.Expect(t, data.CustomerList.Customers[0].Name, "Ромашка")
	st.Expect(t, data.CustomerList.Customers[0].BuyerType.Code, 3)
	st.Expect(t, data.CustomerList.Customers[0].BuyerType.Name, "Юридическое лицо")
	st.Expect(t, data.CustomerList.Customers[0].Country.Code, 30)
	st.Expect(t, data.CustomerList.Customers[0].Country.Name, "Латвия")
	st.Expect(t, data.CustomerList.Customers[0].City.Code, 24)
	st.Expect(t, data.CustomerList.Customers[0].City.Name, "Рига")
	st.Expect(t, data.CustomerList.Customers[0].City.CityType.Code, 2)
	st.Expect(t, data.CustomerList.Customers[0].City.CityType.Name, "Город")
	st.Expect(t, data.CustomerList.Customers[0].City.AdmUnit1.Code, 1)
	st.Expect(t, data.CustomerList.Customers[0].City.AdmUnit1.Name, "Не определено")
	st.Expect(t, data.CustomerList.Customers[0].City.AdmUnit2.Code, 1)
	st.Expect(t, data.CustomerList.Customers[0].City.AdmUnit2.Name, "Не определено")
	st.Expect(t, data.CustomerList.Customers[0].Actual.Code, 1)
	st.Expect(t, data.CustomerList.Customers[0].Actual.Name, "Да")
	st.Expect(t, data.CustomerList.Customers[0].AllowModify.Code, 1)
	st.Expect(t, data.CustomerList.Customers[0].AllowModify.Name, "Да")

	st.Expect(t, data.CustomerList.Customers[1].FullName, "Лютик и Компания")
	st.Expect(t, data.CustomerList.Customers[1].INN, "123456789")
	st.Expect(t, data.CustomerList.Customers[1].Code, 1322077)
	st.Expect(t, data.CustomerList.Customers[1].Name, "Лютик")
	st.Expect(t, data.CustomerList.Customers[1].BuyerType.Code, 3)
	st.Expect(t, data.CustomerList.Customers[1].BuyerType.Name, "Юридическое лицо")
	st.Expect(t, data.CustomerList.Customers[1].Country.Code, 30)
	st.Expect(t, data.CustomerList.Customers[1].Country.Name, "Латвия")
	st.Expect(t, data.CustomerList.Customers[1].City.Code, 25)
	st.Expect(t, data.CustomerList.Customers[1].City.Name, "Юрмала")
	st.Expect(t, data.CustomerList.Customers[1].City.CityType.Code, 2)
	st.Expect(t, data.CustomerList.Customers[1].City.CityType.Name, "Город")
	st.Expect(t, data.CustomerList.Customers[1].City.AdmUnit1.Code, 1)
	st.Expect(t, data.CustomerList.Customers[1].City.AdmUnit1.Name, "Не определено")
	st.Expect(t, data.CustomerList.Customers[1].City.AdmUnit2.Code, 1)
	st.Expect(t, data.CustomerList.Customers[1].City.AdmUnit2.Name, "Не определено")
	st.Expect(t, data.CustomerList.Customers[1].Actual.Code, 1)
	st.Expect(t, data.CustomerList.Customers[1].Actual.Name, "Да")
	st.Expect(t, data.CustomerList.Customers[1].AllowModify.Code, 1)
	st.Expect(t, data.CustomerList.Customers[1].AllowModify.Name, "Да")
}

func TestCustomerInfoRequest_Ok(t *testing.T)  {
	testRequest("customer_info_response_example.xml", false)
	defer gockOff()

	data, err := acApi.CustomerRequestInfo(1322076)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.ActionInfo.Parameters.CustomerCode, 1322076)
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
