package test

import (
	"testing"

	"github.com/nbio/st"
)

func TestAdmUnit1Request_Ok(t *testing.T) {
	testRequest("admunit1_response_example.xml", false)
	defer gockOff()

	data, err := acApi.AdmUnit1Request(9, "", "обл")

	er := getCustomErrorType()
	st.Expect(t, err, er)
	if isMock {
		st.Expect(t, len(data.AdmUnit1), 2)
	} else {
		st.Expect(t, len(data.AdmUnit1), 88)
	}

	st.Expect(t, data.AdmUnit1[0].Code, 4)
	st.Expect(t, data.AdmUnit1[0].Name, "Адыгея республика")
	st.Expect(t, data.AdmUnit1[1].Code, 5)
	st.Expect(t, data.AdmUnit1[1].Name, "Алтай республика")
}

func TestAdmUnit1Request_Error(t *testing.T) {
	testRequest("admunit1_error_example.xml", true)
	defer gockOff()

	_, err := acApi.AdmUnit1Request(9, "", "обл")

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}
