package test

import (
	"testing"
	"github.com/nbio/st"
)

func TestMealRequest_Ok(t *testing.T) {
	testRequest("meal_response_example.xml", false)
	defer gockOff()

	data, err := acApi.MealRequest(0,0,"швед")
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.Action.Name, "LIST")
	st.Expect(t, data.Action.Parameters.MealCode, "")
	st.Expect(t, data.Action.Parameters.MealTypeCode, "")
	st.Expect(t, data.Action.Parameters.MealName, "швед")
	st.Expect(t, len(data.MealList.Items), 5)
	st.Expect(t, data.MealList.Items[0].Code, 3)
	st.Expect(t, data.MealList.Items[0].TypeCode, 2)
	st.Expect(t, data.MealList.Items[0].Name, "Завтрак \"Шведский стол\"")
}

func TestMealRequest_Error(t *testing.T) {
	testRequest("meal_error_example.xml", true)
	defer gockOff()

	_, err := acApi.MealRequest(0,0,"швед")

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}

