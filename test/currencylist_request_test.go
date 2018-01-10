package test

import (
	"testing"

	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
)

func TestCurrencyListRequest_Ok(t *testing.T)  {
	testRequest("currencylistresp_example.xml")
	defer gock.Off()

	data, err := acApi.CurrencyListRequest(0, "","")
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
	testRequest("currencylisterror_example.xml")
	defer gock.Off()

	_, err := acApi.CurrencyListRequest(0, "","")

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}
