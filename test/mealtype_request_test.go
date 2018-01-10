package test

import (
	"testing"
	"gopkg.in/h2non/gock.v1"
	"github.com/nbio/st"
)

func TestMealTypeRequest_Ok(t *testing.T) {
	testRequest("mealtype_response_example.xml")
	defer gock.Off()

	data, err := acApi.MealTypeRequest(nil,nil)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, len(data.Action), 1)
	st.Expect(t, data.Action[0].Name, "LIST")
	st.Expect(t, data.Action[0].Parameters.MealTypeCode, 0)
	st.Expect(t, data.Action[0].Parameters.MealName, "")
	st.Expect(t, len(data.MealTypeList.Items), 4)
	st.Expect(t, data.MealTypeList.Items[0].Code, 1)
	st.Expect(t, data.MealTypeList.Items[0].Name, "-")
	st.Expect(t, data.MealTypeList.Items[1].Code, 2)
	st.Expect(t, data.MealTypeList.Items[1].Name, "Завтрак")
}

func TestMealTypeRequest_Error(t *testing.T) {
	testRequest("mealtype_error_example.xml")
	defer gock.Off()

	_, err := acApi.MealTypeRequest(nil,nil)

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}

