package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestClientCategoryListRequest_Ok(t *testing.T) {
	testRequest("clientcategorylist_response_example.xml", true)
	defer gockOff()

	data, err := acApi.ClientCategoryListRequest(context.Background(), 0, "")
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, len(data.ClientCategories), 4)
	st.Expect(t, data.ClientCategories[0].Code, 1)
	st.Expect(t, data.ClientCategories[0].Name, "Господин")
	st.Expect(t, data.ClientCategories[1].Code, 2)
	st.Expect(t, data.ClientCategories[1].Name, "Госпожа")
	st.Expect(t, data.ClientCategories[2].Code, 3)
	st.Expect(t, data.ClientCategories[2].Name, "Ребенок")
	st.Expect(t, data.ClientCategories[3].Code, 4)
	st.Expect(t, data.ClientCategories[3].Name, "Младенец")
}

func TestClientCategoryListRequest_Error(t *testing.T) {
	testRequest("clientcategorylist_error_example.xml", false)
	defer gockOff()

	_, err := acApi.ClientCategoryListRequest(context.Background(), 0, "")

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}
