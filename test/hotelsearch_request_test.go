package test

import (
	"testing"
	"gopkg.in/h2non/gock.v1"
	"github.com/nbio/st"
)

func TestRequest(t *testing.T) {

}

func TestSearchProductRequest_Ok(t *testing.T)  {
	testRequest("hotelsearch_by_hotellist_response_example.xml")
	defer gock.Off()

	_, err := acApi.HotelSearchRequest("","","","","",
		0, 0, 0, 0, 0, 0, 0,0,0,
		0,0,0.00,0.00,nil,nil)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
	//st.Expect(t, data, 0)
}

func TestHotelSearchRequest_Error(t *testing.T) {
	testRequest("hotelsearch_error_example.xml")
	defer gock.Off()

	_, err := acApi.HotelSearchRequest("","","","","",
		0, 0, 0, 0, 0, 0, 0,0,0,
		0,0,0.00,0.00,nil,nil)

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}

