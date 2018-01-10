package test

import (
	"testing"
	"gopkg.in/h2non/gock.v1"
	"github.com/nbio/st"
)

func TestTypeOfPlaceRequest_Ok(t *testing.T) {
	testRequest("typeofplace_response_example.xml")
	defer gock.Off()

	data, err := acApi.TypeOfPlaceRequest(0, "")
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.Action.Name, "LIST")
	st.Expect(t, len(data.TypeOfPlaceList.TypeOfPlace) , 2)
	st.Expect(t, data.TypeOfPlaceList.TypeOfPlace[0].Code, 1)
	st.Expect(t, data.TypeOfPlaceList.TypeOfPlace[0].Name, "Не определено")
	st.Expect(t, data.TypeOfPlaceList.TypeOfPlace[1].Code, 2)
	st.Expect(t, data.TypeOfPlaceList.TypeOfPlace[1].Name, "город")
}

func TestTypeOfPlaceRequest_Error(t *testing.T) {
	testRequest("typeofplace_error_example.xml")
	defer gock.Off()

	_, err := acApi.TypeOfPlaceRequest(0, "")

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}

