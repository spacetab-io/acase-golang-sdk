package test

import (
	"testing"
	"github.com/nbio/st"
)

func TestPenaltyReasonRequest_Ok(t *testing.T) {
	testRequest("penaltyreason_response_example.xml", false)
	defer gockOff()

	data, err := acApi.PenaltyReasonRequest()
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.Action.Name, "LISTPENALTY")
	st.Expect(t, len(data.PenaltyReason) , 4)
	st.Expect(t, data.PenaltyReason[0].Code, 1)
	st.Expect(t, data.PenaltyReason[0].Name, "-")
	st.Expect(t, data.PenaltyReason[1].Code, 3)
	st.Expect(t, data.PenaltyReason[1].Name, "Штраф за позднее изменение")
	st.Expect(t, data.PenaltyReason[2].Code, 4)
	st.Expect(t, data.PenaltyReason[2].Name, "Штраф за позднее аннулирование")
	st.Expect(t, data.PenaltyReason[3].Code, 5)
	st.Expect(t, data.PenaltyReason[3].Name, "Штраф за невостребование")
}

func TestPenaltyReasonRequest_Error(t *testing.T) {
	testRequest("penaltyreason_error_example.xml", true)
	defer gockOff()

	_, err := acApi.PenaltyReasonRequest()

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}

