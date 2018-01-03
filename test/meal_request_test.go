package test

import (
	"testing"
	"gopkg.in/h2non/gock.v1"
	"github.com/nbio/st"
)

func TestMealRequest_Ok(t *testing.T) {
	testRequest("meal_response_example.xml")
	defer gock.Off()

	data, err := acApi.MealRequest(nil,nil,nil)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, len(data.Action), 1)
	st.Expect(t, data.Action[0].Name, "LIST")
	st.Expect(t, data.Action[0].Parameters.MealCode, 0)
	st.Expect(t, data.Action[0].Parameters.MealTypeCode, 0)
	st.Expect(t, data.Action[0].Parameters.MealName, "швед")
	st.Expect(t, len(data.MealList.Items), 5)
	st.Expect(t, data.MealList.Items[0].Code, 3)
	st.Expect(t, data.MealList.Items[0].TypeCode, 2)
	st.Expect(t, data.MealList.Items[0].Name, "Завтрак \"Шведский стол\"")
}

func TestMealRequest_Error(t *testing.T) {
	testRequest("meal_error_example.xml")
	defer gock.Off()

	_, err := acApi.MealRequest(nil,nil,nil)

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}

