package test

import (
	"testing"
	"github.com/nbio/st"
)

func TestRateGroupRequest_Ok(t *testing.T) {
	testRequest("rategroup_response_example.xml", false)
	defer gockOff()
	pItem := make([]string, 1)
	pItem[0] = "Агент"
	data, err := acApi.RateGroupRequest(pItem)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, len(data.ActionList.Parameters) , 1)
	st.Expect(t, data.ActionList.Parameters[0].Name, "Агент")
	st.Expect(t, len(data.RateGroup) , 2)
	st.Expect(t, data.RateGroup[0].Code, 60)
	st.Expect(t, data.RateGroup[0].Name, "Агенты Hot Key")
	st.Expect(t, data.RateGroup[1].Code, 2)
	st.Expect(t, data.RateGroup[1].Name, "Агенты УК")
}

func TestRateGroupRequest_Error(t *testing.T) {
	testRequest("rategroup_error_example.xml", true)
	defer gockOff()

	_, err := acApi.RateGroupRequest(nil)

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}

