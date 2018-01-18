package test

import (
	"testing"

	"github.com/nbio/st"
)

func TestAdmUnit2Request_Ok(t *testing.T) {
	testRequest("admunit2_response_example.xml", false)
	defer gockOff()

	data, err := acApi.AdmUnit2Request(9, 3, 0, "", "")

	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, len(data.AdmUnit2), 2)
	st.Expect(t, data.AdmUnit2[0].Code, 2)
	st.Expect(t, data.AdmUnit2[0].Name, "Тарусский район")
	st.Expect(t, data.AdmUnit2[0].AdmUnit1.Code, 3)
	st.Expect(t, data.AdmUnit2[0].AdmUnit1.Name, "Калужская область")
	st.Expect(t, data.AdmUnit2[1].Code, 3)
	st.Expect(t, data.AdmUnit2[1].Name, "Козельский район")
	st.Expect(t, data.AdmUnit2[1].AdmUnit1.Code, 3)
	st.Expect(t, data.AdmUnit2[1].AdmUnit1.Name, "Калужская область")
}

func TestAdmUnit2Request_Error(t *testing.T) {
	testRequest("admunit2_error_example.xml", true)
	defer gockOff()

	_, err := acApi.AdmUnit2Request(0, 3, 9, "", "")

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}

