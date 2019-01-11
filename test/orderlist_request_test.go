package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestOrderListRequest_Ok(t *testing.T) {
	testRequest("orderlist_response_example.xml", false)
	defer gockOff()

	data, err := acApi.OrderListRequest(context.Background(), "", "", "", "", "",
		"", "", "", "", "", "",
		"", 0)

	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.IsThatAll.Code, 1)
	st.Expect(t, data.IsThatAll.Name, "Да")
	st.Expect(t, len(data.Orders.Items), 1)
	st.Expect(t, data.Orders.Items[0].Id, 495780)
	st.Expect(t, data.Orders.Items[0].ReferenceNumber, "CYXBW")
	st.Expect(t, data.Orders.Items[0].RegistrationDate, "07.10.2004")
	st.Expect(t, data.Orders.Items[0].DeadlineDate, "12.10.2004")
	st.Expect(t, data.Orders.Items[0].Price, 198.00)
	st.Expect(t, data.Orders.Items[0].TravelAgencyCommission, 0.00)
	st.Expect(t, data.Orders.Items[0].Penalty, 0.00)
	st.Expect(t, data.Orders.Items[0].StartDate, "13.10.2004")
	st.Expect(t, data.Orders.Items[0].EndDate, "15.10.2004")
	st.Expect(t, data.Orders.Items[0].Currency.Code, 3)
	st.Expect(t, data.Orders.Items[0].Currency.Name, "USD")
	st.Expect(t, data.Orders.Items[0].Status.Code, 65)
	st.Expect(t, data.Orders.Items[0].Status.Name, "Забронирован")
	st.Expect(t, data.Orders.Items[0].ContactPerson.Name, "Смирнов Павел")
	st.Expect(t, data.Orders.Items[0].ContactPerson.Phone, "")
	st.Expect(t, data.Orders.Items[0].ContactPerson.Fax, "")
	st.Expect(t, data.Orders.Items[0].ContactPerson.Email, "smirnov@smirnov.ru")
}

func TestOrderListRequest_Error(t *testing.T) {
	testRequest("orderlist_error_example.xml", true)
	defer gockOff()

	_, err := acApi.OrderListRequest(context.Background(), "", "", "", "", "",
		"", "", "", "", "", "",
		"", 0)

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}
