package test

import (
	"testing"

	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
)

func TestAdmUnit1Request_Ok(t *testing.T) {
	testRequest("au1resp_example.xml")
	defer gock.Off()

	data, err := acApi.AdmUnit1Request(9, "", "обл")

	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, len(data.AdmUnit1), 2)
	st.Expect(t, data.AdmUnit1[0].Code, 2)
	st.Expect(t, data.AdmUnit1[0].Name, "Калужская область")
	st.Expect(t, data.AdmUnit1[1].Code, 21)
	st.Expect(t, data.AdmUnit1[1].Name, "Тверская область")
}

func TestAdmUnit1Request_Error(t *testing.T) {
	testRequest("au1error_example.xml")
	defer gock.Off()

	_, err := acApi.AdmUnit1Request(9, "", "обл")

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}

/*
func TestSerializationRequest(t *testing.T) {
	req := &acaseStructs.AdmUnit1RequestType{
		Language: acaseStructs.Ru,
		Password: "Password",
		UserId: "UserId",
		BuyerId: "MyCompanyId",
		Action: acaseStructs.AdmUnit1ActionType{
			Name: acaseStructs.List,
			Parameters: acaseStructs.AdmUnit1ActionTypeParameters{
				CountryCode: 9,
				AdmUnit1Code: "",
				AdmUnit1Name: "обл",
			},
		},
	}
	bItem, err := xml.Marshal(req)
	if err != nil {
		t.Error(err)
	}
	t.Log(xml.Header + string(bItem))
	t.Log("TestSerializationRequest: Ok")
}
*/
