package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestCurrencyListRequest_Ok(t *testing.T) {
	testRequest("currencylist_response_example.xml", false)
	defer gockOff()

	data, err := acApi.CurrencyListRequest(context.Background(), 22, "", "")
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, len(data.Currency), 3)
	st.Expect(t, data.Currency[0].Code, 2)
	st.Expect(t, data.Currency[0].Name, "RUR")
	st.Expect(t, data.Currency[1].Code, 3)
	st.Expect(t, data.Currency[1].Name, "USD")
	st.Expect(t, data.Currency[2].Code, 5)
	st.Expect(t, data.Currency[2].Name, "EUR")
}

func TestCurrencyListRequest_Error(t *testing.T) {
	testRequest("currencylist_error_example.xml", true)
	defer gockOff()

	_, err := acApi.CurrencyListRequest(context.Background(), 0, "", "")

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}
