package test

import (
	"testing"

	"github.com/nbio/st"
)

func TestCountryListRequest_Ok(t *testing.T)  {
	testRequest("countrylist_response_example.xml", false)
	defer gockOff()

	data, err := acApi.CountryListRequest(0,"")
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, len(data.Countries), 15)
	st.Expect(t, data.Countries[0].Code, 68)
	st.Expect(t, data.Countries[0].Name, "Азербайджан")
	st.Expect(t, data.Countries[0].Position.Latitude, "55.7513")
	st.Expect(t, data.Countries[0].Position.Longitude, "35.6156")
	st.Expect(t, data.Countries[1].Code, 78)
	st.Expect(t, data.Countries[1].Name, "Армения")
	st.Expect(t, data.Countries[1].Position.Latitude, "55.7513")
	st.Expect(t, data.Countries[1].Position.Longitude, "35.6156")
	st.Expect(t, data.Countries[2].Code, 34)
	st.Expect(t, data.Countries[2].Name, "Беларусь")
	st.Expect(t, data.Countries[2].Position.Latitude, "55.7513")
	st.Expect(t, data.Countries[2].Position.Longitude, "35.6156")
	st.Expect(t, data.Countries[3].Code, 74)
	st.Expect(t, data.Countries[3].Name, "Грузия")
	st.Expect(t, data.Countries[4].Code, 70)
	st.Expect(t, data.Countries[4].Name, "Казахстан")
	st.Expect(t, data.Countries[4].Position.Latitude, "55.7513")
	st.Expect(t, data.Countries[4].Position.Longitude, "35.6156")
	st.Expect(t, data.Countries[5].Code, 79)
	st.Expect(t, data.Countries[5].Name, "Кыргызстан")
	st.Expect(t, data.Countries[6].Code, 30)
	st.Expect(t, data.Countries[6].Name, "Латвия")
	st.Expect(t, data.Countries[7].Code, 58)
	st.Expect(t, data.Countries[7].Name, "Литва")
	st.Expect(t, data.Countries[8].Code, 75)
	st.Expect(t, data.Countries[8].Name, "Молдова")
	st.Expect(t, data.Countries[8].Position.Latitude, "55.7513")
	st.Expect(t, data.Countries[8].Position.Longitude, "35.6156")
	st.Expect(t, data.Countries[9].Code, 9)
	st.Expect(t, data.Countries[9].Name, "Россия")
	st.Expect(t, data.Countries[9].Position.Latitude, "55.7513")
	st.Expect(t, data.Countries[9].Position.Longitude, "35.6156")
	st.Expect(t, data.Countries[10].Code, 76)
	st.Expect(t, data.Countries[10].Name, "Таджикистан")
	st.Expect(t, data.Countries[10].Position.Latitude, "55.7513")
	st.Expect(t, data.Countries[10].Position.Longitude, "35.6156")
	st.Expect(t, data.Countries[11].Code, 77)
	st.Expect(t, data.Countries[11].Name, "Туркменистан")
	st.Expect(t, data.Countries[11].Position.Latitude, "55.7513")
	st.Expect(t, data.Countries[11].Position.Longitude, "35.6156")
	st.Expect(t, data.Countries[12].Code, 67)
	st.Expect(t, data.Countries[12].Name, "Узбекистан")
	st.Expect(t, data.Countries[12].Position.Latitude, "55.7513")
	st.Expect(t, data.Countries[12].Position.Longitude, "35.6156")
	st.Expect(t, data.Countries[13].Code, 38)
	st.Expect(t, data.Countries[13].Name, "Украина")
	st.Expect(t, data.Countries[13].Position.Latitude, "55.7513")
	st.Expect(t, data.Countries[13].Position.Longitude, "35.6156")
	st.Expect(t, data.Countries[14].Code, 36)
	st.Expect(t, data.Countries[14].Name, "Эстония")
	st.Expect(t, data.Countries[14].Position.Latitude, "55.7513")
	st.Expect(t, data.Countries[14].Position.Longitude, "35.6156")
}

func TestCountryListRequest_Error(t *testing.T) {
	testRequest("countrylist_error_example.xml", true)
	defer gockOff()

	_, err := acApi.CountryListRequest(0,"")

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}

