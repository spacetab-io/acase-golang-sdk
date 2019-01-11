package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestSpecialOfferRequest_Ok(t *testing.T) {
	testRequest("specialoffer_response_example.xml", false)
	defer gockOff()

	data, err := acApi.SpecialOfferTypeRequest(context.Background(), 0, "")
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, len(data.SpecialOfferTypeList.SpecialOfferType), 3)
	st.Expect(t, data.SpecialOfferTypeList.SpecialOfferType[0].Code, 1)
	st.Expect(t, data.SpecialOfferTypeList.SpecialOfferType[0].Name, "-")
	st.Expect(t, data.SpecialOfferTypeList.SpecialOfferType[0].IsMultiple.Code, 2)
	st.Expect(t, data.SpecialOfferTypeList.SpecialOfferType[0].IsMultiple.Name, "Нет")
	st.Expect(t, data.SpecialOfferTypeList.SpecialOfferType[1].Code, 5)
	st.Expect(t, data.SpecialOfferTypeList.SpecialOfferType[1].Name, "Бесплатные ночи")
	st.Expect(t, data.SpecialOfferTypeList.SpecialOfferType[1].IsMultiple.Code, 2)
	st.Expect(t, data.SpecialOfferTypeList.SpecialOfferType[1].IsMultiple.Name, "Нет")
	st.Expect(t, data.SpecialOfferTypeList.SpecialOfferType[2].Code, 17)
	st.Expect(t, data.SpecialOfferTypeList.SpecialOfferType[2].Name, "Бесплатные ночи")
	st.Expect(t, data.SpecialOfferTypeList.SpecialOfferType[2].IsMultiple.Code, 1)
	st.Expect(t, data.SpecialOfferTypeList.SpecialOfferType[2].IsMultiple.Name, "Да")
}

func TestSpecialOfferRequest_Error(t *testing.T) {
	testRequest("specialoffer_error_example.xml", true)
	defer gockOff()

	_, err := acApi.SpecialOfferTypeRequest(context.Background(), 0, "")

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}
