package test

import (
	"testing"

	"github.com/nbio/st"
)

func TestAdmUnit1Request_Ok(t *testing.T) {
	testRequest("au1req_example.xml")

	data, err := acApi.AdmUnit1Request(9, "", "обл")

	st.Expect(t, err, nil)
	st.Expect(t, len(data.AdmUnit1), 2)
	st.Expect(t, data.AdmUnit1[0].Code, 2)
	st.Expect(t, data.AdmUnit1[0].Name, "Калужская область")
	st.Expect(t, data.AdmUnit1[1].Code, 21)
	st.Expect(t, data.AdmUnit1[1].Name, "Тверская область")
}
