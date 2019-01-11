package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestMealTypeRequest_Ok(t *testing.T) {
	testRequest("mealtype_response_example.xml", false)
	defer gockOff()

	data, err := acApi.MealTypeRequest(context.Background(), 0, "")
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.Action.Name, "LIST")
	st.Expect(t, data.Action.Parameters.MealTypeCode, "")
	st.Expect(t, data.Action.Parameters.MealTypeName, "")
	st.Expect(t, len(data.MealTypeList.Items), 4)
	st.Expect(t, data.MealTypeList.Items[0].Code, 1)
	st.Expect(t, data.MealTypeList.Items[0].Name, "-")
	st.Expect(t, data.MealTypeList.Items[1].Code, 2)
	st.Expect(t, data.MealTypeList.Items[1].Name, "Завтрак")
}

func TestMealTypeRequest_Error(t *testing.T) {
	testRequest("mealtype_error_example.xml", true)
	defer gockOff()

	_, err := acApi.MealTypeRequest(context.Background(), 0, "")

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}
